// Copyright 2021 Faye Amacker
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package circlehash

import (
	"unsafe"
)

func circle64f(p unsafe.Pointer, seed uint64, dlen uint64) uint64 {

	startingLength := dlen
	currentState := seed ^ pi0

	if dlen > 64 {
		// Process chunks of 64 bytes.
		duplicatedState := currentState

		for ; dlen > 64; dlen -= 64 {
			a := readUnaligned64(p)
			b := readUnaligned64(add(p, 8))
			c := readUnaligned64(add(p, 16))
			d := readUnaligned64(add(p, 24))
			e := readUnaligned64(add(p, 32))
			f := readUnaligned64(add(p, 40))
			g := readUnaligned64(add(p, 48))
			h := readUnaligned64(add(p, 56))

			cs0 := mix64(a^pi1, b^currentState)
			cs1 := mix64(c^pi2, d^currentState)
			currentState = (cs0 ^ cs1)

			ds0 := mix64(e^pi3, f^duplicatedState)
			ds1 := mix64(g^pi4, h^duplicatedState)
			duplicatedState = (ds0 ^ ds1)

			p = add(p, 64)
		}

		currentState = currentState ^ duplicatedState
	}

	// We have at most 64 bytes to process.
	// Process chunks of 16 bytes
	for ; dlen > 16; dlen -= 16 {
		a := readUnaligned64(p)
		b := readUnaligned64(add(p, 8))

		currentState = mix64(a^pi1, b^currentState)

		p = add(p, 16)
	}

	// We have at most 16 bytes to process.
	a := uint64(0)
	b := uint64(0)

	switch dlen {

	case 9, 10, 11, 12, 13, 14, 15, 16:
		// We have 9-16 bytes to process.
		// a and b might overlap.
		a = readUnaligned64(p)
		b = readUnaligned64(add(p, uintptr(dlen-8)))

	case 4, 5, 6, 7, 8:
		// We have 4-8 bytes to process.
		// a and b might overlap.
		a = uint64(readUnaligned32(p))
		b = uint64(readUnaligned32(add(p, uintptr(dlen-4))))

	case 1, 2, 3:
		// We have 1-3 bytes to process.
		a = uint64(*(*byte)(p)) << 16
		a |= uint64(*(*byte)(add(p, uintptr(dlen>>1)))) << 8
		a |= uint64(*(*byte)(add(p, uintptr(dlen-1))))
		b = 0

	case 0:
		a = 0
		b = 0
	}

	w := mix64(a^pi1, b^currentState)
	z := pi4 ^ startingLength // wyhash reuses salt1 here, but CircleHash64 (like Go 1.17) avoids reusing it here
	return mix64(w, z)
}

// circle64fUint64x2 produces a 64-bit digest from a, b, and seed.
// Digest is compatible with circlehash64f with byte slice of len 16.
func circle64fUint64x2(a uint64, b uint64, seed uint64) uint64 {
	const dataLen = uint64(16)
	currentState := seed ^ pi0
	w := mix64(a^pi1, b^currentState)
	z := pi4 ^ dataLen
	return mix64(w, z)
}
