package variablelengthquantity

import "errors"

func EncodeVarint(input []uint32) []byte {
	output := []byte{}
	for _, integer := range input {
		encoded := []byte{byte(integer % 128)}
		integer /= 128
		for integer > 0 {
			encoded = append([]byte{byte((integer % 128) + 128)}, encoded...)
			integer /= 128
		}
		output = append(output, encoded...)
	}
	return output
}

func DecodeVarint(input []byte) ([]uint32, error) {
	output := []uint32{}
	integer := uint32(0)
	completeSequences := true
	for _, inputByte := range input {
		integer <<= 7
		if inputByte >= 128 {
			integer += uint32(inputByte - 128)
			completeSequences = false
		} else {
			integer += uint32(inputByte)
			output = append(output, integer)
			integer = 0
			completeSequences = true
		}
	}
	if !completeSequences {
		return nil, errors.New("incomplete sequence")
	}
	return output, nil
}
