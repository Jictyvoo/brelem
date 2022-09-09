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
		if iterations < LengthCPF-2 {
			actualNumber := character - '0'
			{
				verifierDigits[0] += (actualNumber) * weights[0]
				verifierDigits[1] += (actualNumber) * weights[1]
				weights[0]--
				weights[1]++
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
			originalVerifier = append(originalVerifier, character-'0')
		}
		iterations++
	}

	// start checking the digits
	var shouldRecalculate bool
	if verifierDigits[0] %= 11; verifierDigits[0] > 9 {
		verifierDigits[0] = 0
		shouldRecalculate = true
	}

	verifierDigits[1] %= 11
	if shouldRecalculate {
		checkValue := verifierDigits[1] - 2
		if checkValue < 0 {
			verifierDigits[1] += 9
		} else {
			verifierDigits[1] -= 2
		}
	}
	// final adjust on second digit
	if verifierDigits[1] > 9 {
		verifierDigits[1] = 0
	}

	hasIncorrectLength := repeatedNumber[1] >= LengthCNH-2 || iterations > LengthCNH || len(originalVerifier) != 2
	hasValidDigits := originalVerifier[0] == verifierDigits[0] && originalVerifier[1] == verifierDigits[1]
	if hasIncorrectLength {
		result = utils.ErrElementIncorrectLength
	} else if !hasValidDigits {
		result = utils.ErrInvalidElement
	}

	return
}
