//go:build amd64 && !gccgo
// +build amd64,!gccgo

// Package base64 is an efficient base64 implementation for Golang.
package base64

import (
	"errors"
	"strconv"
)

type EncodingType int

const (
	EncodeStd EncodingType = iota
	EncodeURL
)

const (
	StdPadding rune = '=' // Standard padding character
	NoPadding  rune = -1  // No padding
)

var (
	StdEncoding = NewEncoding(EncodeStd)
	URLEncoding = NewEncoding(EncodeURL)

	RawStdEncoding = StdEncoding.WithPadding(NoPadding)
	RawURLEncoding = URLEncoding.WithPadding(NoPadding)
)

var ErrFormat = errors.New("base64: invalid input")

type CorruptInputError int64

func (e CorruptInputError) Error() string {
	return "illegal base64 data at input byte " + strconv.FormatInt(int64(e), 10)
}
