package rotationalcipher

import "fmt"

func RotationalCipher(plain string, shiftKey int) string {
	cipher := ""
    for _, c := range plain {
        if c >= 65 && c <= 90 {
            value := int(c) + shiftKey
            if value > 90 {
                value -= 26
            }
        	cipher = fmt.Sprintf("%s%c", cipher, value)
        } else if c >= 97 && c <= 122 {
        	value := int(c) + shiftKey 
            if value > 122 {
                value -= 26
            }
        	cipher = fmt.Sprintf("%s%c", cipher, value)
        } else {
        	cipher = fmt.Sprintf("%s%c", cipher, c)
        }
    }
	return cipher
}
