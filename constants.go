package gobioseq

import (
	"strings"
	"strconv"
)

var Bases = [4]string{ "T", "C", "A", "G" }
var Codons = [64]string{}
i = 0
for _, a := range Bases {
	for _, b := range Bases {
		for _, c := range Bases {
			Codons[i] = a + b + c 
			i++
		}
	}
}
var StopCodons = [3]string{ "TGA", "TAG", "TAA" }
var AminoAcids = [20]string{
	"A",
	"R",
	"N",
	"D",
	"C",
	"Q",
	"E",
	"G",
	"H",
	"I",
	"L",
	"K",
	"M",
	"F",
	"P",
	"S",
	"T",
	"W",
	"Y",
	"V",
}

var translCode = "F2L6I3M1V4S4P4T4A4Y2*2H2Q2N2K2D2E2C2*1W1R4S2R2G4"
var transl []string
for i := 0; i < len(translCode); i += 2 {
	for j := 1; j < len(translCode); j += 2 {
		times, _ := strconv.Atoi(translCode[j])
		transl = append(transl, strings.Repeat(translCode[i], times)) 
	}
}
var GeneticCode = make(map[string]string)
for i := range Codons {
	GeneticCode[Codons[i]] = transl[i]
}

var DegenerateBases = map[string][]string{
	"W": []string{ "A", "T" },  // 2-FOLD
	"S": []string{ "C", "G" },
	"M": []string{ "A", "C" },
	"K": []string{ "G", "T" },
	"R": []string{ "A", "G" },
	"Y": []string{ "C", "T" },
	"B": []string{ "C", "G", "T" },  // 3-FOLD
	"D": []string{ "A", "G", "T" },
	"H": []string{ "A", "C", "T" },
	"V": []string{ "A", "C", "G" },
	"N": []string{ "A", "C", "G", "T" },  // 4-FOLD
}

var AAFold = make(map[string]int)
for _, v := range GeneticCode {
	AAFold[v]++
}

var CodonFold = make(map[string]int)
for k, v := range GeneticCode {
	CodonFold[k] : AAFold[v]
}

type MSA struct {
	ids []string
	alignment [][]string
}

var SequenceTypes = [3]string{ "nucl", "prot", "cod" }