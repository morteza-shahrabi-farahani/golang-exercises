# Chapter 5: Go packages and functions

Go packages, are Go's way of organizing, delivering, and using code. Regarding the visibility of package elements, Go follows a simple rule that states that functions, variables, data types, structure fields, and so forth that begin with an uppercase letter are public, whereas functions, variables, types, and so on that begin with a lowercase letter are private.

## Go packages
Everything in Go is delivered in the form of packages. Apart from the packages of the Go standard library, there are external packages that can be imported using their full address and that should be downloaded on the local machine, before their first use.

packages are mainly used for grouping related functions, variables, and constants so that you can transfer them easily and use them in your own Go programs. Note that apart from the main package, Go packages are not autonomous programs and cannot be compiled into executable files on their own.

\* go get github.com/spf13/cobra => command for downloading packages.

Basically, there are three main directories under ~/go with the following properties:

* The bin directory: This is where binary tools are places.
* the pkg directory: This is where reusable packages are put.
* The src directory: This is where the source code of the packages is located. The underlying structure is based on the URL of the package you are looking for. So, the URL for the github.com/spf13/viper package is ~/go/src/github.com/spf13/viper. If a package is downloaded as  a module, then it will be located under ~/go/pkg/mod.

\* Starting with Go 1.16, go install is the recommended way of building and installing packages in module mode. The use of go get is deprecated, but this chapter uses go get because it's commonly used online and is worth knowing about. However, most of the chapters in this book use go mod init and go mod tidy for downloading external dependencies for your own source files. 

\* If you want to upgrade an existing package, you should execute go get with teh -u option. Additionally, if you want to see what is happening behind the scenes, add the -v option to the go get command - in this case, we are using the viper package as an example, but we abbreviate the output.

## Functions

A piece of advice: functions must be as independent from each other as possible and must do one job (and only one job) well. So, if you find yourself writing functions that do multiple things, you might want to consider replacing them with multiple functions instead.

### Anonymous functions
Anonymous functiosn can be defined inline without the need for a name, and they are usually used for implementing things that require a small amount of code. Note that anonymous functions are called lambdas in functional programming terminology. Similar to that, a closure is a specific type of anonymous function that carries or closes over variables that are in the same lexical scope as the anonymous function that was defined. 