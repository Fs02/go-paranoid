package paranoid_test

import (
	"errors"
	"testing"

	"github.com/Fs02/go-paranoid"
)

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("not panic")
		}
	}()

	paranoid.Panic(errors.New("errors ocurred"))
}

func TestPanicNil(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic")
		}
	}()

	paranoid.Panic(nil)
}
