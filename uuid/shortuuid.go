// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"strings"

	"github.com/google/uuid"
	"github.com/mr-tron/base58"
)

var Alphabet = base58.NewAlphabet("123456789PaymentABCDEFGHJKLMNQRSTUVWXYZbcdfghijkopqrsuvwxz")
var ZeroAlpha = "1"

type Base58Encoder struct{}

func (enc Base58Encoder) Encode(u uuid.UUID) string {
	return base58.EncodeAlphabet(u[:], Alphabet)
}

func (enc Base58Encoder) Decode(s string) (uuid.UUID, error) {
	var uid uuid.UUID
	b, err := base58.DecodeAlphabet(s, Alphabet)
	if err != nil {
		return uid, err
	}

	return uuid.FromBytes(b)
}

const IDWidth = 22 // uuid v4 -> encode58 max width

func New() string {
	uid := uuid.New()
	var encoder = Base58Encoder{}
	shortuid := encoder.Encode(uid)

	// padding zero value - uuidencoder doesn't pad leading zero.
	remainder := IDWidth - len(shortuid)
	if remainder > 0 {
		shortuid = shortuid + strings.Repeat(ZeroAlpha, remainder)
	}

	return shortuid
}
