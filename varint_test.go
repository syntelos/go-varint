/*
 * ULEB64 VARINT I/O
 * Copyright 2023 John Douglas Pritchard, Syntelos
 */
package varint

import (
	"bytes"
	"encoding/binary"
	"testing"
	"github.com/syntelos/go-endian"
)
/*
 * Self consistency
 */
func Test00E7(t *testing.T){
	const value uint64 = 0x00E7
	var count uint64 = Count(value)
	var buf *bytes.Buffer = new(bytes.Buffer)

	sz, er := Write(buf,value)
	if nil != er {
		t.Errorf("Error writing value 0x%X: %v",value,er)
	} else if count != sz {
		t.Errorf("Error writing value 0x%X.  Expected length (%d) found length (%d).",value,count,sz)
	} else {
		var read uint64
		read, er = Read(buf)
		if nil != er {
			t.Errorf("Error reading value 0x%X: %v",value,er)
		} else if value != read {
			t.Errorf("Error reading.  Expected value (0x%X) found value (0x%X).",value,read)
		}
	}
}
/*
 * Self consistency
 */
func Test1200(t *testing.T){
	const value uint64 = 0x1200
	var count uint64 = Count(value)
	var buf *bytes.Buffer = new(bytes.Buffer)

	sz, er := Write(buf,value)
	if nil != er {
		t.Errorf("Error writing value 0x%X: %v",value,er)
	} else if count != sz {
		t.Errorf("Error writing value 0x%X.  Expected length (%d) found length (%d).",value,count,sz)
	} else {
		var read uint64
		read, er = Read(buf)
		if nil != er {
			t.Errorf("Error reading value 0x%X: %v",value,er)
		} else if value != read {
			t.Errorf("Error reading.  Expected value (0x%X) found value (0x%X).",value,read)
		}
	}
}
/*
 * Self consistency
 */
func TestLargeNumbers(t *testing.T){
	var sequence uint8
	for sequence = 1; sequence <= 16; sequence++ {

		var er error

		var value uint64
		value, er = endian.Rand64()

		if nil != er {
			t.Error(er)
		} else {
			var count uint64 = Count(value)
			var buf *bytes.Buffer = new(bytes.Buffer)

			sz, er := Write(buf,value)
			if nil != er {
				t.Errorf("Error writing value 0x%X: %v",value,er)
			} else if count != sz {
				t.Errorf("Error writing value 0x%X.  Expected length (%d) found length (%d).",value,count,sz)
			} else {
				var read uint64
				read, er = Read(buf)
				if nil != er {
					t.Errorf("Error reading value 0x%X: %v",value,er)
				} else if value != read {
					t.Errorf("Error reading.  Expected value (0x%X) found value (0x%X).",value,read)
				}
			}
		}
	}
}
/*
 * GOPL consistency
 */
func TestGoVintGet(t *testing.T){
	var sequence uint8
	for sequence = 1; sequence <= 16; sequence++ {

		var er error

		var value uint64
		value, er = endian.Rand64() 
		if nil != er {
			t.Error(er)
		} else {
			var count uint64 = Count(value)
			var buf *bytes.Buffer = new(bytes.Buffer)

			sz, er := Write(buf,value)
			if nil != er {
				t.Errorf("Error writing value 0x%X: %v",value,er)
			} else if count != sz {
				t.Errorf("Error writing value 0x%X.  Expected length (%d) found length (%d).",value,count,sz)
			} else {
				var read uint64
				var vz int
				read, vz = binary.Uvarint(buf.Bytes())
				if 1 > vz {
					t.Errorf("Error reading value 0x%X",value)
				} else if value != read {
					t.Errorf("Error reading.  Expected value (0x%X) found value (0x%X).",value,read)
				}
			}
		}
	}
}
/*
 * GOPL consistency
 */
func TestGoVintPut(t *testing.T){
	var sequence uint8
	for sequence = 1; sequence <= 16; sequence++ {

		var er error

		var value uint64
		value, er = endian.Rand64() 
		if nil != er {
			t.Error(er)
		} else {
			var bary []byte = make([]byte,10)

			vz := binary.PutUvarint(bary,value)
			if 1 > vz {
				t.Errorf("Error writing value 0x%X",value)
			} else {
				var buf *bytes.Buffer = bytes.NewBuffer(bary)

				var read uint64

				read, er = Read(buf)
				if nil != er {
					t.Errorf("Error reading value 0x%X: %v",value,er)
				} else if value != read {
					t.Errorf("Error reading.  Expected value (0x%X) found value (0x%X).",value,read)
				}
			}
		}
	}
}
