---
title: "Development"
---

# Development


In order to work on this code, you must have Go version 1.9 or higher as well as the [dep](https://golang.github.io/dep/) tool. Run `dep ensure` after cloning to install dependencies.

A `makefile` is included to compile the code for AWS Lambda. After modifying `publish/main.go`, run `make` to compile before deploying the code to AWS.

## Testing

Since serverless doesn't currently support local invocation of Lambda functions written in Go, you can use `publish_test.go` to test your code. Update the code with your testing data (*be sure not to commit your oauth token,*) and run it like this:

```bash
cd publish
go test
```