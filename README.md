# Gencode
Gencode is a code generation based data serialization/deserialization system.

Code is generated from a schema that is similar to native Go semantics, though there are a few differences/additions

For example:
```
struct Person {
  Name string
  Age uint8
}
```

Base data types available:
int/uint/vint/vuint in 8/16/32/64 bit varieties (v prefix indicates varint encoding)
float32/64
string
byte

Extended data types are built off of the base data types and include:
Slices
Pointer
Tagged Unions
Any other gencoded struct (must be declared in the same file)

##Tagged Unions
Tagged unions allow you to have a field that may contain one of a multiple of different types.

Example:
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

Message.Request can contain either a Subscribe or an Unsubscribe.  The protocol tags the field, so deserialization can create the proper objects.
The field itself is declared as an interface{}, and you can type switch on it.  Alternatively, you can give an interface name to use:
```
struct Message {
  Request union Command {
    Subscribe
    Unsubscribe
  }
}
```
The Request field will be declared of type Command, which will be must to be an interface that all the types in that union implement.

# Speed

Gencode encodes to smaller amounts of data, and does so very fast.  Some benchmarks:
```
Gencode encoded size: 47
GOB encoded size: 182
GOB Stream encoded size: 62
JSON encoded size: 138
MSGP encoded size: 115
PASS
BenchmarkFixedBinarySerialize-8          2000000               853 ns/op
BenchmarkFixedBinaryDeserialize-8        3000000               518 ns/op
BenchmarkGencodeSerialize-8              3000000               411 ns/op
BenchmarkGencodeDeserialize-8            2000000               622 ns/op
BenchmarkFixedGencodeSerialize-8         5000000               269 ns/op
BenchmarkFixedGencodeDeserialize-8       5000000               285 ns/op
BenchmarkGobSerialize-8                   200000              9049 ns/op
BenchmarkGobDeserialize-8                  50000             38447 ns/op
BenchmarkGobStreamSerialize-8            1000000              1654 ns/op
BenchmarkGobStreamDeserialize-8          1000000              2057 ns/op
BenchmarkJSONSerialize-8                  500000              2713 ns/op
BenchmarkJSONDeserialize-8                300000              5194 ns/op
BenchmarkMSGPSerialize-8                 5000000               424 ns/op
BenchmarkMSGPDeserialize-8               2000000               863 ns/op
```
