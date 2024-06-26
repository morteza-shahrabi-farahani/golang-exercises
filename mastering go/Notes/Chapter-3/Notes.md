# Chapter 3: Composite Data Types

## Map
\* You should make no assumptions about the order of the elements inside a map. Go randomizes keys when iterating over a map - this is done no purpose and is an intentional part of the language design. 

\* You can delete a key and value pair from a map using the delete() function. Additionally, you can tell whether a key k exists on a map named aMap by the second return value of the v, ok := aMap[k] statement. If it does not exist, v will be set to the zero value of its data type. If you try to get the value of a key that does not exist in a map, Go will not complain about it and returns the zero value of the data type of the value. 

\* Put simply, if you try to store data on a nil map, your program will crash. Testing whether a map points to nil before using it is a good practice. 

### Iterating over maps
When for is combined with the range keyword it implements the functionality of foreach loops found in other programming languages and allows you to iterate over all the elements of a map without knowing its size or its keys. 

```
for key, v := range aMap {
    fmt.Println("key", key, " value", value)
}
```

As you already know, you should make no assumptions about the order that the key and value pairs of a map will be returned in from a for and range loop.

## Structures
The type keyword allows you to define new data types or create aliases for existing ones. Therefore, you are allowed to say type myInt int and define a new data type called myInt that is an alias for int. However, Go considers myInt and int as totally different data types that you cannot compare directly even though they hold the same kind of values. Each structure defines a new data type, hence the use of the type keyword. 

The order in which you put the fields in the definition of a structure type is significant for the type identity of the defined structure. Put simply, two structures with the same fields will not be considered identical in Go if their fields are not in the same order.

\* The new() keyword has the following properties: <br>
* It allocates the proper memory space, which depends on the data type, and then it zeroes it
* It always returns a pointer to the allocated memory
* It works for all data types except channel and map

\* If no initial value is given to a variable, the Go compiler automatically initializes that variable to the zero value of its data type. For structures, this means that a structure variable without an initial value is initialized to the zero values of each one of the data types of its fields.


