package errors

import (
	"fmt"
	"testing"
)

func TestCallers(t *testing.T) {
	fmt.Printf("callers(): %v\n", callers())
}
