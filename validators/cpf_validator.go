package validators

import "github.com/jictyvoo/brelem/utils"

// LengthCPF The CPF Length set to 11
const LengthCPF = 11

// validateAsyncCPF Verify CPF string, remove non-numeric characters, calculate verifier digit
func validateAsyncCPF(cpfString chan rune) (result chan error) {
	result = make(chan error, 1)
	go ValidateCPF(cpfString, result)

	return
}

func ValidateCPF(cpfString chan rune, result chan error) {
	defer close(result)
	var (
		verifierDigits   = [...]rune{0, 0}
		originalVerifier = make([]rune, 0, 2)
		weights          = [...]rune{10, 11}
		repeatedNumber   = [...]rune{-1, 0}
		iterations       = 0
	)

	for intChar := range cpfString {
		if intChar >= '0' && intChar <= '9' {
			if iterations < LengthCPF-2 {
				{
					verifierDigits[0] += (intChar - '0') * weights[0]
					verifierDigits[1] += (intChar - '0') * weights[1]
					weights[0]--
					weights[1]--
				}

				if repeatedNumber[0] == -1 {
					// getting first number
					repeatedNumber[0] = intChar
				}
				if repeatedNumber[0] == intChar {
					// counting number repetitions
					repeatedNumber[1]++
				}
			} else {
				originalVerifier = append(originalVerifier, intChar-'0')
			}
			iterations++
		}
	}
	verifierDigits[0] = _modVerifierDigit(verifierDigits[0])
	verifierDigits[1] += verifierDigits[0] * weights[1]
	verifierDigits[1] = _modVerifierDigit(verifierDigits[1])

	if repeatedNumber[1] >= LengthCPF-2 || iterations > LengthCPF || len(originalVerifier) != 2 {
		result <- utils.ErrElementIncorrectLength
	} else if originalVerifier[0] != verifierDigits[0] || originalVerifier[1] != verifierDigits[1] {
		result <- utils.ErrInvalidElement
	}
}
