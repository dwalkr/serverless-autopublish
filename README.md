## Instructions

1. Install Go and dep
2. Request a Github access token and add it to AWS Systems Manager Parameter Store
3. Update the environment vars in [serverless.yml](serverless.yml) with your `author_name`, `author_email`, and `repos` (semicolon-delimited) and adjust the cron rate if desired
4. `make` to build the binaries
5. `sls deploy`

## Testing

Since serverless doesn't currently support local invocation of Go lambdas, you can use `publish_test.go` to test your code. Update the code with your testing data (*be sure not to commit your oauth token,*) and run it like this:

```
cd publish
go test
```