# Reading & Writing Files

In order to read from files on your local filesystem, you’ll have to use the `io/ioutil` module. You’ll first have to pull of the contents of a file into memory by calling `ioutil.ReadFile("/path/to/my/file.ext")` which will take in the path to the file you wish to read in as it’s only parameter. This will return either the data of the file, or an `err` which can be handled as you normally handle errors in go.

Create a new file called [main.go](main.go) as well as another file called localfile.data. Add a random piece of text to the .data file so that our finished go program has something to read and then do the following:
