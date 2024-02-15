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
