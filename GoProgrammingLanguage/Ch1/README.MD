# Notes

GoLang is a compiled language, it has a compuler in which you need to build the executable first before running the program. Go does this by using `go build` mechanism. With Go, there is also another way to shortened the build and execute the program immediately, this is done via the `go run` command.

## 1.1
- String assignment is held in memory. And every new string value the variable is assigned will adds to the memory. This will be garbage collected in due time according to GoLang's garbage collection. Having said that, it means that multiple occurrence of string assignment can be quite expensive.
- Slices is another way of creating a list in Strings. Once you have created a slice, you can access the given characters via its index. S[i]

## 1.2
- Map holds a set of key/value pair. Map is similar to a dictionary that it provides constant-time operations to store, retrieve or test for an item in the set. O(1) Set, Retrieve, Check
- Printf produce formatted output from expressions.

## 1.3
- Using import packages of IO and Scanner. IO Packages behaves in a way that you can read text file and open it. While, Scanner read lines via a stream, which allow for a better memory configuration.

## 1.4
- Packages in Golang is known by its end packages name. For example to call packages io/ioutil, we need to use the ioutil.ReadAll()

## 1.5
- Packages "net" is an important utility that handles resources with the internet and wider net. It has built in tools to serve a server, to handle HTTP request and all the other things related to the net. For more information: https://pkg.go.dev/net 

## 1.6
- Golang has a large support of concurrent programming. Concurrent Programming is the idea that a program can execute tasks in an interleaved fashion through context switching. Context switching allows the task to run in interleaved process in a way that it is non blocking. In this case, the fetch program that we have can concurrently retrieve multiple resources from the internet.
  - processor Core1 (Context Switching) -> fetch -> URL A, URL B, URL C
  - processor Core1 (Context Switching) -> fetch -> URL A(Finished), URL B, URL C(Finished)
- Comparing the two terminology, there is also the concept of parallel programming, which is different to Concurrent Programming. Parallel programming is the idea that you can simultaneously execute process through different CPU Cores simultaneously. Consider this as an example:
  - processor Core 1- task 1
  - processor Core 2- task 2
  - processor Core 3- task 3
- In comparing these two idea, the ability to perform concurrent programming allow you to complete the execution of the job much faster compared to executing sequentially.
