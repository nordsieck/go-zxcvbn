package defect

import (
	"reflect"
	"testing"
)

// A low cost error
type String string

func (s String) Error() string { return string(s) }

func Equal(t *testing.T, obtained, expected interface{}) {
	if obtained != expected {
		t.Error("obtained:", obtained, "\n",
			"expected:", expected, "\n")
	}
}

func DeepEqual(t *testing.T, obtained, expected interface{}) {
	if !reflect.DeepEqual(obtained, expected) {
		t.Error("obtained:", obtained, "\n",
			"expected:", expected, "\n")
	}
}
