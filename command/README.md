
# COMMAND PATTERN

__Command Pattern__ is a behavioural design pattern in which an object is used to encapsulate all the information needed to trigger a event at a later time.


The Command pattern works with the following terms — Command, Receiver, Invoker, and Client.

Let's understand this with a simple example. Consider the working of a restaurant.
- A customer gives a order: `createOrder()`
- A waiter takes the order: `takeOrder()`  
- Waiter gives the order to chef: `orderUp()`
- Order contains what needs to be prepared: `makeBurger()`
- Chef prepares what is written in the order: `Output`

In this if you notice, the waiter does not need to know whats in the order or how to prepare it. 
It just knows that it needs to take the order and invoke the chef. 
The chef in the other hand knows how to deal with the order.

Mapping this to our pattern terms:
- Client -> customer
- Command -> order
- Invoker -> Waiter
- Receiver -> Chef

The client wants to execute a set of commands. 
It creates the command and passes on to the invoker.
The invoker stores the command until it is needed. 
When the time comes, the invoker executes the commands.
The receiver then receives these commands and takes the action.

But like the waiter does not know how to prepare the order, the invoker does not know how to execute the commands or which receiver should receive this. 
It is stateless.
To solve this problem, we package the commands with the receivers into an object that exposes just one method, `exexute()`.

Refer code for better underestanding.

__Class Diagram__

![Figure](cd.drawio.svg)

### Salient Feautes
- Encapsulating a request as an object.
- Allowing the parameterization of clients with different requests.
- Decoupling of the invoker and receiver. Invoker does not know what actions get performed on what receiver. It just knows that if it calls the `execute()` method, the request will be serviced.

### Undo 
To add the Undo functionality:
- Command should have an `undo()` method that mirrors the `execute()`. 
For example, if the command `LightOnCmd`'s execute() method turns on the light.
Then the undo() method should turn off the light. 
This looks simple when our command is binary, but for more complicated commands we also need to store the previous state of the receiver.
- In the invoker class we need to introduce another variable, `prevCommand` which tracks the last command to get invoked.

### No Command
In the invoker's constructor, we can initialize the command with a `NoCommand`. 
This way we always have some command to call.
We don't have to check whether the command is Nil when using an invoker object.
The NoCommand's `execute()` can be kept empty.

### Macro Command
Lets say we need to create a command which lets us execute several other commands.
We call this multipurpose command a macro command. 
It will also implement the command interface.
It will have list of commands `commands []Command` as an attribute.
The constructor will fill this with the commands needed.
The `execute()` and `undo()` methods will loop over the commands list and call each command's execute() or undo() respectively.

### Use Cases
Since we can create a pool of requests and this requests can be invoked by different thread as well, this application can have several usecases such as:
- Schedulers
- Job queues
- Thread pools

One more use case can be logging of requests. 
Each command executed in a computer can be stored in a database.
And if the computer crashes we can execute all commands from database to restore the computer.

### Problem Statement
We will try to implement a stock exchange with the following features:
- User can buy or sell a stock
- Broker needs to execute the order
- Additionally, support undo of the previous order.
