# 如何写 Go 代码

## 代码组织
- A package is a collection of source files in the same directory that are compiled together.
- A module is a collection of related Go packages that are released together.
- A file named `go.mod` located at the root of the repository declares the module path: 
the import path prefix for all packages within the module.
- the module `github.com/google/go-cmp` contains a package in the directory `cmp/`. 
That package's import path is `github.com/google/go-cmp/cmp`.
- Packages in the standard library do not have a module path prefix.

## First Program

