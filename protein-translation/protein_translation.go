package protein

import (
	"errors"
	"strings"
)

var (
	// ErrStop denotes the end codon.
	ErrStop = errors.New("errStop")
	// ErrInvalidBase denotes an invalid codon base.
	ErrInvalidBase = errors.New("invalid base")
)
var codonMapping = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromRNA calculate the protein sequence.
func FromRNA(rna string) ([]string, error) {
	var protein []string
	for i := 0; i < len(rna); i = i + 3 {
		codon, err := FromCodon(rna[i : i+3])
		if err == ErrInvalidBase {
			return protein, ErrInvalidBase
		}
		if err == ErrStop {
			break
		}
		protein = append(protein, codon)
	}
	return protein, nil
}

// FromCodon calculate the protein sequence.
func FromCodon(rna string) (string, error) {
	// if rna == "" {
	// 	return "", nil
	// }
	mapping, hasItem := codonMapping[strings.ToUpper(rna)]
	if !hasItem {
		return "", ErrInvalidBase
	}
	if mapping == "STOP" {
		return "", ErrStop
	}
	return mapping, nil
}
