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