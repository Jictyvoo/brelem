package validators

import (
	"github.com/jictyvoo/brelem/utils"
)

// LengthCPF The CPF Length set to 11
const LengthCPF = 11

type cpfValidator struct {
	verifierDigits   [2]rune
	originalVerifier [2]rune
	weights          [2]rune
	repeatedNumber   [2]rune
	iterations       uint
	extraDigits      uint64
}

func newCpfValidator() cpfValidator {
	return cpfValidator{
		weights:        [...]rune{10, 11},
		repeatedNumber: [...]rune{-1, 0},
	}
}

func (v *cpfValidator) iterateRune(intChar rune) {
	if v.iterations < LengthCPF-2 {
		{
			v.verifierDigits[0] += (intChar - '0') * v.weights[0]
			v.verifierDigits[1] += (intChar - '0') * v.weights[1]
			v.weights[0]--
			v.weights[1]--
		}

		if v.repeatedNumber[0] == -1 {
			// getting first number
			v.repeatedNumber[0] = intChar
		}
		if v.repeatedNumber[0] == intChar {
			// counting number repetitions
			v.repeatedNumber[1]++
		}
	} else {
		verifierIndex := 2 - (LengthCPF - v.iterations)
		digitValue := intChar - '0'
		if verifierIndex < 2 {
			v.originalVerifier[verifierIndex] = digitValue
		} else {
			v.extraDigits = (v.extraDigits << 1) | uint64(digitValue)
		}
	}
	v.iterations++
}

func (v *cpfValidator) applyVerifierDigits() {
	v.verifierDigits[0] = _modVerifierDigit(v.verifierDigits[0])
	v.verifierDigits[1] += v.verifierDigits[0] * v.weights[1]
	v.verifierDigits[1] = _modVerifierDigit(v.verifierDigits[1])
}

func (v *cpfValidator) finishValidation() (result error) {
	v.applyVerifierDigits()

	if v.HasIncorrectLength() {
		result = utils.ErrElementIncorrectLength
	} else if !v.HasValidDigits() {
		result = utils.ErrInvalidElement
	}
	return
}

func (v *cpfValidator) HasIncorrectLength() bool {
	return v.repeatedNumber[1] >= LengthCPF-2 || v.iterations > LengthCPF || v.extraDigits != 0
}

func (v *cpfValidator) HasValidDigits() bool {
	return v.originalVerifier == v.verifierDigits
}
