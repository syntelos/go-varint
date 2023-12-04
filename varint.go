/*
 * ULEB64 VARINT I/O
 * Copyright 2023 John Douglas Pritchard, Syntelos
 *
 *
 * References
 *
 * https://en.wikipedia.org/wiki/LEB128
 */
package varint

import (
	"fmt"
	"io"
	"math"
	"math/bits"
)
/*
 * Size encoded in octets.
 */
func Count(src uint64) (ct uint64){
	var b float64 = float64(bits.Len64(src))
	var c float64 = (b / 7.0)
	ct = uint64(math.Ceil(c))
	return ct
}
/*
 * Byte buffer writer.
 */
func Write(w io.ByteWriter, src uint64) (ct uint64, er error) {
	ct = 0

	var x int
	var b uint64 = src
	var e byte

	for x = 0; x <= 8; x++ {

		e = byte(b & 0x7F)

		b >>= 7

		if 0 != b {
			w.WriteByte(e | 0x80)
			ct += 1
		} else {
			w.WriteByte(e)
			ct += 1
			break
		}
	}

	return ct, nil
}
/*
 * Byte buffer reader.
 */
func Read(r io.ByteReader) (dst uint64, er error) {
	dst = 0

	var x int
	var b byte
	var s uint8 = 0

	for x = 0; x < 8; x++ {

		b, er = r.ReadByte()
		if nil != er {

			if io.EOF == er {

				return dst, nil
			} else {
				return 0, fmt.Errorf("Reading varint: %w",er)
			}
		} else {
			dst |= (uint64(b & 0x7F)<< s)
			s += 7

			if 0x80 != (b & 0x80) {

				break
			}
		}
	}
	return dst, nil
}
