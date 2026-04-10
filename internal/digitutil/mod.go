package digitutil

// ModVerifierDigit applies the modulo-11 check digit calculation.
// Returns 11 - (toVerify % 11), clamped to 0 when the result is >= 10.
func ModVerifierDigit(toVerify rune) rune {
	toVerify = 11 - (toVerify % 11)
	if toVerify >= 10 {
		toVerify = 0
	}
	return toVerify
}
