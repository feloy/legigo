package legigo

import (
	"testing"
)

func assertStringEquals(t *testing.T, expected string, value string) {
	if expected != value {
		t.Errorf("Error expected value '%s' is '%s'", expected, value)
	}
}

func assertLenEquals(t *testing.T, expected int, value int) {
	if expected != value {
		t.Errorf("Error expected array length '%d' is '%d'", expected, value)
	}
}
