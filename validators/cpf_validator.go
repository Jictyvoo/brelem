package validators

import "github.com/jictyvoo/brelem/utils"

// LengthCPF The CPF Length set to 11
const LengthCPF = 11

func CPF(element string) (result error) {
	var (
		verifierDigits   = [...]rune{0, 0}
		originalVerifier = make([]rune, 0, 5)
		weights          = [...]rune{10, 11}
		repeatedNumber   = [...]rune{-1, 0}
		iterations       = 0
	)

	for _, character := range element {
		if character >= '0' && character <= '9' {
			actualNumber := character - '0'
			if iterations < LengthCPF-2 {
				{
					verifierDigits[0] += actualNumber * weights[0]
					verifierDigits[1] += actualNumber * weights[1]
					weights[0]--
					weights[1]--
				}

				if repeatedNumber[0] == -1 {
					// getting first number
					repeatedNumber[0] = character
				}
				if repeatedNumber[0] == character {
					// counting number repetitions
					repeatedNumber[1]++
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
	verifierDigits[1] += verifierDigits[0] * weights[1]

	if verifierDigits[1] = 11 - (verifierDigits[1] % 11); verifierDigits[1] >= 10 {
		verifierDigits[1] = 0
	}

	hasCorrectLength := repeatedNumber[1] < LengthCPF-2 && iterations <= LengthCPF && len(originalVerifier) == 2
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
