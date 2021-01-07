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
