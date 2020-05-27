# goAdvanced

concurrency : multiple computations happning same time (when 2 task are overlapping is concurrency )  its a design pattern

race condition : 
A race condition is when two or more routines have access to the same resource, such as a variable or data structure and attempt to read and write to that resource without any regard to the other routines. This type of code can create the craziest and most random bugs you have ever seen. 

mutex : to overcome race conditon :
A Mutex, or a mutual exclusion is a mechanism that allows us to prevent concurrent processes from entering a critical section of data whilst it’s already being executed by a given process.


















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


func init() {} - can be added to any go file , you can add any number of init () function in a single go file ( you can not explicitly call init()) 

youtube : = rob pike concurrency parallelism

packages: runtime, "golang.org/x/crypto/bcrypt"