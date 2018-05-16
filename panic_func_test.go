package paranoid_test

import (
	"errors"
	"testing"

	"github.com/Fs02/go-paranoid"
)

func TestPanicFunc(t *testing.T) {
	called := false
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("not panic")
		}
		if called != true {
			t.Errorf("callback is called")
		}
	}()

	paranoid.PanicFunc(errors.New("errors ocurred"), "some context message", func() { called = true })
}

func TestPanicFuncNil(t *testing.T) {
	called := false
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic")
		}
		if called == true {
			t.Errorf("callback is not called")
		}
	}()

	paranoid.PanicFunc(nil, "some context message", func() { called = true })
}
