package cnhcalc

import "github.com/jictyvoo/brelem/internal/docerrors"

// Length is the number of digits in a valid CNH.
const Length = 11


// Validator performs CNH check digit validation using rune-based arithmetic.
// Zero-alloc: all state lives on the stack (no slice allocation).
type Validator struct {
	verifierDigits   [2]rune
	originalVerifier [2]rune
	weights          [2]rune
	repeatedNumber   [2]rune
	iterations       uint
	verifierCount    uint
}

// New returns a Validator ready to receive digits.
func New() Validator {
	return Validator{
		weights:        [...]rune{9, 1},
		repeatedNumber: [...]rune{-1, 0},
	}
}

// Feed processes a single digit rune ('0'-'9').
func (v *Validator) Feed(char rune) {
	actualNumber := char - '0'
	if v.iterations < Length-2 {
		v.verifierDigits[0] += actualNumber * v.weights[0]
		v.verifierDigits[1] += actualNumber * v.weights[1]
		v.weights[1]++
		v.weights[0]--

		if v.repeatedNumber[0] == -1 {
			v.repeatedNumber[0] = char
		}
		if v.repeatedNumber[0] == char {
			v.repeatedNumber[1]++
		}
	} else if v.verifierCount < 2 {
		v.originalVerifier[v.verifierCount] = actualNumber
		v.verifierCount++
	}
	v.iterations++
}

// Finish completes validation and returns the computed verifier digits
// and an ErrorReason indicating the result (ReasonNoError on success).
func (v *Validator) Finish() (verifier [2]uint8, reason docerrors.ErrorReason) {
	var shouldRecalculate bool
	if v.verifierDigits[0] %= 11; v.verifierDigits[0] >= 10 {
		v.verifierDigits[0] = 0
		shouldRecalculate = true
	}

	if v.verifierDigits[1] %= 11; v.verifierDigits[1] >= 10 {
		v.verifierDigits[1] = 0
	} else if shouldRecalculate {
		v.verifierDigits[1] -= 2
	}

	verifier = [2]uint8{uint8(v.verifierDigits[0]), uint8(v.verifierDigits[1])}

	hasCorrectLength := v.repeatedNumber[1] < Length-2 && v.iterations == Length &&
		v.verifierCount == 2
	if !hasCorrectLength {
		reason = docerrors.ReasonIncorrectLength
	} else if v.originalVerifier[0] != v.verifierDigits[0] || v.originalVerifier[1] != v.verifierDigits[1] {
		reason = docerrors.ReasonInvalidDigits
	}
	return
}
