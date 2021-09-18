# Adapter Pattern

The __Adapter Pattern__ converts the interface of a class into another interface the clients expect.
Adapter lets classes work together that couldn’t otherwise because of incompatible interfaces.

![Figre](adptr.png)

__Class Diagram__

![Figure](cd.drawio.svg)

### Problem Statement
We have two kinds of port: Mac Port and Windows Port and both provide different method to insert device.
The client can only use Mac Port. Help the client use windows port as well.
