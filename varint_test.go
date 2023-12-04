/*
 * ULEB64 VARINT I/O
 * Copyright 2023 John Douglas Pritchard, Syntelos
 */
package varint

import (
	"bytes"
	"testing"
)
/*
 */
func Test00E7(t *testing.T){
	const value uint64 = 0x00E7
	var count uint64 = Count(value)
	var buf *bytes.Buffer = new(bytes.Buffer)

	sz, er := Write(buf,value)
	if nil != er {
		t.Errorf("TestIO error writing value 0x%X: %v",value,er)
	} else if count != sz {
		t.Errorf("TestIO error writing value 0x%X.  Expected length (%d) found length (%d).",value,count,sz)
	} else {
		var read uint64
		read, er = Read(buf)
		if nil != er {
			t.Errorf("TestIO error reading value 0x%X: %v",value,er)
		} else if value != read {
			t.Errorf("TestIO error reading.  Expected value (0x%X) found value (0x%X).",value,read)
		}
	}
}
/*
 */
func Test1200(t *testing.T){
	const value uint64 = 0x1200
	var count uint64 = Count(value)
	var buf *bytes.Buffer = new(bytes.Buffer)

	sz, er := Write(buf,value)
	if nil != er {
		t.Errorf("TestIO error writing value 0x%X: %v",value,er)
	} else if count != sz {
		t.Errorf("TestIO error writing value 0x%X.  Expected length (%d) found length (%d).",value,count,sz)
	} else {
		var read uint64
		read, er = Read(buf)
		if nil != er {
			t.Errorf("TestIO error reading value 0x%X: %v",value,er)
		} else if value != read {
			t.Errorf("TestIO error reading.  Expected value (0x%X) found value (0x%X).",value,read)
		}
	}
}
