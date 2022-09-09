package validators

import "github.com/jictyvoo/brelem/utils"

// LengthCNH The CNH Length set to 11
const LengthCNH = 11

func CNH(element string) (result error) {
	var (
		iterations       uint
		verifierDigits   [2]rune
		originalVerifier = make([]rune, 0, 2)
		weights          = [...]rune{9, 1}
		repeatedNumber   = [...]rune{-1, 0}
	)

	for _, character := range element {
		if character >= '0' && character <= '9' {
			actualNumber := character - '0'
			if iterations < LengthCPF-2 {
				{
					verifierDigits[0] += (actualNumber) * weights[0]
					verifierDigits[1] += (actualNumber) * weights[1]
					weights[1]++
					weights[0]--
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

	// start checking the digits
	var shouldRecalculate bool
	if verifierDigits[0] %= 11; verifierDigits[0] >= 10 {
		verifierDigits[0] = 0
		shouldRecalculate = true
	}

	if verifierDigits[1] %= 11; verifierDigits[1] >= 10 {
		verifierDigits[1] = 0
	} else if shouldRecalculate {
		verifierDigits[1] -= 2
	}

	hasCorrectLength := repeatedNumber[1] < LengthCNH-2 && iterations <= LengthCNH && len(originalVerifier) == 2
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
