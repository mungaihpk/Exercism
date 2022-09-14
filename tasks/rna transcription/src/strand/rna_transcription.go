package strand

import (
	"strings"
)

var dna_nucleotide_mapping = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

func ToRNA(dna string) string {
	//panic("Please implement the ToRNA function")
	var rna string

	for _, nucleotide := range strings.Split(dna, "") {
		rna += dna_nucleotide_mapping[nucleotide]
	}
	return rna
}
