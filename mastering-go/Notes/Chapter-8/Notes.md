# Chapter 8: Building Web Services
The http.Response structure embodies the response from an HTTP request. Both http.Client and http.Transport return http.Response values once the response headers have been received.

The http.Request structure represents an HTTP request as constructed by a client in order to be sent or received by an HTTP server.

## Creating a web server
The net/http package offers functions and data types that allow you to develop powerful web servers and clients. The http.Set() and http.Get() methods can be used to make HTTP and HTTPS requests, whereas http.ListenAndServe() is used for creating web servers given the user-specified handeler function or functions that handle incoming requests.

The simplest way to define the supported endpoints, as well as the handler function that responds to each client request, is with the use of http.HandleFunc(), which can be called multiple times.

```
func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
    fmt.Printf("Served: %s\n", r.Host)
}

http.HandleFunc("/time", timeHandler)
http.HandleFunc("/", myHandler)
```

The http.ListenAndServe() call begins the HTTP server using the predefined port number.

```
err := http.ListenAndServe(PORT, nil)
if err != nil {
    fmt.Println(err)
    return
}
```