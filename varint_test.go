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
 * OFFLINE TEST
 * 
 *    --- FAIL: TestGoVintPut (0.00s)
 *  panic: runtime error: index out of range [8] with length 8 [recovered]
 * 	 panic: runtime error: index out of range [8] with length 8
 * 
 *  goroutine 22 [running]:
 *  testing.tRunner.func1.2({0x521dc0, 0xc0000cc000})
 * 	 /usr/local/src/go-1.20.11/src/testing/testing.go:1526 +0x24e
 *  testing.tRunner.func1()
 * 	 /usr/local/src/go-1.20.11/src/testing/testing.go:1529 +0x39f
 *  panic({0x521dc0, 0xc0000cc000})
 * 	 /usr/local/src/go-1.20.11/src/runtime/panic.go:884 +0x213
 *  encoding/binary.PutUvarint(...)
 * 	 /usr/local/src/go-1.20.11/src/encoding/binary/varint.go:58
 *  github.com/syntelos/go-varint.TestGoVintPut(0xc00009aea0)
 * 	 /home/jdp/go/src/github.com/syntelos/go-varint/varint_test.go:145 +0x2e5
 *  testing.tRunner(0xc00009aea0, 0x539598)
 * 	 /usr/local/src/go-1.20.11/src/testing/testing.go:1576 +0x10b
 *  created by testing.(*T).Run
 * 	 /usr/local/src/go-1.20.11/src/testing/testing.go:1629 +0x3ea
 *  exit status 2
 *  FAIL	github.com/syntelos/go-varint	0.006s
 * 
 *  Compilation exited abnormally with code 1 at Mon Dec  4 17:56:45
 * 
 */
func _TestGoVintPut(t *testing.T){
	var sequence uint8
	for sequence = 1; sequence <= 16; sequence++ {

		var er error

		var value uint64
		value, er = endian.Rand64() 
		if nil != er {
			t.Error(er)
		} else {
			var bary []byte = make([]byte,8)

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
