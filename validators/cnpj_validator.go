package validators

import "github.com/jictyvoo/brelem/utils"

// LengthCNPJ The complete cnpj length (only numbers)
const LengthCNPJ = 14

func CNPJ(element string) (result error) {
	var (
		verifierDigits      = [...]rune{0, 0}
		originalVerifier    = make([]rune, 0, 2)
		subscriptionWeights = [...]rune{5, 6}
		branchWeights       = [...]rune{9, 9}
		repeatedNumber      = [...]rune{-1, 0}
		iterations          = 0
	)

	for _, character := range element {
		if character >= '0' && character <= '9' {
			actualNumber := character - '0'
			if iterations < LengthCNPJ-2 {
				if repeatedNumber[0] == -1 {
					// getting first number
					repeatedNumber[0] = character
				}

				if repeatedNumber[0] == character {
					// counting number repetitions
					repeatedNumber[1]++
				}
				if subscriptionWeights[0] >= 2 {
					verifierDigits[0] += actualNumber * subscriptionWeights[0]
					subscriptionWeights[0]--
				} else {
					verifierDigits[0] += actualNumber * branchWeights[0]
					branchWeights[0]--
				}
				if subscriptionWeights[1] >= 2 {
					verifierDigits[1] += actualNumber * subscriptionWeights[1]
					subscriptionWeights[1]--
				} else {
					verifierDigits[1] += actualNumber * branchWeights[1]
					branchWeights[1]--
				}
			} else {
				originalVerifier = append(originalVerifier, actualNumber)
			}
			iterations++
		}
	}

	if verifierDigits[0] = 11 - (verifierDigits[0] % 11); verifierDigits[0] >= 10 {
		verifierDigits[0] = 0
	}
	verifierDigits[1] += verifierDigits[0] * branchWeights[1]

	if verifierDigits[1] = 11 - (verifierDigits[1] % 11); verifierDigits[1] >= 10 {
		verifierDigits[1] = 0
	}

	hasCorrectLength := repeatedNumber[1] < LengthCNPJ-2 && iterations <= LengthCNPJ && len(originalVerifier) == 2

	if hasCorrectLength {
		hasCorrectDigits := originalVerifier[0] == verifierDigits[0] && originalVerifier[1] == verifierDigits[1]
		if !hasCorrectDigits {
			result = utils.ErrInvalidElement
		}
	} else {
		result = utils.ErrElementIncorrectLength
	}

	return
}
