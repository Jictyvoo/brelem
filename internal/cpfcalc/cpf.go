package cpfcalc

import (
	"github.com/jictyvoo/brelem/internal/digitutil"
	"github.com/jictyvoo/brelem/internal/docerrors"
)

// Length is the number of digits in a valid CPF.
const Length = 11

// Validator performs CPF check digit validation using rune-based arithmetic.
// Zero-alloc: all state lives on the stack.
type Validator struct {
	verifierDigits   [2]rune
	originalVerifier [2]rune
	weights          [2]rune
	repeatedNumber   [2]rune
	iterations       uint
	extraDigits      uint64
}

// New returns a Validator ready to receive digits.
func New() Validator {
	return Validator{
		weights:        [...]rune{10, 11},
		repeatedNumber: [...]rune{-1, 0},
	}
}

// Feed processes a single digit rune ('0'-'9').
// Non-digit runes must be filtered by the caller.
func (v *Validator) Feed(char rune) {
	if v.iterations < Length-2 {
		digit := char - '0'
		v.verifierDigits[0] += digit * v.weights[0]
		v.verifierDigits[1] += digit * v.weights[1]
		v.weights[0]--
		v.weights[1]--

		if v.repeatedNumber[0] == -1 {
			v.repeatedNumber[0] = char
		}
		if v.repeatedNumber[0] == char {
			v.repeatedNumber[1]++
		}
	} else {
		verifierIndex := 2 - (Length - v.iterations)
		digitValue := char - '0'
		if verifierIndex < 2 {
			v.originalVerifier[verifierIndex] = digitValue
		} else {
			v.extraDigits = (v.extraDigits << 1) | uint64(digitValue)
		}
	}
	v.iterations++
}

// Finish completes validation and returns the computed verifier digits
// and an ErrorReason indicating the result (ReasonNoError on success).
func (v *Validator) Finish() (verifier [2]uint8, reason docerrors.ErrorReason) {
	v.verifierDigits[0] = digitutil.ModVerifierDigit(v.verifierDigits[0])
	v.verifierDigits[1] += v.verifierDigits[0] * v.weights[1]
	v.verifierDigits[1] = digitutil.ModVerifierDigit(v.verifierDigits[1])

	verifier = [2]uint8{uint8(v.verifierDigits[0]), uint8(v.verifierDigits[1])}

	if v.repeatedNumber[1] >= Length-2 || v.iterations != Length || v.extraDigits != 0 {
		reason = docerrors.ReasonIncorrectLength
	} else if v.originalVerifier != v.verifierDigits {
		reason = docerrors.ReasonInvalidDigits
	}
	return
}
