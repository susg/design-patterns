# DESIGN PATTERNS

Implementation of commonly used design patterns in Go.

## How to read the code

Although the code is self explanatory, but still here are some tips.
- Each go package represents a design pattern.
- Each package has a `README.md` file which briefly describes the design pattern. It also has a problem statement which the code is trying to solve. Going through the class diagram will give a general idea of the flow of data in the code.
- Packages do not have a `main()` method. Instead the objects are instantiated in the tests. So please go through tests as that is the entry point for each package.

## Run tests

Enter into the repo and type:

```bash
go test -v ./...
```

This will run all tests of every packages.
