# Composite Types

## Arrays

Arrays can be declared using the following syntax:

```go
// Declare empty array of strings
var days []string

// Declaring array with elements
days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

// Querying items from the array is very similar to other languages
fmt.Println(days[0]) // Prints 'monday'
fmt.Println(days[5]) // Prints 'saturday'
```

## Slices

The difference between an array and a slice is very subtle, a slice allows you to access a subset of an arrays underlying elements.

Slices are made up of three things, a pointer, a length, and a capacity.

```go
days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
weekDays := days[0:5]
fmt.Println(weekDays)

// This returns [Monday, Tuesday, Wednesday, Thursday, Friday]
```

## Maps

A map is Go's representation of a hash table. It allows you to map one data type to another.

In this example we create a map of YouTube channel names to their number of subscribers:

```go
youtubeSubscribers := map[string]int{
  "TutorialEdge":     2240,
  "MKBHD":            6580350,
  "Fun Fun Function": 171220,
}

fmt.Println(youtubeSubscribers["MKBHD"]) // prints out 6580350
```

## Structs

In Go we can create a `struct` which are data types that are aggregates of other data types.

For example, we could create a `Person` in our application that has a number of fields within it, perhaps a `name` which is a `string` and `age` which is an `int`.

```go
// Person struct
type Person struct {
    name string
    age int
}

// Declare new Person
var myPerson Person
```

The advantage of using these struct is that we can effectively treat all of these values or fields as they are called as a single entity and modify that easily.

```go
// Declare new `jethro`
jethro := Person{name: "Jethro", age: 29}

// Change age to 18
jethro.age = 18
```

## Nested Structs

We can create structs within structs.
