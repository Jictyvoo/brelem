package validators

import "testing"

func _getValidCNH() [14]string {
	return [...]string{
		"22522791500", "46613298880", "18271939762", "03621746707", "26606066255", "33235836407", "35318223990",
		"96195156706", "33004095866", "87672878823", "33282790801", "29231325340", "62504334773", "37882306567",
	}
}

func TestCNH(t *testing.T) {
	testCases := _getValidCNH()
	for _, testCase := range testCases {
		resultErr := CNH(testCase)
		if resultErr != nil {
			t.Error(resultErr)
		}
	}
}
