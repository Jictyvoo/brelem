package brdoc

import (
	"github.com/jictyvoo/brelem/internal/cnpjcalc"
	"github.com/jictyvoo/brelem/internal/docerrors"
)

// LengthCNPJ is the number of digits in a valid CNPJ.
const LengthCNPJ = cnpjcalc.Length

// CNPJ represents a validated Brazilian CNPJ document.
// Each byte in digits stores a numeric value 0-9.
type CNPJ struct {
	digits [LengthCNPJ]byte
}

// NewCNPJ parses and validates a CNPJ string. Accepts formatted ("XX.XXX.XXX/XXXX-XX")
// or unformatted ("XXXXXXXXXXXXXX") input. Returns zero-value CNPJ and error on invalid input.
func NewCNPJ(raw string) (CNPJ, error) {
	var (
		c     CNPJ
		v     = cnpjcalc.New()
		count uint
	)

	for _, r := range raw {
		if r >= '0' && r <= '9' {
			if count < LengthCNPJ {
				c.digits[count] = byte(r - '0')
			}
			v.Feed(r)
			count++
		}
	}

	if count == 0 {
		return CNPJ{}, docerrors.New(docerrors.ReasonEmpty, "cnpj")
	}

	if _, reason := v.Finish(); reason != docerrors.ReasonNoError {
		return CNPJ{}, docerrors.New(reason, "cnpj")
	}
	return c, nil
}

// VerifierDigits returns the two check digits as numeric values (0-9).
func (c CNPJ) VerifierDigits() [2]uint8 {
	return [2]uint8{c.digits[12], c.digits[13]}
}

// String returns the CNPJ formatted as "XX.XXX.XXX/XXXX-XX".
func (c CNPJ) String() string {
	var buf [18]byte // 14 digits + 2 dots + 1 slash + 1 dash
	buf[0] = c.digits[0] + '0'
	buf[1] = c.digits[1] + '0'
	buf[2] = '.'
	buf[3] = c.digits[2] + '0'
	buf[4] = c.digits[3] + '0'
	buf[5] = c.digits[4] + '0'
	buf[6] = '.'
	buf[7] = c.digits[5] + '0'
	buf[8] = c.digits[6] + '0'
	buf[9] = c.digits[7] + '0'
	buf[10] = '/'
	buf[11] = c.digits[8] + '0'
	buf[12] = c.digits[9] + '0'
	buf[13] = c.digits[10] + '0'
	buf[14] = c.digits[11] + '0'
	buf[15] = '-'
	buf[16] = c.digits[12] + '0'
	buf[17] = c.digits[13] + '0'
	return string(buf[:])
}

// Digits returns the raw 14-digit string with no formatting.
func (c CNPJ) Digits() string {
	var buf [LengthCNPJ]byte
	for i, d := range c.digits {
		buf[i] = d + '0'
	}
	return string(buf[:])
}

// Obscure returns a masked representation "**.XXX.XXX/****-**".
// The middle groups (digits[2..7]) are visible; everything else is masked.
func (c CNPJ) Obscure() string {
	var buf [18]byte
	buf[0] = '*'
	buf[1] = '*'
	buf[2] = '.'
	buf[3] = c.digits[2] + '0'
	buf[4] = c.digits[3] + '0'
	buf[5] = c.digits[4] + '0'
	buf[6] = '.'
	buf[7] = c.digits[5] + '0'
	buf[8] = c.digits[6] + '0'
	buf[9] = c.digits[7] + '0'
	buf[10] = '/'
	buf[11] = '*'
	buf[12] = '*'
	buf[13] = '*'
	buf[14] = '*'
	buf[15] = '-'
	buf[16] = '*'
	buf[17] = '*'
	return string(buf[:])
}

// Subscription returns the 8-digit base registration number (first 8 digits).
func (c CNPJ) Subscription() string {
	var buf [8]byte
	for i := 0; i < 8; i++ {
		buf[i] = c.digits[i] + '0'
	}
	return string(buf[:])
}

// Branch returns the 4-digit branch number (digits 8-11).
func (c CNPJ) Branch() string {
	var buf [4]byte
	for i := 0; i < 4; i++ {
		buf[i] = c.digits[8+i] + '0'
	}
	return string(buf[:])
}

// IsZero reports whether the CNPJ is the zero value (not constructed via NewCNPJ).
func (c CNPJ) IsZero() bool {
	return c == CNPJ{}
}
