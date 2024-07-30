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

### The defer keyword

The defer keyword postpones the execution of a function until the surrounding function returns. 

Usually, defer is used in file I/O operations to keep the function call that closes and opened file close to the call that opened it, so that you do not have to remember to close a file that you have opened just before the function exists.

Deferred functions are executed in last in, first out(LIFO) order after the surrounding function has been returned.

## Developing your own packages

It is a best practice to use lowercase package names, even through uppercase package names are allowed. 

Compiling a Go package can be done manually, if the package exists on the local machine, but it is also done automatically after you download the package from the internet. If the package you are downloading contains any errors, you will learn about them at downloading time.

### The init() function
Each Go package con optionally have a private function named init() that is automatically executed at the beginning of execution time.

* init() takes no arguments.
* init() returns no values.
* The init() function is called implicitly by Go.
* All init() functions are always executed prior to the main() function.
* A source file can contain multiple init() functions - these are executed in the order of declaration.
* Go packages can contain multiple files. Each source file can contain one or more init() functions.

init() function is a private function by design and it cannot be called from outside the package in which it is contained. Additionally, as the user of a package has no control over the init() function, you should think carefully before using an init() function in publich packages or changing any global state in init().

There are some situations where the use of init() makes sense:

* For initializing network connections that might take time prior to the execution of package functions or methods.

* For initializing connections to one or more servers prior to the execution of package functions or methods. 

* For creating required files and directories. 

* For checking whether required resources are available or not.

### Order of execution
If a main package imports package A and package A depends on package B, then the following will take place:
* The process starts with main package
* The main package imports package A
* Package A imports package B
* The global variables, if any, in package B are initialized.
* The init() function or functions of package B, if they exist, run. This is the first init() function that gets executed.
* The global variables, if any, in package A are initialized.
* The init() function or functions of package A, if there are any, run.
* The global variables in the main package are initialized.
* The init() function or functions of main package, if they exist, run.
* The main() function of the main package begins execution.


![Local Image](./Order%20of%20execution%20in%20Go.png "Order of execution in Go")

## A package for working with a database

\* The sql.Open() function opens the database connection and keeps it open until the program ends.

\* You need to assign the values returned from the SELECT query into Go variables, in order to use them. This happens with a call to Scan(), which requires pointer parameters. If the SELECT query returns multiple values, you need to put multiple parameters in Scan().

```
insertStatement = `insert into "userData" ("userid", "name", "surname", "description") values ($1, $2, $3, $4)`

_, err = db.Exec(insertStatement, userID, d.Name, d.Surname, d.Description)
```
## Modules 
A Go module is like a Go package with a version - however, Go modules can consist of multiple packages. Go uses semantic versioning for versioning modules. This means that version begin with the letter v, followed by the major.minor.patch version numbers. Therefore, you can have versions such as v1.0.0, v1.0.5 and so on. The v1, v2, and v3 parts signify the major version of a Go package that is usually not backward compatible. This means that if your Go program works with v1, it will not necessarily work with v2 or v3. The second number in a version is about features. Usually, v1.1.0 has more features than v1.0.2 or v1.0.0, while being compatible with all older versions. Lastly, the third number is just about bug fixes without having any new features. Note that semantic versioning is also used for Go versions.

\* It is better to split the functionality of a package unnecessarily into multiple packages than to add too much functionality to a single Go package.

\* When developing a new Go package, try to use multiple files in order to group similar tasks or concepts.

\* Nobody wants a Go package that prints logging information on the screen. It would be more professional to have a flay for turning on logging when needed.

\* Small details make all the difference3 and give people confidence that you are a serious developer!

## Generating documentation
In order to document a function, a method, a variable, or even the package itself, you can write comments, as usual, that should be located directly before the element you want to document, without any empty lines in between. You can use one or more single-line comments, which are lines beginning with //, or block comments, which begin with /* and end with */.

It is highly recommended that each Go package you create has a block comment preceding the package declaration that introduces developersw to the package, and also explains what the package does. 

If a line in a block comment begins with a tab, then it is rendered differently in the graphical output, which is good for differentiating between various kinds of information in the documentation.

There are two ways to see the documentation of the package. The first one involves using go get, which also means creating a GitHub repository of the package, as we did with post05. However, as this is for testing purposes, we are going to do things the easy way: we are going to copy it in ~/go/src and access it from there. As the package is
called document, we are going to create a directory with the same name inside ~/go/ src. After that, we are going to copy document.go in ~/go/src/document and we are done—for more complex packages, the process is going to be more complex as well. In such cases, it would be better to go get the package from its repository.

Either way, the go doc command is going to work just fine with the document package: <br>
-> go doc document <br>
If you want to see information about a specific function, you should use go doc, as follows: <br>
-> go doc document ListUsers
Additionally, we can use the web version of the Go documentation, which can be accessed after running the godoc utility and going to the Third Party section.
