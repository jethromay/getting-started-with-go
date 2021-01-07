# Basic Types

There are 4 main types in Go:

- Basic Types
- Aggregate Types - Arrays & Structs
- Reference Types - Pointers & Slices
- Interface Types - Standard Interfaces

All numeric types default to 0

Unsigned int with 8 bits
Can store: 0 to 255
```go
var myint uint8
```

Signed int with 8 bits
Can store: -127 to 127
```go
var myint int8
```

Unsigned int with 16 bits
```go
var myint uint16
```

Signed int with 16 bits
```go
var myint int16
```

Unsigned int with 32 bits
```go
var myint uint32
```

Signed int with 32 bits
```go
var myint int32
```

Unsigned int with 64 bits
```go
var myint uint64
```

Signed int with 64 bits
```go
var myint int64
```

If you attempt to assign a larger value than the int can handle, the compiler will fail
```go
var myint int8
myint = 2500
```
