# Chapter 13: Go Generics

Let me make something clear from the beginning; you do not have to use Go generics if you do not want to and you can still write wonderful, efficient, maintainable, and correct software in Go! Additionally, the fact that you can use generics and support lots of data types, if not all available data types, does not mean that you should do that. Always support the required data types, no more, no less, but do not forget to keep an eye on the future of your data and the possibility of supporting data types that were not known at the time of writing your code.

## Introducing generics

Generics are a feature that gives you the capability of not precisely specifying the data type of one or more function parameters, mainly because you want to make your functions as generic as possible. In other words, generics allow ffunctions to process several data types without the need to write special code, as is the case with the empty interface or interfaces in general. However, when working with interfaces in Go, you have to write extra code to determine the data type of the interface variable you are working with, which is not the case with generics.

```
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
```

There is a function that accepts a slice of any data type. This is denoted by the use of []T in the function signature in combination with the [T any] part. The [T any] part tells the compiler that the data type T is going to be determined at execution time. We are also free to use multiple data types using the [T, U, W any] notation.

The any keyword tells the compiler that there are no constraints about the data type of T.

Now imagine writing separate functions to implement the functionality of this function for slices of integers, strings, floating-point numbers and so on. So, we have found a profound case where using generics simplifies the code and our programming efforts. However, not all cases are so obvious, and we should be very careful about overusing any.

## Constraints

Let us say that you have a function that works with generics that multiplies two numeric values. Should this function work with all data types? Can this function work with all data types? Can you multiply two strings or two structures?The solution for aboiding that kind of issue is the use of contraints.

```
func Same[T comparable](a, b T) bool {
    if a == b {
        return true
    }

    return false
}
```

The same() function uses the predefined comparable constraint insteasd of any. In reality, the comparable constraint is just a predefined interface that includes all data types that can be compared with == or !=.

