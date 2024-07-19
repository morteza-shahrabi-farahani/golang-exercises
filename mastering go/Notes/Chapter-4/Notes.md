# Chapter 4: Reflection and Interfaces

Interfaces are about expressing abstractions and identifying and defining behaviors that can be shared among different data types. Once you have implemented an interface for a data type, a new world of functionality becomes available to the variables and the values of that type, which can save you time and increase your productivity. Interfaces work with methods on types or type methods, which are like functions attached to given data types, which in Go are usually structures. Remember that once you implement the required type methods of an interface, that interface is satisfied implicitly, which is also the case with the empty interface that is explained in this chapter. 

Another handy Go feature is reflection, which allows you to examine the structure of a data type at execution time. However, as reflection is an advanced Go feature, you do not need to use it on a regular basis. 

## Reflection
If you want to find out the names of the fields of a structure at execution time, you need to use reflection. Apart from that, reflection also allows you to explore and manipulate unknown structures like the ones created from decoding JSON data.

Reflection allows you to dynamically learn the type of an arbitrary object along with information about its structure. Additionally, eflection might come in handy when you have to work with data types that do not implement a common interface and therefore have an uncommon or unknown behavior.

reflect.Value is used for storing values of any type, whereas reflect.Type is used for representing Go types. There exist two functions named reflect.TypeOf() and reflect.ValueOf() that return the reflect.Type and reflect.Value values, respectively. Note that reflect.TypeOf() returns the actual type of variable - if we are examining a structure, it returns the name of the structure. <br> reflect.NumField() => listing the number of fields in a structure. <br> reflect.Field() => getting reflect.Value value of a specific field of a structure. <br> reflect.King => representing the specific data type of a variable: int, struct, etc.

\* Reflection code can look unpleasant and hard to read sometimes. Therefore, according to the Go philosophy, you should rarely use reflection unless it is absolutely necessary because despite its cleverness, it does not create clean code. 

example usage:
```
type Secret struct {
Username string
Password string
}

type Record struct {
Field1 string
Field2 float64
Field3 Secret
}

r := reflect.ValueOf(A)
fmt.Println("String value:", r.String())
=> String value: <main.Record Value>

iType := r.Type()
fmt.Printf("i Type: %s\n", iType)
fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)
=> i Type: main.Record
=> The 3 fields of main.Record are

for i := 0; i < r.NumField(); i++ {
    fmt.Printf("\t%s ", iType.Field(i).Name)
    fmt.Printf("\twith type: %s ", r.Field(i).Type())
    fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())
    => Field1 with type: string and value _String value_
    => Field2 with type: float64 and value _-12.123_
    => Field3 with type: main.Secret and value _{Mihalis Tsoukalos}_

    // Check whether there are other structures embedded in Record
    k := reflect.TypeOf(r.Field(i).Interface()).Kind()
    // Need to convert it to string in order to compare it
    if k == reflect.Struct {
        fmt.Println(r.Field(i).Type())
        => main.Secret
    }
}

```

If you were to make changes to the values of the structure fields, you would use the Elem()  method and pass the structure as a pointer to ValueOf(). There exist methods that allow you to modify an existing value. For example SetString() and SetInt().

### The three disadvantages of reflection
* The first reason is that extensive use of reflection will make your programs hard to read and maintain.
* The second reason is that the Go code that uses reflection makes your programs slower. Generally speaking, Go code that works with a particular data type is always faster than Go code that uses reflection to dynamically work with any Go data type. Additionally, such dynamic code makes it difficult for tools to refactor or analyze your code.
* The last reason is that reflection errors cannot be caught at build time and are reported at runtime as panics, which means that reflection errors can potentially crash your programs.