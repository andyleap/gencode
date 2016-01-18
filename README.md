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
Gencode encoded size: 48
GOB encoded size: 182
GOB Stream encoded size: 62
JSON encoded size: 138
MSGP encoded size: 115
PASS
BenchmarkFixedBinarySerialize-8          2000000               894 ns/op
BenchmarkFixedBinaryDeserialize-8        3000000               539 ns/op
BenchmarkGencodeSerialize-8             10000000               174 ns/op
BenchmarkGencodeDeserialize-8           10000000               219 ns/op
BenchmarkFixedGencodeSerialize-8        20000000                75.7 ns/op
BenchmarkFixedGencodeDeserialize-8      100000000               20.7 ns/op
BenchmarkGobSerialize-8                   200000              9370 ns/op
BenchmarkGobDeserialize-8                  30000             40337 ns/op
BenchmarkGobStreamSerialize-8            1000000              1694 ns/op
BenchmarkGobStreamDeserialize-8          1000000              2125 ns/op
BenchmarkJSONSerialize-8                  500000              2780 ns/op
BenchmarkJSONDeserialize-8                300000              5263 ns/op
BenchmarkMSGPSerialize-8                 5000000               277 ns/op
BenchmarkMSGPDeserialize-8               2000000               608 ns/op
```
