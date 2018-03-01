package errorlog

import (
	"errors"
	"testing"
)

func TestLog(t *testing.T) {
	l := NewErrorLog()

	msgs := []string{"First", "Second"}

	i := 1
	for _, msg := range msgs {
		l.Append(errors.New(msg))

		length := l.Len()
		if length != i {
			t.Fatalf("Expected Len() to be %d, but got %d", i, length)
		}

		lenErrs := len(l.Errs())
		if lenErrs != i {
			t.Fatalf("Expected len(Errors()) to be %d, but got %d", i, lenErrs)
		}
		i++
	}

	errs := l.Errs()

	for i, err := range errs {
		expect := msgs[i]
		got := err.Error()
		if expect != got {
			t.Fatalf("Expected error string to be %s, but got %s", expect, got)
		}
	}
}
