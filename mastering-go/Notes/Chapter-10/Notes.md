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

## Using gorilla/mux 
The gorilla/mux package is a popular and powerful alternative to the default Go router that allows you to match incoming requests to their respective handler.

r.HandleFunc("/url", UrlHandlerFunction).Methods(http.MethodPut):
This example shows how you can tell Gorilla to match a specific HTTP method, which saves you from having to write code to do that manually. 

s.HandleFunc("/users/{id:[0-9]+}"), HandlerFunction): This last
example shows that you can define a variable in a path using a name (id)
and a pattern and Gorilla does the matching for you! If there is not a regular
expression, then the match is going to be anything from the beginning slash
to the next slash in the path.

### The use of subrouters
A subrouter is a nested route that will only be examined for potential matches if the parent route matches the parameters of the subrouter. The good thing is that the paent route can contain conditions that are common among all paths that ae defined under a subrouter, which includes hots, path prefixes, and, as it happens in our case, HTTP request methods.

## Working with multiple REST API versions

There exist various approaches on how to implement REST API versioning, including the following:

* Using a custom HTTP header (version-used) to define the used version

* Using a different subdomain for each version (v1.servername and v2.servername)

* Using a combination of Accept and Content-Type headers -this method is based on content negotiation

* Using a different path for each version (/v1 and /v2 if the RESTful server supports two REST API versions.)

* Using a query parameter to reference the desired version (.../endpoint?version=v1 or .../endpoint?v=1)

There is no correct answer for how to implement REST API versioning. Use what seems more natural to youou and your users.

## Uploading and downloading binary files

There exist three main ways to save the files you upload:

* On the local filesystem

* On a database management system that supports the storing of binary files.

* On the clouad using a cloud provider

\* If we had a scenario which we can have a function which we can reuse it in other parts of the program, it would better to seperate the code scope of this function for reusability.

