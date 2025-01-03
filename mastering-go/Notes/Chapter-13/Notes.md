# Chapter 13: Go Generics

Let me make something clear from the beginning; you do not have to use Go generics if you do not want to and you can still write wonderful, efficient, maintainable, and correct software in Go! Additionally, the fact that you can use generics and support lots of data types, if not all available data types, does not mean that you should do that. Always support the required data types, no more, no less, but do not forget to keep an eye on the future of your data and the possibility of supporting data types that were not known at the time of writing your code.

## Introducing generics

Generics are a feature that gives you the capability of not precisely specifying the data type of one or more function parameters, mainly because you want to make your functions as generic as possible. In other words, generics allow ffunctions to process several data types without the need to write special code, as is the case with the empty interface or interfaces in general. However, when working with interfaces in Go, you have to write extra code to determine the data type of the interface variable you are working with, which is not the case with generics.

