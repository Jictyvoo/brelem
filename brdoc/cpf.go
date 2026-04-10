package brdoc

import (
	"github.com/jictyvoo/brelem/internal/cpfcalc"
	"github.com/jictyvoo/brelem/internal/docerrors"
)

// LengthCPF is the number of digits in a valid CPF.
const LengthCPF = cpfcalc.Length

// CPF represents a validated Brazilian CPF document.
// Each byte in digits stores a numeric value 0-9.
type CPF struct {
	digits [LengthCPF]byte
}

// NewCPF parses and validates a CPF string. Accepts formatted ("XXX.XXX.XXX-XX")
// or unformatted ("XXXXXXXXXXX") input. Returns zero-value CPF and error on invalid input.
func NewCPF(raw string) (CPF, error) {
	var (
		c     CPF
		v     = cpfcalc.New()
		count uint
	)

	for _, r := range raw {
		if r >= '0' && r <= '9' {
			if count < LengthCPF {
				c.digits[count] = byte(r - '0')
			}
			v.Feed(r)
			count++
		}
	}

	if count == 0 {
		return CPF{}, docerrors.New(docerrors.ReasonEmpty, "cpf")
	}

	if _, reason := v.Finish(); reason != docerrors.ReasonNoError {
		return CPF{}, docerrors.New(reason, "cpf")
	}
	return c, nil
}

// VerifierDigits returns the two check digits as numeric values (0-9).
func (c CPF) VerifierDigits() [2]uint8 {
	return [2]uint8{c.digits[9], c.digits[10]}
}

// String returns the CPF formatted as "XXX.XXX.XXX-XX".
func (c CPF) String() string {
	var buf [14]byte // 11 digits + 2 dots + 1 dash
	buf[0] = c.digits[0] + '0'
	buf[1] = c.digits[1] + '0'
	buf[2] = c.digits[2] + '0'
	buf[3] = '.'
	buf[4] = c.digits[3] + '0'
	buf[5] = c.digits[4] + '0'
	buf[6] = c.digits[5] + '0'
	buf[7] = '.'
	buf[8] = c.digits[6] + '0'
	buf[9] = c.digits[7] + '0'
	buf[10] = c.digits[8] + '0'
	buf[11] = '-'
	buf[12] = c.digits[9] + '0'
	buf[13] = c.digits[10] + '0'
	return string(buf[:])
}

// Digits returns the raw 11-digit string with no formatting.
func (c CPF) Digits() string {
	var buf [LengthCPF]byte
	for i, d := range c.digits {
		buf[i] = d + '0'
	}
	return string(buf[:])
}

// Obscure returns a masked representation "***.XXX.***-**".
// The middle group (digits[3..5]) is visible; everything else is masked.
func (c CPF) Obscure() string {
	var buf [14]byte
	buf[0] = '*'
	buf[1] = '*'
	buf[2] = '*'
	buf[3] = '.'
	buf[4] = c.digits[3] + '0'
	buf[5] = c.digits[4] + '0'
	buf[6] = c.digits[5] + '0'
	buf[7] = '.'
	buf[8] = '*'
	buf[9] = '*'
	buf[10] = '*'
	buf[11] = '-'
	buf[12] = '*'
	buf[13] = '*'
	return string(buf[:])
}

// IsZero reports whether the CPF is the zero value (not constructed via NewCPF).
func (c CPF) IsZero() bool {
	return c == CPF{}
}
