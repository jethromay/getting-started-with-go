# Functions

Function declarations in Go are very similar to other languages, they begin with `func` followed by the name of the function.

```go
func test() {
    // Function body
}
```

> Capitalization Matters! If you want your functions to be accessible in other packages then you will have to make the first
> letter of your function name Uppercase!

## A simple example

```go
func myFunction(firstName string, lastName string) (string) {
    fullName := firstName + " " + lastName
    return fullName
}
```

See full example in [main.go](main.go)
