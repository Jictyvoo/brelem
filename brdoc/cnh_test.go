package brdoc

import (
	"testing"
	"time"

	"github.com/jictyvoo/brelem/internal/docerrors"
)

func validCNHInputs() [14]string {
	return [...]string{
		"22522791500", "46613298880", "18271939762", "03621746707", "26606066255", "33235836407", "35318223990",
		"96195156706", "33004095866", "87672878823", "33282790801", "29231325340", "62504334773", "37882306567",
	}
}

func TestNewCNH_Valid(t *testing.T) {
	for _, input := range validCNHInputs() {
		t.Run(input, func(t *testing.T) {
			cnh, err := NewCNH(input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := cnh.Digits(); got != input {
				t.Errorf("Digits() = %q, want %q", got, input)
			}
			if got := cnh.String(); got != input {
				t.Errorf("String() = %q, want %q", got, input)
			}
			if cnh.IsZero() {
				t.Error("IsZero() should be false for valid CNH")
			}
		})
	}
}

func TestNewCNH_Invalid(t *testing.T) {
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
		// empty
		"",
	}

	for _, input := range tests {
		name := input
		if name == "" {
			name = "empty"
		}
		t.Run(name, func(t *testing.T) {
			_, err := NewCNH(input)
			if err == nil {
				t.Errorf("expected error for %q, got nil", input)
			}
		})
	}
}

func TestNewCNH_Options(t *testing.T) {
	expires := time.Date(2030, 6, 15, 0, 0, 0, 0, time.UTC)
	firstLic := time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)

	cnh, err := NewCNH("22522791500",
		WithExpiresAt(expires),
		WithFirstLicenseDate(firstLic),
		WithType(DriverTypeB),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := cnh.ExpiresAt(); !got.Equal(expires) {
		t.Errorf("ExpiresAt() = %v, want %v", got, expires)
	}
	if got := cnh.FirstLicenseDate(); !got.Equal(firstLic) {
		t.Errorf("FirstLicenseDate() = %v, want %v", got, firstLic)
	}
	if got := cnh.Type(); got != DriverTypeB {
		t.Errorf("Type() = %v, want %v", got, DriverTypeB)
	}
	if cnh.IsExpired() {
		t.Error("CNH with future expiry should not be expired")
	}
}

func TestCNH_IsExpired(t *testing.T) {
	past := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	cnh, err := NewCNH("22522791500", WithExpiresAt(past))
	if err != nil {
		t.Fatal(err)
	}
	if !cnh.IsExpired() {
		t.Error("CNH with past expiry should be expired")
	}

	// No expiry set should not be expired
	cnh2, err := NewCNH("22522791500")
	if err != nil {
		t.Fatal(err)
	}
	if cnh2.IsExpired() {
		t.Error("CNH without expiry should not be expired")
	}
}

func TestNewCNH_ErrorTypes(t *testing.T) {
	_, err := NewCNH("")
	if !IsCNHError(err) {
		t.Errorf("expected IsCNHError for empty, got: %v", err)
	}
	if !docerrors.IsEmpty(err) {
		t.Errorf("expected IsEmpty for empty, got: %v", err)
	}
}

func TestCNH_IsZero(t *testing.T) {
	var zero CNH
	if !zero.IsZero() {
		t.Error("zero CNH should be IsZero")
	}
}

func BenchmarkNewCNH(b *testing.B) {
	b.ReportAllocs()
	inputs := validCNHInputs()
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			if _, err := NewCNH(input); err != nil {
				b.Fatal(err)
			}
		}
	}
}
