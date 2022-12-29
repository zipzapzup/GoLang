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
- In comparing these two idea, the ability to perform concurrent programming allow you to complete the execution of the job much faster compared to executing sequentially. In this lesson, this is called FetchALL.
- In Chapter 6 we learn about go routine and channel. Channel is the communications mechanism or channel into which allow go routine to pass values to antoher go routine. In the fetchAll program, we can spot that the main function is a go routine that creates a channel. Main runs a loop of other go statements which creates additional goroutines.
  - main -> 1 go routine and create 1 channel
    - loop go 1 sends message to ch
    - loop go 2 sends message to ch
    - loop go 3 sends message to ch
- Note how in the execution of fetch all URL, the print command fmt.Println(<-ch) will print on whichever go routine finish first. ch <- denotes that we are sending a message to the channel and <- ch means this function will receive all of the channel.
 
## 1.7
- In Chapter 1.7 these present introductions to web server and how through the http library, we can serve the request to a handler via the DefaultServeMux pattern matching.

## 1.8
- Switch is one form of control flow in golang aside from if statement and for loops. Switch has the same usage as a case in bash and you can use switch using two format, where in one format you can specify a function and switch will wait for the functions return value OR a tagless switch inside a function, where it is generally assumed switch is true.:
  - switch funcA() { case "heads": heads++; case "tails" L tails++}
  - switch {case x> 5: x++; case x <=5: x--}
- break and continue is another statements in golang, where it is used in control flow (for loops, ifs). 
  - Break is used to break the loop and move to the next statement.
  - Continue is used to continue the iteration of the loop, moving to the next iteration.
- Type is used to create a structure in golang programs.


- Pointers: Golang provides pointers as a way to use them as references or use them to make changes to the value. There are 2 different types of pointers
  - & operator is a pointer that will return the address of a variable.
  - * operator retrieves the variable that the pointer refers to.
- Methods and Interfaces: Methods is a function that is associated with a named type. Interfaces are abstract types that allow a different treatment in the same way based on what methods they have.
- Packages in golang refers to as modules. You can find and try to get yourselves familiarised with golang standard library: https://golang.org/pkg or pkg.go.dev/std. Go doc is also an important tools to view documentation, one example is: go doc http.ListenAndServe.


## 2
- Go Lang builds large programs from a small set of basic constructus. Variables store the values and simple expressions are combined into larger ones. In Golang, these are the structural elements of a Go Program.

## 2.1
2.1 Names: are a way to define variables in Golang, it is case sensitive and any name needs to begin with letters. In Golang, names have different declaration scope.
    - function scope: when an entity is defined within the function scope. It will be a local scope in the fucntion of where it belongs.
    - package scope: To define an entity that is exportable outside the package scope, you need to define the entity starting with an upper case. For example "PersonAge" is exportable outside the packages compared to "personAge".

## 2.2
Declaration: A declaration create a program entity and specify its properties. In Golang, there are 4 concepts of declarations:
    - var: variables, which are names that can have its value changed from time to time. Variables are mutables.
    - const: a constant variables in which it is an immutable values.
    - type: defining your own type and structure.
    - func: declaring a function.

## 2.3 Variables
Variables is a way for you to define variables in Golang. Declarations follows the format of `var name type = expression` to define any variable.

In golang, there are the concept of zero variables, in which any variables that is declared and not assigned, it will be given its zero value.

Golang also allow you to declare multiple variables at the same time known as initialisers.

Lifetime of variables depends on whether the variables are traceable by Golang. If its not traceable by Golang, then it will be cleaned up in due time by its Garbage collection. Defining the lifetime of a variable is different and it depends on whether the variable is declared inside a function (local) or global. Global variable tends to have a longer lifetime since its not a type of stack that is not used after the function is created and returned.

2 Main differences between int and *int

- Int in golang is a type of integers, which means that the variable wil lreturn the value of int. Zero value of int is 0.
- *int on the other hand represents the memory address of int and its a pointer type of int. Zero value of *int is <nil>.

var x *int is a pointer of type int
  - returns memory address
var x int is a variable of type int 
  - returns value


## 2.4 Assignments

A few different ways to assign variables exist in Golang. One of the ways is to assign it: `var x int = 10`. In Golang, you can also do `x *= 10` to apply an assignment operator, where you apply the value of x by another value and assign it simultaneously.

Tuple assignment is also allowed in golang: `x, y = y, x`. To perform assignments, both the variables on the left hand side and the right hand side must have the same types. If the types are not the same, then it will error out.

## 2.5 Type Declarations
In golang, you can declare your own type and perform your own type conversions. Once you performed the type conversions, you can then perform operations or comparison operators. Comparison is allowed for those of the same types in Golang, note the type needs to be the same. 

Operations can be performed on the same type. 

## 2.6 Packages and Files

Packages is an important concept in Golang that support libraries or modules. It is a way to support modularity, encapsulation, separate compilation and reuse. Source code for package are usually in a directory that ends with the import path.

Package are stored in directory: `$GOPATH/src/gopl.io/ch1/helloworld` 

Each package serves as a separate name space for its declarations. Within image package, for example, the identifier refers to different function than in the other package. Concept of namespaces exist in Golang packages.

