package strand

// ToRNA converts DNA strands into RNA strands
func ToRNA(dna string) string {
	if len(dna) == 0 {
		return ""
	}

	dnaToRna := map[byte]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}

	rna := ""
	for i := 0; i < len(dna); i++ {
		strand := dna[i]
		rna += dnaToRna[strand]
	}

	return rna
}
