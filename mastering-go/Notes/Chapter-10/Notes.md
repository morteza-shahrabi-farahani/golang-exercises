# Chapter 10: Working with REST APIs
REST is an acronyum for Representational State Transfer and is primarily an architecture for designing web services. REST is not tied to any operating system or system architecture and is not a protocol; however, to implement a RESTful service, you need to use a protocol such as HTTP.

## An introduction to REST
Most modern web applications work by exposing their APIs and allowing clients to use these APIs to interact and communicate with them. Although REST is not tied to HTTP, most web services use HTTP as their underlying protocol. Additionally, although REST can work with any data format, usually REST means JSON over HTTP because most of the time, data is exchanged in JSON format in RESTful services.

Due to the way a RESTful service works, it should have an architecture that follows the next principles:

* Client-server design
* Stateless implementation - this means that each interaction does not depend on others
* Cacheable
* Uniform interface
* Layered system

/* As a convention, a PUT request should contain the full and updated version of an existing resource. A PATCH request only contains the modifications to an existing resource.

## A RESTful server
```
if r.Method != http.MethodPOST {
    http.Error(w, "Error:", http.StatusMethodNotAllowed)
    fmt.Fprintf(w, "%s\n", "Method not allowed!")
    return
}
```

The http.Error() function sends a reply to the client request that includes the specified error message, which should be in plain text, as well as the desired HTTP code.