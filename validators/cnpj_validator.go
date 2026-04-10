package validators

import "github.com/jictyvoo/brelem/utils"

// LengthCNPJ The complete cnpj length (only numbers)
const LengthCNPJ = 14

type cnpjValidator struct {
	verifierDigits      [2]rune
	originalVerifier    [2]rune
	subscriptionWeights [2]rune
	branchWeights       [2]rune
	repeatedNumber      [2]rune
	iterations          uint
	extraDigits         uint64
}

func newCnpjValidator() cnpjValidator {
	return cnpjValidator{
		subscriptionWeights: [...]rune{5, 6},
		branchWeights:       [...]rune{9, 9},
		repeatedNumber:      [...]rune{-1, 0},
	}
}

func (v *cnpjValidator) iterateRune(intChar rune) {
	if v.iterations < LengthCNPJ-2 {
		if v.repeatedNumber[0] == -1 {
			// getting first number
			v.repeatedNumber[0] = intChar
		}

		if v.repeatedNumber[0] == intChar {
			// counting number repetitions
			v.repeatedNumber[1]++
		}
		if v.subscriptionWeights[0] >= 2 {
			v.verifierDigits[0] += (intChar - '0') * v.subscriptionWeights[0]
			v.subscriptionWeights[0]--
		} else {
			v.verifierDigits[0] += (intChar - '0') * v.branchWeights[0]
			v.branchWeights[0]--
		}
		if v.subscriptionWeights[1] >= 2 {
			v.verifierDigits[1] += (intChar - '0') * v.subscriptionWeights[1]
			v.subscriptionWeights[1]--
		} else {
			v.verifierDigits[1] += (intChar - '0') * v.branchWeights[1]
			v.branchWeights[1]--
		}
	} else {
		verifierIndex := 2 - (LengthCNPJ - v.iterations)
		digitValue := intChar - '0'
		if verifierIndex < 2 {
			v.originalVerifier[verifierIndex] = digitValue
		} else {
			v.extraDigits = (v.extraDigits << 1) | uint64(digitValue) // TODO: 0 is valid as last digit, this approach doesn't check it
		}
	}
	v.iterations++
}

func (v *cnpjValidator) applyVerifierDigits() {
	v.verifierDigits[0] = _modVerifierDigit(v.verifierDigits[0])
	v.verifierDigits[1] += v.verifierDigits[0] * v.branchWeights[1]
	v.verifierDigits[1] = _modVerifierDigit(v.verifierDigits[1])
}

func (v *cnpjValidator) finishValidation() (result error) {
	v.applyVerifierDigits()

	if v.HasIncorrectLength() {
		result = utils.ErrElementIncorrectLength
	} else if !v.HasValidDigits() {
		result = utils.ErrInvalidElement
	}
	return
}

func (v *cnpjValidator) HasIncorrectLength() bool {
	return v.repeatedNumber[1] >= LengthCNPJ-2 || v.iterations > LengthCNPJ || v.extraDigits != 0
}

func (v *cnpjValidator) HasValidDigits() bool {
	return v.originalVerifier == v.verifierDigits
}
