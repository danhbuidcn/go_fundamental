# go_fundamental

This repository contains resources and code examples for learning the Go programming language. The `README` provides instructions for setting up and running Go, along with links to different folders that cover various aspects of Go programming.

The Go language was developed with the criteria of simplicity, efficiency, and ease of maintenance, aiming at writing clear and fast code.

## Table of Contents
- [Project Overview](#project-overview)
- [Setup](#setup)
- [Folder Structure](#folder-structure)
- [How to Run Go Code](#how-to-run-go-code)
- [Learning Resources](#learning-resources)

## Project Overview

This project demonstrates how to build and run Go applications. It includes several folders, each focusing on different areas of Go programming, from basic syntax to best practices for writing Go code.

For more information about Go, visit the [Go documentation](https://golang.org/doc/).

## Setup

### 1. Install Go

For detailed instructions on installing Go, refer to the [official Go installation guide](https://golang.org/doc/install).

### 2. Clone the Repository

```bash
git clone https://github.com/danhbuidcn/go_fundamental.git
cd go_fundamental
```

### 3. Build & Run the Application

Compile and run a specific Go file:
```bash
go build -o main main.go
./main
```
This command compiles the Go file, creates an executable named `main`, and runs it. The executable can be run directly without recompiling.

Alternatively, run the Go file directly without creating an executable:
```bash
go run main.go
```
This command compiles and runs the Go file immediately. The source code is recompiled each time you run it.

## Folder Structure

- [Tour of Go](./tour_of_go/README.md): Introduction to the Go programming language, with modules, exercises, and example programs that you can navigate through and run directly in your browser.
- [Basic](./basic/README.md): Provides examples and explanations of fundamental Go concepts such as syntax, variables, data types, and control structures.
- [How to Write Go Code](./how_to_write_go_code/README.md): A guide to writing and organizing Go code, including project structure and coding best practices.
- [Language specification](./language_specification/README.md): This manual details the Go programming language, focusing on its strong typing, concurrency, and package-based structure.
- [Effective Go](./effective_go/README.md): Provides guidelines and best practices for writing clear, idiomatic Go code, emphasizing that understanding Go's unique properties and conventions is crucial for producing high-quality Go programs.
- [Writing Web pplications](./writing_web_applications/README.md): A guide to writing and organizing Go code, including project structure and coding best practices.
- [Advance](./advance/README.md): Explores advanced Go topics like concurrency, goroutines, channels, performance optimization, and system interactions.
- [Projects](./projects/README.md): Offers practical examples and step-by-step guides for building complete projects with Go, applying learned concepts to real-world problems.

## How to Run Go Code

To run Go code within this repository:

1. Navigate to the folder containing the Go file you want to run.
2. Use `go run filename.go` to run the file directly.
3. Alternatively, use `go build -o outputname filename.go` to build an executable, then run the executable with `./outputname`.

## Learning Resources

Enhance your Go programming skills with these resources:

- <a href="https://roadmap.sh/golang" target="_blank">Go Roadmap</a>: A comprehensive guide to learning Go and its ecosystem.
- <a href="https://go.dev/wiki/" target="_blank">Go Wiki</a>: A collection of information about the Go Programming Language.
- <a href="https://www.cloudbees.com/blog/best-practices-for-a-new-go-developer" target="_blank">Best Practices for a New Go Developer</a>: Key practices for new Go developers to understand and adopt.
- <a href="https://go.dev/ref/spec" target="_blank">The Go Programming Language Specification</a>: This is the reference manual for the Go programming language
- <a href="https://200lab.io/blog/tag/golang/" target="_blank">Go blogs</a>: Collection of articles sharing Golang programming techniques
