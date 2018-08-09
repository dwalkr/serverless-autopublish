---
title: Usage
---

# Usage

This serverless function is written in Go, and the project includes binaries precompiled for the AWS Lambda environment. This means that you don't need to have Go installed on your local machine to configure and deploy this task.

## Install Serverless
Follow the [AWS quick start](https://serverless.com/framework/docs/providers/aws/guide/quick-start/) to install and configure Serverless.


## Install the Autopublisher
You can clone this repo, or create it as a serverless template:

```bash
serverless create --template-url https://github.com/dwalkr/serverless-autopublish --path serverless-autopublish
```

The [serverless.yml](serverless.yml) file is where you will configure the autopublish script.

```yaml
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

## Github Access Token

[Create an access token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) and add it to [AWS Systems Manager Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html) using the name `github_token`. To use a different parameter name, replace `github_token` with your parameter name in the `github_token` environment variable: `${ssm:your_parameter_name}`

## Deploy to AWS

Run `serverless deploy` to deploy the function to AWS. Your command should start running on the specified schedule. To invoke it immediately, run `serverless invoke -f publish` in your project directory.

## Troubleshooting
If something doesn't seem to be working, you can inspect the output of your function by running `serverless logs -f publish`

## Remove From AWS

Run the `serverless remove` command in your project directory to remove all objects from AWS.