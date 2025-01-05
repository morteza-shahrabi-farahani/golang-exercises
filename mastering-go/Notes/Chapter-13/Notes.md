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

### Creating constraints

```
type Numeric interface {
    type int, int8, int16, int32, int64, float64
}
```

In here, we define a new interface called Numeric that specifies the list of supported data types. You can use any data type you want as long as it can be used with the generic function that you are going to implement. In this case, we could have added string or uint to the list of supported data types.

```
func Add[T Numeric](a, b T) T {
    return a + b
}
```

## Interfaces versus generics
When using interfaces, we must using a type switch to differentiate between the supported data types.

```
func Print(s interface{}) {
    switch s.(type) {
        case int:
            fmt.Println(s.(int)+1)
    }
}
```

The biggest issue with Print() is that due to the use of the empty interface, it accepts all kinds of input. As a result, the function signature does not help us limit the allowed data types. The second issue with Print() is that we need to specifically handle each case.

On the other hand, the compiler does not have to guess many things with that code, which is not the case with generics, where the compiler and the runtime have more work to do. This kind of work intruduces delays in the execution time.

```
func PrintGenerics[T any] (s T) {
    fmt.Println(s)
}
```

PrintGenerics() is a generic function that can handle all available data types simply and elegantly.

## Summary

Although a function with generics is more flexible, code with generics usually runs slower than code that works with predefined static data types. So, the price you pay for flexibility is execution speed. Similarly, Go code with generics has a bigger compilation time than equivalent code that does not use generics.

At the end of the day, programming is about understanding the cost of your decisions. Only then can you consider yourself a programmer. So, understanding the cost of using generics instead of interfaces, reflection, or other techniques is important.



