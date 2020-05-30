# goAdvanced

concurrency : multiple computations happning same time (when 2 task are overlapping is concurrency )  its a design pattern , allow to run multiple go routine at the same time.

race condition : 
A race condition is when two or more routines have access to the same resource, such as a variable or data structure and attempt to read and write to that resource without any regard to the other routines. This type of code can create the craziest and most random bugs you have ever seen. 

mutex : to overcome race conditon :
A Mutex, or a mutual exclusion is a mechanism that allows us to prevent concurrent processes from entering a critical section of data whilst it’s already being executed by a given process.

Atomic : used for synchronize algorithm used to avoid race condition

Semaphore: a semaphore is a variable or abstract data type used to control access to a common resource by multiple processes in a concurrent system such as a multitasking operating system.

chan (channel) :=  You can not use send and receive at the same time or at the same function. 
                c := make(chan int)
                c <- 10 ///send
                 <- c     //receive

                 directional channel :
                 c := make(chan <- int) 
                c<- 10 
                <- c  //not possible
chanels : are like relay racer , you put on it and it stops until some one takes it (unless you using buffered channel). 
when called close() function you can not put value to channel but you can read from it.

if range is used then you need to use close in the channel:
     c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				c <- i

			}
			close(c)
		}()
	}

	for c1 := range c { // with the range code block here. until it gets values.
		fmt.Println(c1)
	}

method set -  A type may have a method associated with it is called method set 
            abc type int 
            func (t *abc) printMe() 
            func (t abc) printMe()  
            Receiver  value 
            -------------------
            (t T)      T and *T
            (t T)         *T

func init() {} - can be added to any go file , you can add any number of init () function in a single go file ( you can not explicitly call init()) 

packages: runtime, "golang.org/x/crypto/bcrypt"

func Gosched(): (package runtime)
how to yield : allow other routines to run  

context pattern in go : - Request handlers often start additional go routines to access backends such as databases  and RPC services.
When a request is canceled or timed out, all the routines working on that request need to close quickly. We have context package for that.
withCanel : arranges for done to be closed when cancel called 
withDedline: arragnges for done to be closed when deadline called 


A select is only used with channels.

A switch is used with concrete types.

A select will choose multiple valid options at random, while aswitch will go in sequence (and would require a fallthrough to match multiple.)

Note that a switch can also go over types for interfaces when used with the keyword .(type)

var a interface{}
a = 5
switch a.(type) {
case int:
     fmt.Println("an int.")
case int32:
     fmt.Println("an int32.")



Error handling :
Do does not have exceptions (as goes multi value return (from function ) makes it easy to report error without overloading the return values)

defer: you can add as many defer as you can , it will be executed in revers sequences.
         if program exists before defer function call statement position sequence, then it wont call defer (LIFO)
         you can not use combination of both go and defer keywoard

error is just another type (interface). 

type error interface {
     Error() string
}
if you want to create custom error type you need to implement Error() method.

errors packages: (implements error interface)
it has a func New(text string) error  

Panic(): a built-in function that stops the ordinary flow of control and beginging panicking. when a function calls panic , execution stops , any defered function registered will be executed.

Recover is the built-in function that regains control of a panicking routine. useful in the defered function , useful to resume call.




Important resource :
https://godoc.org/golang.org
https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
https://divan.dev/posts/go_concurrency_visualize/
https://medium.com/hackernoon/dancing-with-go-s-mutexes-92407ae927bf#.wjr1u2xjm
https://stackoverflow.com/questions/34197248/how-can-i-store-reference-to-the-result-of-an-operation-in-go
http://golangtutorials.blogspot.com/2011/06/memory-variables-in-memory-and-pointers.html

https://github.com/GoesToEleven/go-programming

https://github.com/GoesToEleven/GolangTraining


youtube : = rob pike concurrency parallelism


toDo : 

fanout pattern in go 
