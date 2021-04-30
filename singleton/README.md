# SINGLETON PATTERN

Ensures that a class has only one instance and provides a global point of access to it.

__Class Diagram__

![Figure](cd.drawio.svg)

Singleton Pattern comes into play where having multiple object would lead to chaos. We would run into problems like incorrect programming behavior, overuse of resources or inconsistent results. There are many places where we need only one object thread pools, caches, objects used for logging etc.

We need to keep oour implementation thread safe else it might happen that two threads are using different instances of our Singleton class. The most simple approach would be to keep a lock in the `getInstance()` method, hence it can be accessed by only one thread at a time.

In our implementation we are using Go's [`func (*Once) Do(f func())`](https://golang.org/pkg/sync/#Once.Do). Do calls the function `f` if and only if `Do` is being called for the first time for this instance of `Once`. Hence, in this way we ensure that our Singleton's class object is created only once even when multiple go routines are calling it.
