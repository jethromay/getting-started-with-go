# Cross Compatibility Issues

Please note that some of these commands may not work on your operating system. If you are trying to write code that is compatible on multiple platforms then it would be wise to select commands that only feature on all platforms. If this is un-achievable then I recommend you add conditional logic to your program that executes a different system command depending on the system it’s executing on top of.

## Checking Current Operating System

In order to check what operating system our code is running on we can use the runtime package and check the GOOS constant. This will return the operating system target:

```go
    if runtime.GOOS == "windows" {
        fmt.Println("Can't Execute this on a windows machine")
    } else {
        execute()
    }
```

## Implementation

Let’s have a look at how we can start executing some really simple commands such as `ls` and `pwd` using the `os/exec` package, and once we have the basics covered, we can move on to more advanced examples.

We’ll first of all need to import the 3 key packages for this example, the `fmt`, `os/exec` and the `runtime` package.

Once, we’ve done this, we’ll define an `execute()` function which will attempt to execute the following code:

```go
package main

import (
    "fmt"
    "os/exec"
    "runtime"
)

func execute() {

    // here we perform the pwd command.
    // we can store the output of this in our out variable
    // and catch any errors in err
    out, err := exec.Command("ls").Output()

    // if there is an error with our execution
    // handle it here
    if err != nil {
        fmt.Printf("%s", err)
    }
    // as the out variable defined above is of type []byte we need to convert
    // this to a string or else we will see garbage printed out in our console
    // this is how we convert it to a string
    fmt.Println("Command Successfully Executed")
    output := string(out[:])
    fmt.Println(output)

    // let's try the pwd command herer
    out, err = exec.Command("pwd").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Command Successfully Executed")
    output = string(out[:])
    fmt.Println(output)
}

func main() {
    if runtime.GOOS == "windows" {
        fmt.Println("Can't Execute this on a windows machine")
    } else {
        execute()
    }
}
```

The output should list the files in the directory.

## Passing in Arguments

We have so far only used standard commands, to pass in arguments we will use `.Command()`
