package iteration

import "strings"

func Repeat(char string) string {
	// go'da immutable olduğu için stringler her loopta memoryde yeni yer açar.
	// direkt bunu concantenate yaparken
	//  memoryde dinamik yer açıp stringleri ekliyor
	var myString strings.Builder
	for i := 0; i < 4; i++ {
		myString.WriteString(char)
	}
	return myString.String()
}
