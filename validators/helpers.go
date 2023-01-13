package validators

type ElementType uint8

const (
	TypeUNKNOWN ElementType = iota
	TypeCPF
	TypeCNPJ
)

func _modVerifierDigit(toVerify rune) rune {
	toVerify = 11 - (toVerify % 11)
	if toVerify >= 10 {
		toVerify = 0
	}
	return toVerify
}
