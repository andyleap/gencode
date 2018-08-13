# Gencode
Gencode is a code generation based data serialization/deserialization system.  Gencode attempts to both encode/decode fast, and have a small data size.

Code is generated from a schema that is similar to native Go semantics, though there are a few differences/additions

For example:
```
struct Person {
  Name string
  Age uint8
}
```

Run through using `gencode go -schema test.schema -package test`

Yields:
```
package test

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type Person struct {
	Name string
	Age  uint8
}

func (d *Person) Size() (s uint64) {
 ...
}
func (d *Person) Marshal(buf []byte) ([]byte, error) {
 ...
}
func (d *Person) Unmarshal(buf []byte) (uint64, error) {
 ...
}
```
(bulk removed for size reasons)

# Data Types
## Struct
Structs are built, similar to native Go, from various fields of various types.  The format is slightly different, putting the `struct` keyword in front of the name of the struct and dropping the `type` keyword, in order to differentiate Gencode schemas from Go code.  Structs may optionally be "framed", adding `Serialize` and `Deserialize` functions taking a `io.Writer` or `io.Reader` respectively.  These structs have a prefixed `vuint64` for the length of the whole struct, minus the prefix length.  This allows efficient reading from network sockets and other streams.

### Int
Integer data types consist of both signed and unsigned ints, in 8, 16, 32, and 64 bit lengths.  In addition, any type can be varint encoded by prefixing it with the letter `v`.  Some examples:

* `uint16`
* `vuint32`
* `vint64`
* `int32`

### Float
Float types are allowed in either 32 or 64 bit lengths.

### String
Strings are encoded with a prefixed `vuint64` for length, so short strings only require a 1 or 2 byte prefix, but strings of practically any length can be used.

### Byte
Bytes are basically an alias to uint8, though there is an optimization for a slice of bytes, i.e. []byte

### Bool
Bools are stored as either a 0 or a 1 for false or true

### Fixed Length Arrays
Fixed Length Arrays as encoded as the designated number of elements, with no length prefix.  Note that the number of elements is fixed, but the elements themselves may take a variable number of bytes to actually encode.  Examples:
* `[5]vuint64`
* `[16]float64`

### Slices
Slices, as in go, are a variable length sequence that can be made out of any other valid gencode type.  Slices are also prefixed with a `vuint64` for length.  Examples:
* `[]byte`
* `[][]int64`

### Pointers
Pointers translate directly into pointers on the Go struct as well, and are also used to allow potentially empty fields.  A pointer field has a "prefix" of 1 byte, though if that byte is 0, the field will be set to nil, and there will be no more data for that field in the marshalled data.

### Tagged Unions
Tagged Unions are one of the high points of the Gencode format and system.  There is no direct match in the Go language itself, so tagged unions are handled on the Go side using interfaces.  Tagged unions have a prefix vuint64 specifying the actual type of the field, and that field's serialization semantics then take over.  This allows widely disjoint data types to be stored in the same field.  While tagged unions can use all other types in the gencode system, the standard use is to use structs defined in the schema.  Example:
```
struct Subscribe {
  Topic string
}

struct Unsubscribe {
  Topic string
}

struct Message {
  Request union {
    Subscribe
    Unsubscribe
  }
}
```
Message.Request can contain either a Subscribe or an Unsubscribe.
The field itself is declared as an interface{}, and you can type switch on it.  Alternatively, you can give an interface name to use:
```
struct Message {
  Request union Command {
    Subscribe
    Unsubscribe
  }
}
```
The Request field will be declared of type Command, which must be an interface that all the types in that union implement.

# Speed

Gencode encodes to smaller amounts of data, and does so very fast.  Some benchmarks (using schemas and test files located in the bench folder):
```
Colfer encoded size: 62
Gencode encoded size: 48
GOB encoded size: 183
GOB Stream encoded size: 62
JSON encoded size: 138
MSGP encoded size: 115
goos: darwin
goarch: amd64
pkg: github.com/andyleap/gencode/bench
BenchmarkFixedBinarySerialize-12       	 3000000	       569 ns/op
BenchmarkFixedBinaryDeserialize-12     	 5000000	       358 ns/op
BenchmarkColferSerialize-12            	30000000	        45.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkColferDeserialize-12          	10000000	       232 ns/op	     144 B/op	       6 allocs/op
BenchmarkFixedColferSerialize-12       	200000000	         9.10 ns/op
BenchmarkFixedColferDeserialize-12     	200000000	         9.62 ns/op
BenchmarkGencodeSerialize-12           	20000000	        90.3 ns/op	      48 B/op	       1 allocs/op
BenchmarkGencodeDeserialize-12         	20000000	        97.5 ns/op	      16 B/op	       4 allocs/op
BenchmarkFixedGencodeSerialize-12      	100000000	        11.4 ns/op
BenchmarkFixedGencodeDeserialize-12    	200000000	         7.79 ns/op
BenchmarkGobSerialize-12               	  200000	      6737 ns/op
BenchmarkGobDeserialize-12             	   50000	     28500 ns/op
BenchmarkGobStreamSerialize-12         	 1000000	      1211 ns/op
BenchmarkGobStreamDeserialize-12       	 1000000	      1420 ns/op
BenchmarkJSONSerialize-12              	 1000000	      1847 ns/op
BenchmarkJSONDeserialize-12            	  300000	      5087 ns/op
BenchmarkMSGPSerialize-12              	10000000	       132 ns/op	     144 B/op	       1 allocs/op
BenchmarkMSGPDeserialize-12            	 5000000	       286 ns/op	      16 B/op	       4 allocs/op
```
