package day_6

func DNACheckValue(dna string) (DNA, string) {
	result := DNA{
		A: 0,
		C: 0,
		G: 0,
		T: 0,
	}
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'A':
			result.A += 1
			break
		case 'C':
			result.C += 1
			break
		case 'G':
			result.G += 1
			break
		case 'T':
			result.T += 1
			break
		default:
			return result, "error"
		}
	}
	return result, ""
}
