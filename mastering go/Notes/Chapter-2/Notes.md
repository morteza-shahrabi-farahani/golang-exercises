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

How arrays are stored in Go? <br>
Go arrays are laid out contigously in memory. Then since Go types are statically sized, the address of the nth item is equal to the address of the 0th element plus a byte offset equal to the size of the type of the item.

Go's arrays are values. An array variable denotes the entire array; it is not a pointer to the first array element (as would be the case in C). This means that when you assign or pass around an array value you will make a copy of its contents.

### Slices
Slices in Go are more powerful than arrays mainly because they are dynamic, which means that they can grow or shrink after creation if needed. Additionally, any changes you make to a slice inside a function also affect the original slice. But how does this happen? Strictly speaking, all parameters in Go are passed by value - there is no other way to pass parameters in Go.

In reality, a slice value is a header that contains a pointer to an underlying array where the elements are actually stored, the length of the array, and its capacity. Nota that the slice value does not include its elements, just a pointer to the underlying array. So, when you pass a slice to a function, Go makes a copy of that header and passes it to the function. This copy of the slice header includes the pointer to the underlying array. The slice header is defined in the reflect package as follows:

```
type SliceHeader struct {
    Data uintptr
    Len int
    Cap int
}
```

You can create a slice using make() or like an array without specifying its size or using [...]. If you do not want to initialize a slice, then using make() is better and faster.

```
aSlice := []float64{1.2, 3.2, -4.5}
aSlice := make([]float64, 3) 
\* Each element of this slice has a value of 0, which is the zero value of the float64 data type. *\

bSlice := make([][]int, 2) 
\* This returns a slice with two dimensions where the first dimension is 2 (rows) and the second dimension (columns) is unspecified and should be explicitly specified when adding data to it. *\

twoD := [][]int{{1, 2, 3}, {4, 5, 6}}.
```

You can add new elements to a full slice using the append() function. append() automatically allocates the required memory space.

#### Capacity
The capacity shows how much a slice can be expanded without the need to allocate more memory and change the underlying array. The first argument of make() is the type of the slice and its dimensions, the second is its initial length and the third, which is optional, is the capacity of the slice. Although the data type of a slice cannot change after creation, the other two properties can change.

how changing the capacity works in Go by a picture. When you add a new element to a full capacity slice, its capacity will be doubled.
![Local Image](./How%20capacity%20works%20in%20slices.jpeg "how capacity works in slices")

\* aSlice = append(aSlice, []int{-1, -2, -3, -4}...) <br>
The ... operator expands []int{-1, -2, -3, -4} into multiple arguments and
append() appends each argument one by one to aSlice.

\* Setting the correct capacity of a slice, if known in advance, will make your programs faster because Go will not have to allocate a new underlying array and have all the data copied over.

#### Selecting a part of a slice
There is a variation where you can add a third parameter that controls the capacity of the resulting slice. So, using aSlice[0:2:4] selects the first 2 elements of a slice and creates a new slice with a maximum capacity of 4. The resulting capacity is defined as the result of the 4 - 0 subtraction where 4 is the maximum capacity and 0 is the first index. 

\* Selecting last 2 elements of slice <br>
```
aSlice[len(aSlice)-2:]
```

#### Byte slices
bytes are a universal unit among computer systems.<br>
As Go does not have a char data type, it uses byte and rune for storing character values. A single byte can only store a single ASCII character whereas a rune can store Unicode characters. However, a rune can occupy multiple bytes.

\* Convert string to a byte slice<br>
```
b = []byte("Byte slice $")
```
In this case, if you print b, you will get an array of integers which are values of each byte. <br>
As Unicode characters like $ need more than one byte for their representation, the length of the byte slice might not be the same as the length of the string that is stores. Although the b byte slice has 12 characters, it has a size of 14!!

You can print the text of array of the byte with these commands.
```
fmt.Printf("Byte slice as text: %s\n", b)
fmt.Println("Byte slice as text:", string(b))
```

#### Deleting an element from a slice
```
aSlice = append(aSlice[:i], aSlice[i+1:]...)
```
Here we logically divide the original slice into two slices. The two slices are split at the index of the element that needs to be deleted. After that, we concatenate these two slices with the help of ...

#### How slices are connected to arrays
As mentioned before, behind the scenes, each slice is implemented using an underlying array. The length of the underlying array is the same as the capacity of the slice and there exist pointers that connect the slice elements to the appropriate array elements. 

You can understand that by connecting an existing array with a slice, Go allows you to reference an array of a part of an array using a slice. This has some strange capabilities including the fact that the changes to the slice affect the referenced array! However, when the capacity of the slice changes, the connection to the array ceases to exist! This happens because when the capacity of a slice changes, so does the underlying array, and the connection between the slice and the original arrya does not exist anymore!

```
a := [4]string{"Zero", "One", "Two", "Three"}
fmt.Println("a:", a)

// The result is => a: [Zero, One, Two, Three]

var S0 = a[0:1]
fmt.Println(S0)

// The result is => [Zero]

var S12 = a[1:3]
fmt.Println(S12)

// The result is => a: [One, Two]

S0[0] = "S0"
S12[0] = "S12_0"
S12[1] = "S12_1"

fmt.Println("a:", a)

// The result is => a: [S0, S12_0, S12_1, Three]
// The changes in S0 and S12 also change the contents of array. a // has not changedin a direct way. However, due to the connections // of a, the contents of a have changed!

// Adding 4 elements to S0
S0 = append(S0, "N1")
S0 = append(S0, "N2")
S0 = append(S0, "N3")
a[0] = "-N1"

// Changing the capacity of S0
// Not the same underlying array any more!
S0 = append(S0, "N4")

fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))
// The result is => Capacity of S0: 8 Length of S0: 5
// This change does not go to S0
a[0] = "-N500-"

fmt.Println("S0:", S0)
// The result is => S0:[-N1, N1, N2, N3, N4]
// As the capacity of S0 changes, it is no longer connected to the same underlying array. 

```

#### The copy() function

copy(destination, input)

The use of copy() can be tricky because the destination slice is not auto-expanded if the source slice is bigger than the destination slice. Additionally, if the destination slice is bigger than the source slice, then copy() does not empty the elements from the destination slice that did not get copied. 

```
a1 := []int{1}
a2 := []int{-1, -2}
a5 := []int{10, 11, 12, 13, 14}
fmt.Println("a1", a1)
fmt.Println("a2", a2)
fmt.Println("a5", a5)
// copy(destination, input)
// len(a2) > len(a1)
copy(a1, a2)
fmt.Println("a1", a1)
fmt.Println("a2", a2)

// The result is 
// a1 [-1] The first element of a1 is changed
// a2 [-1 -2]

// len(a5) > len(a1)
copy(a1, a5)
fmt.Println("a1", a1)
fmt.Println("a5", a5)

// The result is
// a1 [10] same as last scenario
// a5 [10 11 12 13 14]

// len(a2) < len(a5) -> OK
copy(a5, a2)
fmt.Println("a2", a2)
fmt.Println("a5", a5)

// The result is 
// a2 [-1 -2]
// a5 [-1 -2 12 13 14] Now first two elements of 15 changes and other elements, remained the same
```

#### Sorting Slices

The sort package can sort slices of built-in data types without the need to write any extra code. <br>
sort.Ints => sort integers <br>
sort.Reverse() => sorting in the reverse order than difault.

### Pointers

Go has support for pointers but not for pointer arithmetic, which is the cause of many bugs and errors in programming languages like C. A pointer is the memory address of a variable. You need to dereference a pointer in order to get its value dereferencing is performed using the * character in front of the pointer variable. Additionally, you can get the memory address of a normal variable using an & in front of it.

![Local Image](./Go%20pointers.png "Go pointers")

If a pointer variable points to an existing regular variable, then any changes you make to the stored value using the pointer variable will modify the regular variable.

Pointers are also very handy when you want to tell the difference between the zero value of a variable and a value that is not set (nil). This is particularly useful with structures because pointers can have the nil value, which means that you can compare a pointer to a structure with the nil value, which is not allowed for normal structure variables. 

\* The zero value for pointers is nil. when pointer variable points to nowhere, Go makes it point to nil. However, when we assign it to an empty structure value with new keyword, now this variable is no longer nil but fields of structure have the zero values of their data types.





