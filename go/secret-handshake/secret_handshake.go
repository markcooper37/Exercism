package secret

var action = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) []string {
	handshake := []string{}
	for i := 0; i <= 3; i++ {
		if code % 2 == 1 {
			handshake = append(handshake, action[i])
		}
		code /= 2
	}
	if code % 2 == 1 {
		for i := 0; i < len(handshake)/2; i++ {
			handshake[i], handshake[len(handshake)-i-1] = handshake[len(handshake)-i-1], handshake[i]
		}
	}
	return handshake
}
