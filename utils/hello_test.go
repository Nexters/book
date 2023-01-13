package utils

import "testing"

func TestHell(t *testing.T) {
	if res := Hello(); res != "hello" {
		t.Errorf("expected: 'hello' but received: %v", res)
	}
}
