package brdoc

import (
	"testing"

	"github.com/jictyvoo/brelem/internal/docerrors"
)

func TestNewCPF_Valid(t *testing.T) {
	tests := []struct {
		input    string
		digits   string
		str      string
		obscure  string
		verifier [2]uint8
	}{
		{
			input: "298.074.850-10", digits: "29807485010",
			str: "298.074.850-10", obscure: "***.074.***-**",
			verifier: [2]uint8{1, 0},
		},
		{
			input: "241.525.560-21", digits: "24152556021",
			str: "241.525.560-21", obscure: "***.525.***-**",
			verifier: [2]uint8{2, 1},
		},
		{
			input: "501.585.380-72", digits: "50158538072",
			str: "501.585.380-72", obscure: "***.585.***-**",
			verifier: [2]uint8{7, 2},
		},
		{
			input: "97949526050", digits: "97949526050",
			str: "979.495.260-50", obscure: "***.495.***-**",
			verifier: [2]uint8{5, 0},
		},
		{
			input: "43793938018", digits: "43793938018",
			str: "437.939.380-18", obscure: "***.939.***-**",
			verifier: [2]uint8{1, 8},
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			cpf, err := NewCPF(tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := cpf.Digits(); got != tc.digits {
				t.Errorf("Digits() = %q, want %q", got, tc.digits)
			}
			if got := cpf.String(); got != tc.str {
				t.Errorf("String() = %q, want %q", got, tc.str)
			}
			if got := cpf.Obscure(); got != tc.obscure {
				t.Errorf("Obscure() = %q, want %q", got, tc.obscure)
			}
			if got := cpf.VerifierDigits(); got != tc.verifier {
				t.Errorf("VerifierDigits() = %v, want %v", got, tc.verifier)
			}
			if cpf.IsZero() {
				t.Error("IsZero() should be false for valid CPF")
			}
		})
	}
}

func TestNewCPF_Invalid(t *testing.T) {
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
		{"empty", ""},
		{"too short", "12"},
		{"extra digit", "241.525.560-210"},
		{"many extra", "298.074.850-1034"},
		{"wrong digits", "000.000.001-00"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewCPF(tc.input)
			if err == nil {
				t.Errorf("expected error for %q, got nil", tc.input)
			}
		})
	}
}

func TestNewCPF_ErrorTypes(t *testing.T) {
	_, err := NewCPF("")
	if !IsCPFError(err) {
		t.Errorf("expected IsCPFError for empty, got: %v", err)
	}
	if !docerrors.IsEmpty(err) {
		t.Errorf("expected IsEmpty for empty, got: %v", err)
	}

	_, err = NewCPF("000.000.000-00")
	if !IsCPFError(err) {
		t.Error("expected IsCPFError for all-same-digit")
	}

	_, err = NewCPF("000.000.001-00")
	if !docerrors.IsInvalidDigits(err) {
		t.Errorf("expected IsInvalidDigits, got: %v", err)
	}
}

func TestCPF_IsZero(t *testing.T) {
	var zero CPF
	if !zero.IsZero() {
		t.Error("zero CPF should be IsZero")
	}
}

func BenchmarkNewCPF(b *testing.B) {
	b.ReportAllocs()
	inputs := [...]string{
		"298.074.850-10", "241.525.560-21", "501.585.380-72", "97949526050", "43793938018",
	}
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			if _, err := NewCPF(input); err != nil {
				b.Fatal(err)
			}
		}
	}
}
