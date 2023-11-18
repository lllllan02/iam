package errors

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"strconv"
	"strings"
)

// Frame 表示堆栈帧内的程序计数器。
// 由于历史原因，如果 Frame 被解释为 uintptr，它的值表示程序计数器 +1。
type Frame uintptr

// pc 返回该 frame 的程序计数器，多个 frame 可能具有相同的 PC 值。
func (f Frame) pc() uintptr { return uintptr(f) - 1 }

// file 返回包含此 Frame 电脑功能的文件的完整路径。
func (f Frame) file() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	file, _ := fn.FileLine(f.pc())
	return file
}

// line 返回此 Frame 电脑的函数源代码行号。
func (f Frame) line() int {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(f.pc())
	return line
}

// name 返回此函数的名称（如果已知）。
func (f Frame) name() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

// Format 根据 mt.Formatter 格式化 frame。
//
//	%s    源代码文件名
//	%d    源代码行号
//	%n    函数名
//	%v    相当于 %s:%d
//
// Format 接受改变某些动词打印的标志，如下所示：
//
//	%+s   函数名和源文件相对于编译时 GOPATH 的路径以 \n\t（<函数名>\n\t<路径>）分隔
//	%+v   e相当于 %+s:%d
func (f Frame) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		switch {
		case s.Flag('+'):
			io.WriteString(s, f.name())
			io.WriteString(s, "\n\t")
			io.WriteString(s, f.file())
		default:
			io.WriteString(s, path.Base(f.file()))
		}
	case 'd':
		io.WriteString(s, strconv.Itoa(f.line()))
	case 'n':
		io.WriteString(s, funcname(f.name()))
	case 'v':
		f.Format(s, 's')
		io.WriteString(s, ":")
		f.Format(s, 'd')
	}
}

// MarshalText 将 stacktrace Frame 格式化为文本字符串。
// 输出与 fmt. Sprintf("%+v"，f) 相同，但没有换行符或制表符。
func (f Frame) MarshalText() ([]byte, error) {
	name := f.name()
	if name == "unknown" {
		return []byte(name), nil
	}
	return []byte(fmt.Sprintf("%s %s:%d", name, f.file(), f.line())), nil
}

// StackTrace 是从最内部（最新）到最外部（最旧）的帧堆 frame。
type StackTrace []Frame

// Format 根据 fmt.Formatter 格式化接口格式化 StackTrace。
//
//	%s	列出堆栈中每个 Frame 的源代码文件名
//	%v	列出堆栈中每个 Frame 的源代码行号
//
// Format 接受改变某些动词打印的标志，如下所示：
//
//	%+v   列出堆栈中每个 Frame 的文件名、函数名和行号
func (st StackTrace) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('+'):
			for _, f := range st {
				io.WriteString(s, "\n")
				f.Format(s, verb)
			}
		case s.Flag('#'):
			fmt.Fprintf(s, "%#v", []Frame(st))
		default:
			st.formatSlice(s, verb)
		}
	case 's':
		st.formatSlice(s, verb)
	}
}

// formatSlice 将此 StackTrace 格式化为给定缓冲区中的 Frame 切片，
// 仅在使用 "%s" 或 "%v" 调用时有效。
func (st StackTrace) formatSlice(s fmt.State, verb rune) {
	io.WriteString(s, "[")
	for i, f := range st {
		if i > 0 {
			io.WriteString(s, " ")
		}
		f.Format(s, verb)
	}
	io.WriteString(s, "]")
}

// stack 表示程序计数器的堆栈。
type stack []uintptr

func (s *stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				f := Frame(pc)
				fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

func (s *stack) StackTrace() StackTrace {
	f := make([]Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = Frame((*s)[i])
	}
	return f
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

// funcname 删除由 func.Name() 报告的函数名称的路径前缀组件。
func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}
