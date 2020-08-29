// Copyright 2020 The go-benchmarks Authors.
// SPDX-License-Identifier: BSD-3-Clause

package pseudo

import (
	_ "runtime" // used with go:linkname
	_ "unsafe"  // required for go:linkname
)

//go:linkname fastrandn runtime.fastrandn
func fastrandn(n uint32) uint32
