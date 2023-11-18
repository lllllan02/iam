package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// WithCode 返回携带有错误码及堆栈信息的错误。
func WithCode(code int, format string, args ...interface{}) error {
	return &withCode{
		code:  code,
		err:   fmt.Errorf(format, args...),
		stack: callers(),
	}
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	return &withCode{
		code:  Code(err),
		err:   fmt.Errorf(message),
		cause: err,
		stack: callers(),
	}
}

// Wrapc returns an error annotating err with a stack trace at the
// point Wrapf is called, status code and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapc(err error, code int, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return &withCode{
		code:  code,
		err:   fmt.Errorf(format, args...),
		cause: err,
		stack: callers(),
	}
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return &withCode{
		code:  Code(err),
		err:   fmt.Errorf(format, args...),
		cause: err,
		stack: callers(),
	}
}

// Cause returns the underlying cause of the error, if possible.
func Cause(err error) error {
	if e, ok := err.(*withCode); ok {
		if e.cause == nil {
			return e
		}
		return Cause(e.cause)
	}

	return err
}

// Code returns the status code of the error, if possible.
func Code(err error) int {
	if e, ok := err.(*withCode); ok {
		return e.code
	}
	return 0
}

type withCode struct {
	*stack

	code  int
	err   error
	cause error
}

// Error return the externally-safe error message.
func (w *withCode) Error() string { return fmt.Sprintf("%v", w) }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withCode) Unwrap() error { return w.cause }

// Format implements fmt.Formatter.
//
// Verbs:
//
//	%s  - Returns the user-safe error string mapped to the error code or
//	  ┊   the error message if none is specified.
//	%v      Alias for %s
//
// Flags:
//
//	#      JSON formatted output, useful for logging
//	-      Output caller details, useful for troubleshooting
//	+      Output full error stack details, useful for debugging
//
// Examples:
//
//	%s:    error for internal read B
//	%v:    error for internal read B
//	%-v:   error for internal read B - #0 [/home/lk/workspace/golang/src/github.com/marmotedu/iam/main.go:12 (main.main)] (#100102) Internal Server Error
//	%+v:   error for internal read B - #0 [/home/lk/workspace/golang/src/github.com/marmotedu/iam/main.go:12 (main.main)] (#100102) Internal Server Error; error for internal read A - #1 [/home/lk/workspace/golang/src/github.com/marmotedu/iam/main.go:35 (main.newErrorB)] (#100104) Validation failed
//	%#v:   [{"error":"error for internal read B"}]
//	%#-v:  [{"caller":"#0 /home/lk/workspace/golang/src/github.com/marmotedu/iam/main.go:12 (main.main)","error":"error for internal read B","message":"(#100102) Internal Server Error"}]
//	%#+v:  [{"caller":"#0 /home/lk/workspace/golang/src/github.com/marmotedu/iam/main.go:12 (main.main)","error":"error for internal read B","message":"(#100102) Internal Server Error"},{"caller":"#1 /home/lk/workspace/golang/src/github.com/marmotedu/iam/main.go:35 (main.newErrorB)","error":"error for internal read A","message":"(#100104) Validation failed"}]
func (w *withCode) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		str := bytes.NewBuffer([]byte{})
		jsonData := []map[string]interface{}{}

		var (
			flagDetail bool
			flagTrace  bool
			modeJSON   bool
		)

		if state.Flag('#') {
			modeJSON = true
		}

		if state.Flag('-') {
			flagDetail = true
		}
		if state.Flag('+') {
			flagTrace = true
		}

		sep := ""
		errs := list(w)
		length := len(errs)
		for k, e := range errs {
			jsonData, str = format(length-k-1, jsonData, str, e, sep, flagDetail, flagTrace, modeJSON)
			sep = "; \n\t"

			if !flagTrace {
				break
			}

			if !flagDetail && !flagTrace && !modeJSON {
				break
			}
		}
		if modeJSON {
			var byts []byte
			byts, _ = json.Marshal(jsonData)

			str.Write(byts)
		}

		fmt.Fprintf(state, "%s", strings.Trim(str.String(), "\r\n\t"))
	default:
		fmt.Fprint(state, w.err.Error())
	}
}

func format(k int, jsonData []map[string]interface{}, str *bytes.Buffer, w *withCode,
	sep string, flagDetail, flagTrace, modeJSON bool) ([]map[string]interface{}, *bytes.Buffer) {
	if modeJSON {
		data := map[string]interface{}{}
		if flagDetail || flagTrace {
			data = map[string]interface{}{
				"message": w.err.Error(),
				"code":    w.code,
				"error":   w.err,
			}

			caller := fmt.Sprintf("#%d", k)
			if w.stack != nil {
				f := Frame((*w.stack)[0])
				caller = fmt.Sprintf("%s %s:%d (%s)",
					caller,
					f.file(),
					f.line(),
					f.name(),
				)
			}
			data["caller"] = caller
		} else {
			data["error"] = w.err.Error()
		}
		jsonData = append(jsonData, data)
	} else {
		if flagDetail || flagTrace {
			if w.stack != nil {
				f := Frame((*w.stack)[0])
				fmt.Fprintf(str, "%s%s - #%d [%s:%d (%s)] (%d) %s",
					sep,
					w.err,
					k,
					f.file(),
					f.line(),
					f.name(),
					w.code,
					w.err.Error(),
				)
			} else {
				fmt.Fprintf(str, "%s%s - #%d %s", sep, w.err, k, w.err.Error())
			}
		} else {
			fmt.Fprint(str, w.err.Error())
		}
	}

	return jsonData, str
}

// list will convert the error stack into a simple array.
func list(e error) []*withCode {
	ret := []*withCode{}

	if e != nil {
		if w, ok := e.(*withCode); ok {
			ret = append(ret, w)
			ret = append(ret, list(w.cause)...)
		} else {
			ret = append(ret, &withCode{err: e})
		}
	}

	return ret
}