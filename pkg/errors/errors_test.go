package errors

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestWrapNil(t *testing.T) {
	got := Wrap(nil, "no error")
	if got != nil {
		t.Errorf("Wrap(nil, \"no error\"): got %#v, expected nil", got)
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		err     error
		message string
		want    string
	}{
		{io.EOF, "read error", "read error"},
		{Wrap(io.EOF, "read error"), "client error", "client error"},
	}

	for _, tt := range tests {
		got := Wrap(tt.err, tt.message).Error()
		if got != tt.want {
			t.Errorf("Wrap(%v, %q): got: %v, want %v", tt.err, tt.message, got, tt.want)
		}
	}
}

type nilError struct{}

func (nilError) Error() string { return "nil error" }

func TestCause(t *testing.T) {
	tests := []struct {
		err  error
		want error
	}{{
		// nil error is nil
		err:  nil,
		want: nil,
	}, {
		// explicit nil error is nil
		err:  (error)(nil),
		want: nil,
	}, {
		// typed nil is nil
		err:  (*nilError)(nil),
		want: (*nilError)(nil),
	}, {
		// uncaused error is unaffected
		err:  io.EOF,
		want: io.EOF,
	}, {
		// caused error returns cause
		err:  Wrap(io.EOF, "ignored"),
		want: io.EOF,
	}}

	for i, tt := range tests {
		got := Cause(tt.err)
		if got != tt.want {
			t.Errorf("test %d: got %#v, want %#v", i+1, got, tt.want)
		}
	}
}

func TestWithCode(t *testing.T) {
	tests := []struct {
		code     int
		message  string
		wantType string
		wantCode int
	}{
		{1, "ConfigurationNotValid error", "*withCode", 1},
	}

	for _, tt := range tests {
		got := WithCode(tt.code, tt.message)
		err, ok := got.(*withCode)
		if !ok {
			t.Errorf("WithCode(%v, %q): error type got: %T, want %s", tt.code, tt.message, got, tt.wantType)
		}

		if Code(err) != tt.wantCode {
			t.Errorf("WithCode(%v, %q): got: %v, want %v", tt.code, tt.message, err.code, tt.wantCode)
		}
	}
}

func TestWithCodef(t *testing.T) {
	tests := []struct {
		code       int
		format     string
		args       string
		wantType   string
		wantCode   int
		wangString string
	}{
		{1, "ConfigurationNotValid %s", "error", "*withCode", 1, `ConfigurationNotValid error`},
	}

	for _, tt := range tests {
		got := WithCode(tt.code, tt.format, tt.args)
		err, ok := got.(*withCode)
		if !ok {
			t.Errorf("WithCode(%v, %q %q): error type got: %T, want %s", tt.code, tt.format, tt.args, got, tt.wantType)
		}

		if err.code != tt.wantCode {
			t.Errorf("WithCode(%v, %q %q): got: %v, want %v", tt.code, tt.format, tt.args, err.code, tt.wantCode)
		}

		if got.Error() != tt.wangString {
			t.Errorf("WithCode(%v, %q %q): got: %v, want %v", tt.code, tt.format, tt.args, got.Error(), tt.wangString)
		}
	}
}

func TestFormatWrap(t *testing.T) {
	tests := []struct {
		error
		format string
		want   string
	}{{
		Wrap(fmt.Errorf("error"), "error2"),
		"%s",
		"error2",
	}, {
		Wrap(fmt.Errorf("error"), "error2"),
		"%v",
		"error2",
	}, {
		Wrap(fmt.Errorf("error"), "error2"),
		"%+v",
		"error2 - #1 [/Users/lllllan/Downloads/github/lllllan02/iam/pkg/errors/errors_test.go:139 (github.com/lllllan02/iam/pkg/errors.TestFormatWrap)] (0) error2; \n\terror - #0 error",
	}, {
		Wrap(io.EOF, "error"),
		"%s",
		"error",
	}, {
		Wrap(io.EOF, "error"),
		"%v",
		"error",
	}, {
		Wrap(io.EOF, "error"),
		"%+v",
		"error - #1 [/Users/lllllan/Downloads/github/lllllan02/iam/pkg/errors/errors_test.go:151 (github.com/lllllan02/iam/pkg/errors.TestFormatWrap)] (0) error; \n\t" +
			"EOF - #0 EOF",
	}, {
		Wrap(
			Wrap(io.EOF, "error1"),
			"error2",
		),
		"%+v",
		"error2 - #2 [/Users/lllllan/Downloads/github/lllllan02/iam/pkg/errors/errors_test.go:156 (github.com/lllllan02/iam/pkg/errors.TestFormatWrap)] (0) error2; \n\t" +
			"error1 - #1 [/Users/lllllan/Downloads/github/lllllan02/iam/pkg/errors/errors_test.go:157 (github.com/lllllan02/iam/pkg/errors.TestFormatWrap)] (0) error1; \n\t" +
			"EOF - #0 EOF",
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}

func testFormatRegexp(t *testing.T, n int, arg interface{}, format, want string) {
	t.Helper()
	got := fmt.Sprintf(format, arg)
	gotLines := strings.Split(got, "\n\t")
	wantLines := strings.Split(want, "\n\t")

	if len(wantLines) > len(gotLines) {
		t.Errorf("test %d: wantLines(%d) > gotLines(%d):\n got: %q\nwant: %q", n+1, len(wantLines), len(gotLines), got, want)
		return
	}

	for i, w := range wantLines {
		if w != gotLines[i] {
			t.Errorf("test %d: line %d: fmt.Sprintf(%q, err):\n got: %q\nwant: %q", n+1, i+1, format, got, want)
		}
	}
}
