Go Concepts

1. Statically typed language, produces binary executables like C++.
2. Go has built in support for concurrency through goroutines and channels. Goroutine are lightweight threads managed by Go runtime, allowing for efficient concurrent execution of code. Channels faciliate communication and synchronization between goroutines, enabling developers to write highly concurrent programs without overhead typically associated with threading in other languages.
3. Go's garbage collector is designed to minimize pause times and memeory overhead. It uses a concurrent, tri-color mark and sweep algorithm, which allows garbage collection to be performed concurrently with the execution of Go code.
4. Cross platform (runs on linux/mac/windows). It means the ability to compile code to run on different os and architecture.
5. Companies that uses Golang: Dropbox, Uber, Alibaba, Paypal

6. Go module is to initialize ur project in go, its like a collection of packages, it defines all the dependencies which we install in our project.
7. Named return parameter and variadic functions (variadic function accpets variable number of arguments of a specific type, it is decalred by using an ellipsis as suffix)
8. Naming a function with uppercase letter meaning it is exported
