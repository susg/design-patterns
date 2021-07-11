# DESIGN PATTERNS |  [![Go](https://github.com/susg/design-patterns/actions/workflows/go.yml/badge.svg)](https://github.com/susg/design-patterns/actions/workflows/go.yml) [![CodeFactor](https://www.codefactor.io/repository/github/susg/design-patterns/badge)](https://www.codefactor.io/repository/github/susg/design-patterns) [![codecov](https://codecov.io/gh/susg/design-patterns/branch/master/graph/badge.svg?token=3AVWKA89AJ)](https://codecov.io/gh/susg/design-patterns) [![Open in Visual Studio Code](https://open.vscode.dev/badges/open-in-vscode.svg)](https://open.vscode.dev/organization/repository)

Implementation of commonly used design patterns in Go.

## How to read the code

Although the code is self-explanatory, but still here are some tips.
- Each go package represents a design pattern.
- Each package has a `README.md` file which briefly describes the design pattern. It also has a problem statement which the code is trying to solve. Going through the class diagram will give a general idea of the flow of data in the code.
- Packages do not have a `main()` method. Instead, the objects are instantiated in the tests. So please go through tests as that is the entry point for each package.

## Run tests

Enter into the repo and type:

```bash
go test -v ./...
```

This will run all tests of every package.
