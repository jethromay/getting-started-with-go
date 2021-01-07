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

## Multiple results from a function

Often we want to return multiple values from a function, with the result being the first value returned and any potential errors being returned in the second result.

This practice is useful as it allows us to determine what to do with errors returned from the original function:

```go
func myfunction(firstName string, lastName string) (string, error) {
  return firstName + " " + lastName, nil
}
```

See full example in [multiple.go](multiple.go)
