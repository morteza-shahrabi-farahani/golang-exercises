# Chapter 11: Code Testing and Profiling

Code optimization is the process where one or more developers try to make certain parts of a program run faster, be more efficient, or use fewer resources. 
Put simplly, code optimization is about eliminating the bottlenecks of a program that matter.

Code testing is about making sure that your code does what you want it to do.
The best time to write testing code is during development, as this can help to reveal bugs in the code as early as possible. 
Code profiling relates to measuring certain aspects of a program to get a detailed understanding of the way the code works. The results of code profiling may help you to decide which parts of your code need to change.

Have in mind that when writing code, we should focus on its correctness as well as other desirable properties such as readability, simplicity, and maintainability, not its performance. Once we are sure that the code is correct, then we might need to focus on its performance. A good trick on performance is to execute the code on machines that are going to be a bit slower than the ones that are going to be used in production.

## Optimizing code

There is no deterministic way to help you optimize your code and that you should use your brain and try many things if you want to make your code faster. However, the general principle regarding code optimization is first make it correct, then make it fast.

> The real problem is that programmers have spent far too much time worrying about efficiency in the wrong places and at the wrong times; premature optimization is the root of all evil (or at least most of it) in the programming.

Said by Donald Knuth

> Make it work, then make it beautiful, then if you really, really have to, make it fast. 90 percent of the time, if you make it beautiful, it will already be fast. So really, just make it beautiful.

Said by Joe Armstron,g one of the developers of Erlang.

## Benchmarking code 

Benchmarking measures the performance of a function or program, allowing you to compare implementations and to understand the performance impact of code changes. Using that information, you can easily reveal the part of the code that needs to be rewritten to improve its performance. 

\* Most of the time, the load of the operating system plays a key role in the performance of your code. 

Go follows certain conventions reggarding benchmarking. The most important convention is that the name of a benchmark function must begin with Benchmark. After the Benchmark word, we can put an underscore or an uppercase letter. The same rule applies to testing functions that begin with Test. By convention such functions are put in files that end with _test.go. Once the benchmarking or the testing code is correct, the go test subcommand does all the dirty work for you.

\* Benchmark functions use testing.B variables whereas testing functions use testing.T variables. 

