# Template Method Pattern

 The Template Method Pattern defines the skeleton of an algorithm in a method, deferring some steps to subclasses. 
 Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm’s structure.

 __Class Diagram__

![Figure](cd.drawio.svg)

### Hook
A hook is a method declared in the abstract class, but only given an empty or a default implementation.
This gives subclasses the ability to “hook into” the algorithm at various points, if they wish; a subclass is also free to ignore the hook.
give the subclass
Hook's also give a chance to react to some step in the template method that is about to happen, or just happened.

### Dependency Rot
Dependency rot happens when you have high-level components depending on low-level components
depending on high-level components depending on sideways components depending on low-level components, and so on.
Hooks help to solve this problem. We allow low level components to hook themselves into a system but the high level components determine when they are needed and how.

### Template VS Strategy Pattern
Both the patterns are used to encapsulate the algorithm.
But the intent of Template Pattern is to define the outline of an algorithm 
and let its subclasses do some of the work.
This way there can be different implementations of the algorithms individual steps but the overall structure remains the same.
Whereas the Strategy Pattern gives the client the choice of algorithm implementation through composition.

### Problem Statement
A tour and travels company need to prepare the itinerary of the tour by 
taking inputs from the customer.
