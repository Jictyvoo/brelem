package validators

import "testing"

func _getValidDocuments() [11]string {
	return [...]string{
		// CPF
		"298.074.850-10", "241.525.560-21", "501.585.380-72", "97949526050", "43793938018",

		// CNPJ
		"54.942.449/0001-06", "38.904.678/0001-71", "71.581.155/0001-07", "30.591.729/0001-40",
		"90718654000148", "73329532000140",
	}

}

func TestCPFCNPJ(t *testing.T) {
	testCases := _getValidDocuments()
	for _, testCase := range testCases {
		resultErr := CPFCNPJ(testCase)
		if resultErr != nil {
			t.Error(resultErr)
		}
	}
}

func TestValidateDetermineCPFCNPJ(t *testing.T) {
	t.Run(
		"Determine CPF Test", func(cpfT *testing.T) {
			correctCPFs := [...]string{
				"298.074.850-10", "241.525.560-21", "501.585.380-72", "97949526050", "43793938018",
			}
			for _, cpf := range correctCPFs {
				eType, err := DetermineCPFCNPJ(cpf)
				if err != nil {
					t.Error(err)
				}
				if eType != TypeCPF {
					t.Errorf("Wrong element type returned: `%v`", eType)
				}
			}
		},
	)

	t.Run(
		"Determine CNPJ Test", func(cpfT *testing.T) {
			correctCNPJs := [...]string{
				"54.942.449/0001-06", "38.904.678/0001-71", "71.581.155/0001-07", "30.591.729/0001-40",
				"90718654000148", "73329532000140",
			}
			for _, cnpj := range correctCNPJs {
				eType, err := DetermineCPFCNPJ(cnpj)
				if err != nil {
					t.Error(err)
				}
				if eType != TypeCNPJ {
					t.Errorf("Wrong element type returned: `%v`", eType)
				}
			}
		},
	)
}

func TestValidateDetermineCPFCNPJ_IncorrectValues(t *testing.T) {
	testCases := [...]string{
		// CPF
		"000.000.000-00", "111.111.111-11", "222.222.222-22",
		"33333333333", "44444444444", "55555555555", "66666666666", "77777777777", "88888888888", "99999999999",
		// CPF - Incomplete
		"111.111.", "46477108071", "", "12",
		// Correct but with more digits
		"241.525.560-210", "298.074.850-1034", "000.000.001-00",

		// CNPJ
		"00.000.000/0000-00", "11.111.111/1111-11", "22.222.222/2222-22", "33.333.333/3333-33", "44.444.444/4444-44",
		"55.555.555/5555-55", "66.666.666/6666-66", "77.777.777/7777-77", "88.888.888/8888-88", "99.999.999/9999-99",
		"69.372.070/0001-", "10004621459503", "00.000.000/0001-00",
	}

	for _, testCase := range testCases {
		eType, resultErr := DetermineCPFCNPJ(testCase)
		if resultErr == nil {
			t.Errorf("It's supposed to fail in validation, but it validates the `%s` as correct", testCase)
		}
		if eType != TypeUNKNOWN {
			t.Errorf("Wrong element type returned: `%v`", eType)
		}
	}
}

func BenchmarkCPFCNPJ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, element := range _getValidDocuments() {
			if err := CPFCNPJ(element); err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkAsyncCPFCNPJ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, element := range _getValidDocuments() {
			if err := AsyncCPFCNPJ(element); err != nil {
				b.Error(err)
			}
		}
	}
}
