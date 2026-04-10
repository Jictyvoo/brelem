package cnhcalc

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

func validCNHs() [14]string {
	return [...]string{
		"22522791500", "46613298880", "18271939762", "03621746707", "26606066255", "33235836407", "35318223990",
		"96195156706", "33004095866", "87672878823", "33282790801", "29231325340", "62504334773", "37882306567",
	}
}

func TestValidator_Valid(t *testing.T) {
	for _, input := range validCNHs() {
		t.Run(input, func(t *testing.T) {
			v := New()
			feedString(&v, input)
			_, reason := v.Finish()
			if reason != docerrors.ReasonNoError {
				t.Fatalf("unexpected reason: %d", reason)
			}
		})
	}
}

func TestValidator_Invalid(t *testing.T) {
	tests := [...]string{
		"30231030596", "66199902513", "80169632370", "92505342466", "65022784469", "51070280999", "58049585715",
		"00618463148", "45643514075", "27850130146", "50535482964", "72800394363", "34237320316", "57271673315",
		"90356036254", "42645673394", "81866483793", "58241900906", "19268477514", "06990392146", "46629197135",
		"11891511083", "05305602596", "55512917711", "57056475202", "47840603918", "64959312184", "90712479234",
		"52372377844", "48146789947", "32514787270", "07006144716", "00506885075", "66464618306", "86129232978",
		"49545976082", "20539958980", "98316782009", "35440631332",
		// all same digit
		"00000000000", "11111111111", "22222222222", "33333333333", "44444444444", "55555555555", "66666666666",
		"77777777777", "88888888888", "99999999999",
		// too long
		"983820091678241",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			v := New()
			feedString(&v, input)
			_, reason := v.Finish()
			if reason == docerrors.ReasonNoError {
				t.Errorf("expected error for %q, got ReasonNoError", input)
			}
		})
	}
}

func BenchmarkValidator(b *testing.B) {
	b.ReportAllocs()
	inputs := validCNHs()
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
	feedString(&v, "00000000000")
	_, reason := v.Finish()
	if reason != docerrors.ReasonIncorrectLength {
		t.Errorf("expected ReasonIncorrectLength for all-same-digit, got: %d", reason)
	}

	v2 := New()
	feedString(&v2, "30231030596")
	_, reason2 := v2.Finish()
	if reason2 != docerrors.ReasonInvalidDigits {
		t.Errorf("expected ReasonInvalidDigits, got: %d", reason2)
	}
}
