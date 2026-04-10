package brdoc

import "testing"

func TestDetermineCPFCNPJ_CPF(t *testing.T) {
	cpfs := [...]string{
		"298.074.850-10", "241.525.560-21", "501.585.380-72", "97949526050", "43793938018",
	}
	for _, input := range cpfs {
		t.Run(input, func(t *testing.T) {
			result, err := DetermineCPFCNPJ(input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Type != TypeCPF {
				t.Errorf("got type %v, want TypeCPF", result.Type)
			}
			if result.CPF.IsZero() {
				t.Error("CPF should not be zero")
			}
		})
	}
}

func TestDetermineCPFCNPJ_CNPJ(t *testing.T) {
	cnpjs := [...]string{
		"54.942.449/0001-06", "38.904.678/0001-71", "71.581.155/0001-07",
		"30.591.729/0001-40", "90718654000148", "73329532000140",
	}
	for _, input := range cnpjs {
		t.Run(input, func(t *testing.T) {
			result, err := DetermineCPFCNPJ(input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Type != TypeCNPJ {
				t.Errorf("got type %v, want TypeCNPJ", result.Type)
			}
			if result.CNPJ.IsZero() {
				t.Error("CNPJ should not be zero")
			}
		})
	}
}

func TestDetermineCPFCNPJ_Invalid(t *testing.T) {
	tests := [...]string{
		"000.000.000-00", "111.111.111-11", "222.222.222-22",
		"33333333333", "44444444444", "55555555555", "66666666666", "77777777777", "88888888888", "99999999999",
		"111.111.", "46477108071", "", "12",
		"241.525.560-210", "298.074.850-1034", "000.000.001-00",
		"00.000.000/0000-00", "11.111.111/1111-11", "22.222.222/2222-22", "33.333.333/3333-33", "44.444.444/4444-44",
		"55.555.555/5555-55", "66.666.666/6666-66", "77.777.777/7777-77", "88.888.888/8888-88", "99.999.999/9999-99",
		"69.372.070/0001-", "10004621459503", "00.000.000/0001-00",
	}

	for _, input := range tests {
		name := input
		if name == "" {
			name = "empty"
		}
		t.Run(name, func(t *testing.T) {
			result, err := DetermineCPFCNPJ(input)
			if err == nil {
				t.Errorf("expected error for %q, got nil", input)
			}
			if result.Type != TypeUnknown {
				t.Errorf("got type %v, want TypeUnknown", result.Type)
			}
		})
	}
}

func BenchmarkDetermineCPFCNPJ(b *testing.B) {
	b.ReportAllocs()
	inputs := [...]string{
		"298.074.850-10", "241.525.560-21", "501.585.380-72", "97949526050", "43793938018",
		"54.942.449/0001-06", "38.904.678/0001-71", "71.581.155/0001-07",
		"30.591.729/0001-40", "90718654000148", "73329532000140",
	}
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			if _, err := DetermineCPFCNPJ(input); err != nil {
				b.Fatal(err)
			}
		}
	}
}
