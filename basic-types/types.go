// all numeric types default to 0

// unsigned int with 8 bits
// Can store: 0 to 255
var myint uint8

// signed int with 8 bits
// Can store: -127 to 127
var myint int8

// unsigned int with 16 bits
var myint uint16

// signed int with 16 bits
var myint int16

// unsigned int with 32 bits
var myint uint32

// signed int with 32 bits
var myint int32

// unsigned int with 64 bits
var myint uint64

// signed int with 64 bits
var myint int64

// If you attempt to assign a larger value than the int can handle, the compiler will fail
var myint int8
myint = 2500
