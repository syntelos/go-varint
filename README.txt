ULEB64 I/O for GOPL

  func Write(io.ByteWriter, uint64) (uint64, error)
  func Read(io.ByteReader) (uint64, error)

References

 [ULEB64] https://en.wikipedia.org/wiki/LEB128
 [VARINT] https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/encoding/binary/varint.go
 [GOPL] https://go.dev/

