# Interfaces

Why are they useful? By defining a function that takes in an `interface{}`, we essentially give ourselves the flexibility to pass in anything we want. It’s a Go programmers way of saying, this function takes in something, but I don’t necessarily care about its type.

## Defining Interfaces

An interface is essentially a contract. If we define a type based off this interface then we are forced to implement all of the functions or methods defined within that interface type.

Say, for example, we wanted to define an interface for a Guitarist. We could define our interface to include a `PlayGuitar()` function like so:

```go
type Guitarist interface {
  // PlayGuitar prints out "Playing Guitar"
  // to the terminal
  PlayGuitar()
}
```

With our Guitarist interface defined, we could define a BaseGuitarist and an AcousticGuitarist struct which implements the Guitarist interface.

```go
package main

import "fmt"

type Guitarist interface {
    // PlayGuitar prints out "Playing Guitar"
    // to the terminal
    PlayGuitar()
}

type BaseGuitarist struct {
    Name string
}

type AcousticGuitarist struct {
    Name string
}

func (b BaseGuitarist) PlayGuitar() {
    fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
    fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

func main() {
    var player BaseGuitarist
    player.Name = "Paul"
    player.PlayGuitar()

    var player2 AcousticGuitarist
    player2.Name = "Ringo"
    player2.PlayGuitar()
}
```

Should we wish, we could then create an array of type Guitarist which could store both our BaseGuitarist and AcousticGuitarist objects.

```go
var guitarists []Guitarist
guitarists = append(guitarists, player)
guitarists = append(guitarists, player2)
```

## Return values

In real world applications we would have more complex interfaces which specify a return value.

```go
type Employee interface {
    Name() string
    Language() string
    Age() int
    Random() (string, error)
}
```

## Satisfying Interfaces

Say we wanted to create an array of all Employee's in the firm. Within this array, we’d want all of our Engineers.

Now, in order for this to work, we’d need our Engineer type to satisfy the Employee interface or it will not allow us to compile our program:

```go
package main

type Employee interface {
    Language() string
    Age() int
    Random() (string, error)
}

type Engineer struct {
    Name string
}

func (e Engineer) Language() string {
    return e.Name + " programs in Go"
}

func main() {
    // This will throw an error
    var programmers []Employee
    elliot := Engineer{Name: "Elliot"}
    // Engineer does not implement the Employee interface
    // you'll need to implement Age() and Random()
    programmers = append(programmers, elliot)
}
```
