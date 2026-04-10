package cnpjcalc

import (
	"testing"

	"github.com/jictyvoo/brelem/internal/docerrors"
)

func feedString(v *Validator, s string) {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			v.Feed(r)
		}
	}
}

func TestValidator_Valid(t *testing.T) {
	tests := []struct {
		input    string
		verifier [2]uint8
	}{
		{"54.942.449/0001-06", [2]uint8{0, 6}},
		{"38.904.678/0001-71", [2]uint8{7, 1}},
		{"71.581.155/0001-07", [2]uint8{0, 7}},
		{"30.591.729/0001-40", [2]uint8{4, 0}},
		{"90718654000148", [2]uint8{4, 8}},
		{"73329532000140", [2]uint8{4, 0}},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			v := New()
			feedString(&v, tc.input)
			verifier, reason := v.Finish()
			if reason != docerrors.ReasonNoError {
				t.Fatalf("unexpected reason: %d", reason)
			}
			if verifier != tc.verifier {
				t.Errorf("got verifier %v, want %v", verifier, tc.verifier)
			}
		})
	}
}

func TestValidator_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"all zeros", "00.000.000/0000-00"},
		{"all ones", "11.111.111/1111-11"},
		{"all twos", "22.222.222/2222-22"},
		{"all threes", "33.333.333/3333-33"},
		{"all fours", "44.444.444/4444-44"},
		{"all fives", "55.555.555/5555-55"},
		{"all sixes", "66.666.666/6666-66"},
		{"all sevens", "77.777.777/7777-77"},
		{"all eights", "88.888.888/8888-88"},
		{"all nines", "99.999.999/9999-99"},
		{"incomplete", "69.372.070/0001-"},
		{"too many digits", "10004621459503"},
		{"wrong digits", "00.000.000/0001-00"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := New()
			feedString(&v, tc.input)
			_, reason := v.Finish()
			if reason == docerrors.ReasonNoError {
				t.Errorf("expected error for %q, got ReasonNoError", tc.input)
			}
		})
	}
}

func BenchmarkValidator(b *testing.B) {
	b.ReportAllocs()
	inputs := [...]string{
		"54.942.449/0001-06", "38.904.678/0001-71", "71.581.155/0001-07",
		"30.591.729/0001-40", "90718654000148", "73329532000140",
	}
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			v := New()
			feedString(&v, input)
			if _, reason := v.Finish(); reason != docerrors.ReasonNoError {
				b.Fatal(reason)
			}
		}
	}
}

func TestValidator_ErrorReason(t *testing.T) {
	v := New()
	feedString(&v, "00.000.000/0000-00")
	_, reason := v.Finish()
	if reason != docerrors.ReasonIncorrectLength {
		t.Errorf("expected ReasonIncorrectLength, got: %d", reason)
	}

	v2 := New()
	feedString(&v2, "00.000.000/0001-00")
	_, reason2 := v2.Finish()
	if reason2 != docerrors.ReasonInvalidDigits {
		t.Errorf("expected ReasonInvalidDigits, got: %d", reason2)
	}
}
