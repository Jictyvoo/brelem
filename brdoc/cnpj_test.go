package brdoc

import (
	"testing"

	"github.com/jictyvoo/brelem/internal/docerrors"
)

func TestNewCNPJ_Valid(t *testing.T) {
	tests := []struct {
		input        string
		digits       string
		str          string
		obscure      string
		subscription string
		branch       string
		verifier     [2]uint8
	}{
		{
			input: "54.942.449/0001-06", digits: "54942449000106",
			str: "54.942.449/0001-06", obscure: "**.942.449/****-**",
			subscription: "54942449", branch: "0001",
			verifier: [2]uint8{0, 6},
		},
		{
			input: "38.904.678/0001-71", digits: "38904678000171",
			str: "38.904.678/0001-71", obscure: "**.904.678/****-**",
			subscription: "38904678", branch: "0001",
			verifier: [2]uint8{7, 1},
		},
		{
			input: "71.581.155/0001-07", digits: "71581155000107",
			str: "71.581.155/0001-07", obscure: "**.581.155/****-**",
			subscription: "71581155", branch: "0001",
			verifier: [2]uint8{0, 7},
		},
		{
			input: "30.591.729/0001-40", digits: "30591729000140",
			str: "30.591.729/0001-40", obscure: "**.591.729/****-**",
			subscription: "30591729", branch: "0001",
			verifier: [2]uint8{4, 0},
		},
		{
			input: "90718654000148", digits: "90718654000148",
			str: "90.718.654/0001-48", obscure: "**.718.654/****-**",
			subscription: "90718654", branch: "0001",
			verifier: [2]uint8{4, 8},
		},
		{
			input: "73329532000140", digits: "73329532000140",
			str: "73.329.532/0001-40", obscure: "**.329.532/****-**",
			subscription: "73329532", branch: "0001",
			verifier: [2]uint8{4, 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			cnpj, err := NewCNPJ(tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := cnpj.Digits(); got != tc.digits {
				t.Errorf("Digits() = %q, want %q", got, tc.digits)
			}
			if got := cnpj.String(); got != tc.str {
				t.Errorf("String() = %q, want %q", got, tc.str)
			}
			if got := cnpj.Obscure(); got != tc.obscure {
				t.Errorf("Obscure() = %q, want %q", got, tc.obscure)
			}
			if got := cnpj.Subscription(); got != tc.subscription {
				t.Errorf("Subscription() = %q, want %q", got, tc.subscription)
			}
			if got := cnpj.Branch(); got != tc.branch {
				t.Errorf("Branch() = %q, want %q", got, tc.branch)
			}
			if got := cnpj.VerifierDigits(); got != tc.verifier {
				t.Errorf("VerifierDigits() = %v, want %v", got, tc.verifier)
			}
			if cnpj.IsZero() {
				t.Error("IsZero() should be false for valid CNPJ")
			}
		})
	}
}

func TestNewCNPJ_Invalid(t *testing.T) {
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
		{"empty", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewCNPJ(tc.input)
			if err == nil {
				t.Errorf("expected error for %q, got nil", tc.input)
			}
		})
	}
}

func TestNewCNPJ_ErrorTypes(t *testing.T) {
	_, err := NewCNPJ("")
	if !IsCNPJError(err) {
		t.Errorf("expected IsCNPJError for empty, got: %v", err)
	}
	if !docerrors.IsEmpty(err) {
		t.Errorf("expected IsEmpty for empty, got: %v", err)
	}

	_, err = NewCNPJ("00.000.000/0001-00")
	if !IsCNPJError(err) {
		t.Error("expected IsCNPJError for wrong digits")
	}
}

func TestCNPJ_IsZero(t *testing.T) {
	var zero CNPJ
	if !zero.IsZero() {
		t.Error("zero CNPJ should be IsZero")
	}
}

func BenchmarkNewCNPJ(b *testing.B) {
	b.ReportAllocs()
	inputs := [...]string{
		"54.942.449/0001-06", "38.904.678/0001-71", "71.581.155/0001-07",
		"30.591.729/0001-40", "90718654000148", "73329532000140",
	}
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			if _, err := NewCNPJ(input); err != nil {
				b.Fatal(err)
			}
		}
	}
}
