// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"
)

// BookType is the 'book_type' enum type from schema 'public'.
type BookType uint16

const (
	// FictionBookType is the 'FICTION' BookType.
	FictionBookType = BookType(1)

	// NonfictionBookType is the 'NONFICTION' BookType.
	NonfictionBookType = BookType(2)
)

// String returns the string value of the BookType.
func (bt BookType) String() string {
	var enumVal string

	switch bt {
	case FictionBookType:
		enumVal = "FICTION"

	case NonfictionBookType:
		enumVal = "NONFICTION"
	}

	return enumVal
}

// MarshalText marshals BookType into text.
func (bt BookType) MarshalText() ([]byte, error) {
	return []byte(bt.String()), nil
}

// UnmarshalText unmarshals BookType from text.
func (bt *BookType) UnmarshalText(text []byte) error {
	switch string(text) {
	case "FICTION":
		*bt = FictionBookType

	case "NONFICTION":
		*bt = NonfictionBookType

	default:
		return errors.New("invalid BookType")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for BookType.
func (bt BookType) Value() (driver.Value, error) {
	return bt.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for BookType.
func (bt *BookType) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid BookType")
	}

	return bt.UnmarshalText(buf)
}
