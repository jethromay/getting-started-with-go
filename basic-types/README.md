# Basic Types

There are 4 main types in Go:

- Basic Types
- Aggregate Types - Arrays & Structs
- Reference Types - Pointers & Slices
- Interface Types - Standard Interfaces

## Integer

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

For simplicity, it’s best to default to int which will be converted to 32bit or 64bit depending on your underlying operating system.

## Floating Point Numbers

There are 2 distinct sizes, `float32` and `float64` which deal with numbers larger than what `int64` can handle.

```go
var f1 float32
var f2 float64

var maxFloat32 float32
maxFloat32 = 16777216
fmt.Println(maxFloat32 == maxFloat32+10) // you would typically expect this to return false
// it returns true
fmt.Println(maxFloat32+10) // 16777216
fmt.Println(maxFloat32+2000000) // 16777216
```

Converting floats to int and back to floats can be achieved by casting the variable:

```go
// Convert from int to float
var myint int
myfloat := float64(myint)

// Convert from float to int
var myfloat2 float64
myfloat2 := int(myfloat2)
```
