# Serverless Autopublisher

This is a task built on the [Serverless Framework](https://serverless.com/) to push commits to a repo at regular intervals, providing a universal solution to triggering deployments on a schedule.

[Read the announcement post on Forestry.io](https://forestry.io/blog/automatically-publish-scheduled-posts-for-static-site/)

## Instructions

This serverless function is written in Go, and the project includes binaries precompiled for the AWS Lambda environment. This means that you don't need to have Go installed on your local machine to configure and deploy this task.

Follow the [AWS quick start](https://serverless.com/framework/docs/providers/aws/guide/quick-start/) to install and configure Serverless.

You can clone this repo, or create it as a serverless template:

```
serverless create --template-url https://github.com/dwalkr/serverless-autopublish --path serverless-autopublish
```

The [serverless.yml](serverless.yml) file is where you will configure the autopublish script.

```
functions:
  publish:
    handler: bin/publish
    timeout: 15
    events:
      - schedule: rate(6 hours)
    environment:
      github_token: ${ssm:github_token}
      author_name: your-author-name
      author_email: your-author-email
      repos: https://github.com/FIRST-REPO;https://github.com/SECOND-REPO
```

This configuration defines a function called `publish` that will run every hour. You can adjust this by modifying the `schedule` parameter.

The variables in the `environment` section are passed to the function as environment variables. This is how we will configure the publisher. Change `author_name` and `author_email` to your desired Git signature, and add the repositories you want to publish to the `repos` variable. Separate multiple repos with a semicolon.

### Github Access Token

[Create an access token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) and add it to [AWS Systems Manager Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html) using the name `github_token`. To use a different parameter name, replace `github_token` with your parameter name in the `github_token` environment variable: `${ssm:your_parameter_name}`

### Deploy to AWS

Run `serverless deploy` to deploy the function to AWS. Your command should start running on the specified schedule. To invoke it immediately, run `serverless invoke -f publish` in your project directory.

### Troubleshooting
If something doesn't seem to be working, you can inspect the output of your function by running `serveless logs -f publish`

### Remove From AWS

Run the `serverless remove` command in your project directory to remove all objects from AWS.


## Development

In order to work on this code, you must have Go version 1.9 or higher as well as the [dep](https://golang.github.io/dep/) tool. Run `dep ensure` after cloning to install dependencies.

A `makefile` is included to compile the code for AWS Lambda. After modifying `publish/main.go`, run `make` to compile before deploying the code to AWS.

## Testing

Since serverless doesn't currently support local invocation of Lambda functions written in Go, you can use `publish_test.go` to test your code. Update the code with your testing data (*be sure not to commit your oauth token,*) and run it like this:

```
cd publish
go test
```