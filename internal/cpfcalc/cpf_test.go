package cpfcalc

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
		{"298.074.850-10", [2]uint8{1, 0}},
		{"241.525.560-21", [2]uint8{2, 1}},
		{"501.585.380-72", [2]uint8{7, 2}},
		{"97949526050", [2]uint8{5, 0}},
		{"43793938018", [2]uint8{1, 8}},
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
		{"all zeros", "000.000.000-00"},
		{"all ones", "111.111.111-11"},
		{"all twos", "222.222.222-22"},
		{"all threes", "33333333333"},
		{"all fours", "44444444444"},
		{"all fives", "55555555555"},
		{"all sixes", "66666666666"},
		{"all sevens", "77777777777"},
		{"all eights", "88888888888"},
		{"all nines", "99999999999"},
		{"incomplete", "111.111."},
		{"wrong check", "46477108071"},
		{"too short", "12"},
		{"extra digit after valid", "241.525.560-210"},
		{"many extra digits", "298.074.850-1034"},
		{"wrong digits", "000.000.001-00"},
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
		"298.074.850-10", "241.525.560-21", "501.585.380-72", "97949526050", "43793938018",
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
	feedString(&v, "000.000.000-00")
	_, reason := v.Finish()
	if reason != docerrors.ReasonIncorrectLength {
		t.Errorf("expected ReasonIncorrectLength, got: %d", reason)
	}

	v2 := New()
	feedString(&v2, "000.000.001-00")
	_, reason2 := v2.Finish()
	if reason2 != docerrors.ReasonInvalidDigits {
		t.Errorf("expected ReasonInvalidDigits, got: %d", reason2)
	}
}
