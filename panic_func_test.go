package paranoid_test

import (
	"errors"
	"testing"

	"github.com/Fs02/paranoid"
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

	paranoid.PanicFunc(errors.New("errors ocurred"), func() { called = true }, "some context %s", "message")
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

	paranoid.PanicFunc(nil, func() { called = true }, "some context message")
}
