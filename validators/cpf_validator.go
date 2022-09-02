package validators

import "github.com/jictyvoo/brelem/utils"

// LengthCPF The CPF Length set to 11
const LengthCPF = 11

type cpfValidator struct {
	verifierDigits   [2]rune
	originalVerifier []rune
	weights          [2]rune
	repeatedNumber   [2]rune
	iterations       int
}

func newCpfValidator() cpfValidator {
	return cpfValidator{
		verifierDigits:   [...]rune{0, 0},
		originalVerifier: make([]rune, 0, 2),
		weights:          [...]rune{10, 11},
		repeatedNumber:   [...]rune{-1, 0},
		iterations:       0,
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
		v.originalVerifier = append(v.originalVerifier, intChar-'0')
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
	return v.repeatedNumber[1] >= LengthCPF-2 || v.iterations > LengthCPF || len(v.originalVerifier) != 2
}

func (v *cpfValidator) HasValidDigits() bool {
	return v.originalVerifier[0] == v.verifierDigits[0] && v.originalVerifier[1] == v.verifierDigits[1]
}

// AsyncCPF Verify CPF string, remove non-numeric characters, calculate verifier digit
func AsyncCPF(cpfString chan rune) (result chan error) {
	result = make(chan error, 1)
	go ChannelCheckCPF(cpfString, result)

	return
}

func ChannelCheckCPF(cpfString chan rune, result chan error) {
	defer close(result)
	validator := newCpfValidator()

	for intChar := range cpfString {
		validator.iterateRune(intChar)
	}
	if err := validator.finishValidation(); err != nil {
		result <- err
	}
}
