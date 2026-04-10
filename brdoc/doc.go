// Package brdoc provides Brazilian document types (CPF, CNPJ, CNH) with
// validation on creation, formatting, and masking.
package brdoc

// Document is the common interface for Brazilian document types.
type Document interface {
	// String returns the human-readable formatted representation.
	String() string
	// Digits returns the raw numeric string with no formatting.
	Digits() string
	// VerifierDigits returns the two check digits as numeric values (0-9).
	VerifierDigits() [2]uint8
	// IsZero reports whether the document is the zero value (not constructed).
	IsZero() bool
}

// Compile-time interface checks.
var (
	_ Document = CPF{}
	_ Document = CNPJ{}
	_ Document = CNH{}
)
