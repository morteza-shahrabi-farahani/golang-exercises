# Chapter 2: Basic Go data types

## The error data type
It should be clearly documented how to handle critical errors. This means that there will be situations where a critical error should terminate the program and other times where a critical error might just create a warning message onscreen.

It is considered a good practice to send all error messages to the log service of your machine because this way the error messages can be examined at a later time.

\* create your own error messages => `errors.New()` <br>
\* format your error messages like fmt.Printf => `fmt.Errorf()` <br>
\* `strconv.Atoi()` => Convert string (ASCII) to Int

## Numeric data types
int and uint can be either 32 or 64 bits each. Their size is defined by Go itself. <br>
When using := for defining integer variables, Go prefers to use int data type for storing these values.

## Non-numeric data types
Strings, Characters, Runes, Dates and Times.

### Strings, Characters, and Runes
A Go strings is just a collection of bytes and can be accessed as a whole or as an array. A single byte can store any ASCII character -- however, multiple bytes are usually needed for storing a single Unicode character. And the main reason for having the rune data type is supporting these Unicode characters. A rune is an int32 value that is used for representing a single Unicode code point, which is an integer value that is used for representing single Unicode characters or, less frequently, providing formatting information.

`[]byte("Given string")` => Create a new byte slice from given string. <br>
`string(byteSlice)` => Convert byte slice into string. <br>

When working with byte slices that contain Unicode characters, the number of bytes in a byte slice is not always connected to the number of characters in the byte slice, because most Unicode characters require more than one byte for their representation. 

\* when you try to print each single byte of a byte slice using fmt.Println() or fmt.Print(), the output is not text presented as characters but integer values. If you want to print the contents of a byte slice as text, you should either print it using string(byteSliceVar) or using fmt.Printf() with %s to tell fmt.
Printf() that you want to print a string. You can initialize a new byte slice with the desired string by using a statement such as []byte("My Initialization String").

`r := '$'` => define a rune. <br>
print the integer value => `fmt.Println(r)`<br>
printing as a single Unicode character => `fmt.Printf("%c", r)`

For accessing rune value from a string => `for index, value range str {fmt.Println(value)}`<br>
For accessing byte value from a string => `for i:=0; i<len(str); i++ {fmt.Println(str(i))}`

`[]rune("Given string")` => Create a new rune slice from given string. <br>
`string(runeSlice)` => Convert rune slice into string. <br>

\* For finding complete list of available functions of strings package in go, you can visit https://golang.org/pkg/strings/

## Times and dates
The king of working with times and dates in Go is the time.Time data type, which 
represents an instant in time with nanosecond precision. Each time.Time value is 
associated with a location (time zone).

The time.Now().Unix() function returns the popular UNIX epoch time, which is the number of seconds that have elapsed since 00:00:00 UTC, January 1, 1970. And if you want to convert UNIX time to time.Time, you can use time.Unix() function.

\* Go parses a string in order to convert it into a date and a time. The function used for parsing is time.Parse() and its full signature is Parse(layout, value string) where layout is the parse string and value is the input that is being parsed.

most widely used strings for parsing dates and times are like this.
![Local Image](./parsing%20dates%20and%20time%20table.png "parsing dates and time table")

For example if you want to parse the 30 January 2020, you should use 02 January 2006 as layout. Pay attention that you can not use anything else in layout if you want that format.

\* The formatting strings can be also used for printing dates and times in the desired format. So in order to print the current date in the 01-02-2006 format, you should use time.Now().Format("01-02-2006")

\* If a command-line argument such as 14 December 2020 contains space characters, you should put it in double quotes for the UNIX shell to treat it as a single command-line argument. Running go run dates.go 14 December 2020 does not work.

\* Use this command to convert time to different time zones and locations. <br>
```
// now is a time variable
loc, _ := time.LoadLocation("Local")
fmt.Printf("Current Location: %s\n", now.In(loc))
```

## Go constants
The value of a constant variable is defined at compile time, not at runtime, this means that it is included in the binary executable. The value of constant variables cannot be changed. Their can be either global or local variables. However, usually they used as global variables.

### The constant generator iota
The constant generator iota is used for declaring a sequence of related values that use incrementing numbers without the need to explicitly type each one of them. 

```
const (
    Zero = iota // 0
    One         // 1
    Two         // 2
    Three       // 3
)
```

\* Using underscore character in a const block with a constant generator iota, allows you to skip unwanted values.

## Grouping similar data
You can use slices instead of arrays almost anywhere in Go but we are also demonstrating arrays because they can still be useful and because slices are implemented by Go using arrays!

### Arrays
When defining an array variable, you must define its size. Otherwise, you should put [...] in the array declaration and let the Go compiler find out the length for you. So you can create an array with 4 string elements either as [4]string{"Zero", "One", "Two", "Three"} or as [...]string{"Zero", "One", "Two", "Three"}. If you put nothing in the square brackets, then a slice is going to be created instead. <br>

\* You cannot change the size of an array after you have created it. <br> 

\* When you pass an array to a function, what is happening is that Go creates a copy of that array and passes that copy to that function - therefore any changes you make to an array in side a function are lost when the function returnes. 
As a result, arrays in Go are not very powerful, which is the main reason that Go has introduced an additional data structure named slice that is similar to an array but is dynamic in nature.

### Slices
Slices in Go are more powerful than arrays mainly because they are dynamic, which means that they can grow or shrink after creation if needed. Additionally, any changes you make to a slice inside a function also affect the original slice. But how does this happen? Strictly speaking, all parameters in Go are passed by value - there is no other way to pass parameters in Go.

In reality, a slice value is a header that contains a pointer to an underlying array where the elements are actually stored, the length of the array, and its capacity.

