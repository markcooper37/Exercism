package variablelengthquantity

func EncodeVarint(input []uint32) []byte {
	output := []byte{}
	for _, integer := range input {
		encoded := []byte{byte(integer % 128)}
		integer /= 128
		for integer > 0 {
			encoded = append([]byte{byte((integer % 128)+128)}, encoded...)
			integer /= 128
		}
		output = append(output, encoded...)
	}
	return output
}

func DecodeVarint(input []byte) ([]uint32, error) {
	return []uint32{}, nil
}
