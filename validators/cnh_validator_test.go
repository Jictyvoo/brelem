package validators

import "testing"

func TestCNH(t *testing.T) {
	const value = "22522791500"
	if err := CNH(value); err != nil {
		t.Error(err)
	}
}
