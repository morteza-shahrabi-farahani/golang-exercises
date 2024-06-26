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


