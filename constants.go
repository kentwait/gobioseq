package gobioseq

import (
	"strings"
	"strconv"
)

// Bases lists the allowed nucleotide base characters in the SequenceArrays or SequenceAlignments.
var Bases = [4]string{ "T", "C", "A", "G" }

// Codons lists the allowed codon strings in CodonAlignments.
var Codons = [64]string{}
var i = 0
for _, a := range Bases {
	for _, b := range Bases {
		for _, c := range Bases {
			Codons[i] = a + b + c 
			i++
		}
	}
}

// StopCodons lists the codons that yield to a termination signal.
var StopCodons = [3]string{ "TGA", "TAG", "TAA" }

// AminoAcids lists the allowed single-character amino acid representations in SequenceArrays or SequenceAlignments.
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

// DegenerateBases lists IUPAC characters to represent degenerate bases and the nucleotide bases they represent.
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

// AAFold lists the number of different codons that code for a particular amino acid.
var AAFold = make(map[string]int)
for _, v := range GeneticCode {
	AAFold[v]++
}

// CodonFold lists which fold category a particular codon belongs to. 
// This count is based on the number of different codons that result in particular amino acid.
var CodonFold = make(map[string]int)
for k, v := range GeneticCode {
	CodonFold[k] : AAFold[v]
}

// MSA is a struct for a multiple sequence alignment datatype
type MSA struct {
	ids []string
	alignment [][]string
}

// SequenceTypes lists the possible sequence types in a SequenceArray or SequenceAlignment
var SequenceTypes = [3]string{ "nucl", "prot", "cod" }
