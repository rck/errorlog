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

func TestLogWithIDs(t *testing.T) {
	l := NewErrorLogWithIDs()
	l.AppendWithID(errors.New("test1"), "myid1")
	l.AppendWithID(errors.New("test2"), "myid2")

	expected := 2
	got := l.Len()
	if expected != got {
		t.Fatalf("Expected length to be %d, but got: %d", expected, got)
	}

	e, err := l.GetID("myid2")
	if err != nil {
		t.Fatalf("Expected to find existing ID")
	}

	if e.Error() != "test2" {
		t.Fatalf("Expected error sting to be 'test2'")
	}

	if _, err := l.GetID("doesnotexist"); err == nil {
		t.Fatalf("Expected an error for non existing ID")
	}
}
