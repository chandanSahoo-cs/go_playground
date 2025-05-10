/*
Implement the isValidPassword function by looping through each character in the password string. Make sure the password is long enough and includes at least one uppercase letter and one digit.

Assume passwords consist of ASCII characters only.
*/


package main

import (
	"fmt"
	"unicode"
)


func isValidPassword(password string) bool {
	
	if len(password)<5 || len(password)>12{
		return false;
	}

	uppercaseLetter :=false
	digit :=false

	for _,ch := range(password) {
		if(unicode.IsDigit(ch)){
			digit = true
		}else if(unicode.IsUpper(ch)){
			uppercaseLetter = true
		}
	}

	if uppercaseLetter && digit {
		return true;
	}
	return false
}

func main() {
	fmt.Println(isValidPassword("Hello1"))     // true
	fmt.Println(isValidPassword("hello1"))     // false (no uppercase)
	fmt.Println(isValidPassword("HELLO"))      // false (no digit)
	fmt.Println(isValidPassword("H1"))         // false (too short)
	fmt.Println(isValidPassword("Hello1234567")) // false (too long)
}

/*
Sample Output:
true
false
false
false
false
*/
