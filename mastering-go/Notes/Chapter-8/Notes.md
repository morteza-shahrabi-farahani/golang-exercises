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

## Implementing the handlers
Usually, handlers are put in a separate package.

```
func deleteHandler(w http.ResponseWriter, r *http.Request) {
    // Get telephone
    paramStr := strings.Split(r.URL.Path, "/")
    fmt.Println("Path:", paramStr)
    if len(paramStr) < 3 {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintln(w, "Not found: "+r.URL.Path)
        return
    }
```

If we do no have enough parameters, we should send an error message back to the client with the desired HTTP code, which in this case is http.StatusNotFound.

```
log.Println("Serving:", r.URL.Path, "from", r.Host)
```

This is where the HTTP server sends data to log files - this mainly happens for debugging reasons.

```
mux := http.NewServeMux()
s := &http.Server{
    Addr: PORT,
    Handler: mux,
    IdleTimeout: 10 * time.Second,
    ReadTimeout: time.Second,
    WriteTimeout: time.Second,
}

mux.Handle("/list", http.HandlerFunc(listHandler))
mux.Handle("/insert/", http.HandlerFunc(insertHandler))
mux.Handle("/insert", http.HandlerFunc(insertHandler))

```

Here, we store the parameters of the HTTP server in the http.Server structure and use our own http.NewServeMux() instead of the default one.

```
err = s.ListenAndServe()
if err != nil {
    fmt.Println(err)
    return
}
```
The ListenAndServe() method starts the HTTP server using the parameters defined previously in the http.Server structure.

\* The http package uses multiple goroutines for interacting with clients - in practice, this means that you application runs concurrently!

## Exposing metrics to Prometeus
The list of supported data types for metrics is the following:

Counter: Counters are usually used for representing cumulative values such as the number of requests served so far, the total number of errors, etc.

Gauge: Gauges are usually used for representing values that can go up or down such as the number of requests, time durations, etc.

Histogram: A histogram is used for sampling observations and creating counts and buckets. Histograms are usually used for counting request durations, response times, etc. 

Summary: A summary is like a histogram but can also calculate quantiles over sliding windows that work with times.

The runtime/metrics package makes metrics exported by the Go runtime available to the developer. If you want to collect all available metrics, you should use metrics.All().

\* You might ask, "why not use a normal Go binary instead of a Docker image?" The answer is simple: Docker images can be put in docker-compose.yml files and can be deployed using Kubernetes. The same is not true about Go binaries.






