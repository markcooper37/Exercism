package strand

import "fmt"

func ToRNA(dna string) string {
	RNA := ""
    nucleotideMap := map[rune]string{'A': "U", 'C': "G", 'G': "C", 'T': "A"}
    for _, char := range dna {
        RNA = fmt.Sprintf("%s%s", RNA, nucleotideMap[char])
    }
	return RNA
}
