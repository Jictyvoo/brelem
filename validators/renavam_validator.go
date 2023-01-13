package validators

import "log"

// Renavam obtained on https://gist.github.com/cagartner/efe5e37c9c52063660cd
func Renavam(element string) (result error) {
	for _, character := range element {
		if character >= '0' && character <= '9' {
			actualNumber := character - '0'
			log.Println(actualNumber)
		}
	}

	return
}
