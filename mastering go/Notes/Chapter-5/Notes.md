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

\* Anonymous functiosn can be defined inline without the need for a name, and they are usually used for implementing things that require a small amount of code. Note that anonymous functions are called lambdas in functional programming terminology. Similar to that, a closure is a specific type of anonymous function that carries or closes over variables that are in the same lexical scope as the anonymous function that was defined. 

\* If you have a function that returns more than 3 values, you should reconsider that decision and maybe redesing it to use a single structure or slice for grouping and returning the desired values as a single entity.

\* When a function has a return statement without any arguments, the function automatically returns the current value of each named return value.

```
func minMax(x, y int) (min, max int) {
    if x > y {
        min = y
        max = x
        return min, max
    }

    min = x
    max = y
    return
}
```

\* You can provide the sort.Slice() function with another function as an argument that specifies the way sorting is implemented. The signature of sort.Slice() is func Slice(slice interface{}, less func(i, j int) bool)

* The sort.Slice() function does not return any data.
* The sort.Slice() function requires two arguments, a slice of type interface{} and another function - the slice variable is modified inside sort.Slice()
* there is no need for you to name the anonymous function. The name less is required because all function parameters should have a name.

### Functional programming in Go
Go is a functional programming language, albeit not a pure one, and allows Go to benefit from the functional programming paradigm.

Here are some key characteristics of functional programming languages:

* **Immutability:** <br>
Functional programming languages promote the use of immutable data structures, meaning once a value is assigned to a variable, it cannot be changed. This helps eliminate side effects, making programs more predictable and easier to reason about.

* **First-Class and Higher-Order Functions :** <br>
Functions may be treated as first-class citizens in functional programming, having the ability to be assigned to variables, provided as arguments to other functions, and returned as values. As they have the ability to return their own values or take in other functions as inputs, higher-order functions are a basic idea in functional programming.

* **Pure Functions:** <br>
Declarative programming, in which programmers state their goals rather than outlining the procedures necessary to get them, is encouraged by functional programming. As a result, the code becomes more legible and succinct.

* **Recursion :** <br>
Recursion is a common technique in functional programming for performing repetitive tasks. Instead of using loops, functional languages often rely on recursive functions, which contribute to the elegant and concise nature of functional code.

### Variadic functions
Variadic functions are functions that can accept a variable number  of parameters.

\* Variadic functions use the pack operator, which consists of a ..., followed by a data type. So, for a variadic function to accept a variable number of int values, the pack operator should be ...int.

\* The pack operator can only be used once in any given function. 

\* The variable that holds the pack operation is a slice and, therefore, is accessed as a slice inside the variadic function.

\* The variable name that is related to the pack operator is always last in the list of function parameters. 

\* ...interface{} accept a variable number of arguments of all data types.

```
func addFloats(message string, s ...float64) float64 {
    fmt.Println(message)
    sum := float64(0)
    for _, a := range s {
        sum = sum + a
    }

    s[0] = -1000
    return sum
} 

s := []float64{1.1, 2.12, 3.14}
sum = addFloats("Adding numbers...", s...)
```

