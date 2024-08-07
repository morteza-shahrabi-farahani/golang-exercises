# Chapter 6: Telling a UNIX system what to do
This chapter teaches you about systems programming in Go.

\* Starting with Go 1.16, the GO111MODULE environment variable defaults to on - this affects the use of Go packages that do not belong to the Go standard library. In practice, this means that you must put your code under ~/go/src. You can go back to the previous behavior by setting GO111MODULE to auto, but you do not want to do that  - modules are the future.

## stdin, stdout, and stderr
UNIX considers everything, even a printer or your mouse, as a file. UNIX uses file descriptors, which are positive integer values, an an internal representation for accessing open files, which is much prettier than using long paths. So, by default, all UNIX systems support three special and standard filenames: /dev/stdin, /dev/stdout, and /dev/stderr, which can also be accessed using file descriptors 0,1, and 2, respectively. These three file descriptors are also called standard input, standard output, and standard error, respectively.

Go uses os.Stdin for accessing standard input, os.Stdout for accessing standard output, and os.Stderr for accessing standard error. Although you can still use /dev/stdin, /dev/stdout, and /dev/stderr or the related file descriptor values for accessing the same devices, it is better, safer, and more portable to stick with os.Stdin, os.Stdout, and os.Stderr.

## UNIX processes
Strictly speaking, a process is an execution environment that contains instructions, user data and system data parts, and other types of resources that are obrained during runtime. On the other hand, a program is a binary file that contains instructions and data that are used for initializing the instruction and user data parts of a process. Each running UNIX process is uniquely identified by an unsigned integer, which is called the process ID of the process.

There are three process categories: user processes, daemon processes, and kernel processes. User processes run in user space and usually have no special access rights. Daemon processes are programs that can be found in the user space and run in the background without the need for a terminal. Kernel processes are executed in kernel space only and can fully access all kernel data structures.

## Buffered and unbuffered file I/O
Buffered file I/O happens when there is a buffer for temporarily storing data before reading data or writing data. Thus, instead of reading a file byte by byte, you read many bytes at once. You put the data in a buffer and wait for someone to read it in the desired way. 

Unbuffered file I/O happens when there is no buffer to temporarily store data before actually reading or writint it.

The next question that you might ask is how to decide when to use buffered and when to use unbuffered file I/O. When dealing with critical data, unbuffered file I/O is generally a better choice because buffered reads might result in out-of-date data and buffered writes might result in data loss when the power of your computer is interrupted. 

Buffered readers can also improve performance by reducing the number of system calls needed to read from a file or socket, so there can be a real performance impact on what the programmer decides to use.

### Reading from /dev/random 
The purpose of the /dev/random system device is to generate random data, which you might use for testing your programs.

## Working with JSON
Go allows you to add support for JSON fiellds in Go structures using tags. Tags control the encoding and decoding of JSON records to and from Go structures.

Marshaling is the process of converting a Go structure iinto a JSON record. Unmarshaling is the process of converting a JSON record given as a byte slice into a Go structure. 

```
type UseAll struct {
    Name string `json:"username"`
}
```

The previous metadata tells us that the Name field of the UseAll structure is translated to username in the JSON record and vice versa.

Imagine tht you have a Go structure that you want to convert into a JSON record without including any empty fields - the next code illustrates how to perform that task with the use of omitempty

```
 type NoEmpty struct {
    Name    string `json:"username"`
    Surname string `json:"surname"`
    Year    int    `json:"creationyear,omitempty"`
}

// now if we have a variable NoEmpty{Name:"Morteza"} and marshal it to JSON, 
// it will give this result {"username": "Morteza", "surname":""} 
```

The noEmpty structure has the default values for surname and year fields. However, as they are not specifically defined, json.Marshal() ignores the Year field because it has the omitempty tag but does not ignore the Surname fields, which has the empty string value.

Last, imagine that you have some sensitive data on some of the fields of a Go structure that you do not want to include in the JSON records. You can do that by including the "-" special value in the desired json structure tags.

```
 type Password struct {
    Name     string `json:"username"`
    Surname  string `json:"surname,omitempty"`
    Year     int    `json:"creationyear,omitempty"`
    Pass     strubg `json:"-"`
}
```

## Working with XML
The idea behind XML and Go is the same as with JSON and Go. You put tags in Go structures in order to specify the XML tags and you can still serialize and deserialize XML records using xml.Unmarshal() and xml.Marshal()

\* A field with the omitempty option is omitted from the output if it is empty. An empty value is any of 0,false, a nil pointer or interface, and any array, slice, map, or string with a length of zero.

\* The go mod init command initializes and writes a new go.mod file in the current directory whereas the go mod tidy command synchronizes go.mod with the source code.

\* If you want to play it safe and you are using packages that do not belong to the standard library, then developing inside ~/go/src, committing to a Github repository, and using Go modules for all dependencies might be the best option. However, this does not mean that you must develop your own packages in the form of Go modules.

## The viper package
In Go, there are many packages to handle application configuration. The viper package is most popular among them in providing a complete configuration solution of an application. It supports numerous configuration file formats such as JSON, YAML, TOML, HCL and Java properties format. This programming tutorial introduces Golang’s viper package with Go code examples.

## The cobra package
Cobra is a very handy and popular Go package that allows you to develop command-line utilities with commands, subcommands, and aliases.

\* It is not necessary to know about all of the supported environment variables such as GO111MODULE, but sometimes they can help you resolve tricky problems with your GO installation. So, if you want to learn about your current Go environment, you can use the go env command.

