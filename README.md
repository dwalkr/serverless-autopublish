# Serverless Autopublisher

This is a task built on the [Serverless Framework](https://serverless.com/) to push commits to a repo at regular intervals, providing a universal solution to triggering deployments on a schedule.

[Read the announcement post on Forestry.io](https://forestry.io/blog/automatically-publish-scheduled-posts-for-static-site/)

## Instructions

Install as a serverless template:

```bash
serverless create --template-url https://github.com/dwalkr/serverless-autopublish --path serverless-autopublish
```

[View the full usage documentation](./docs/usage.md)


## Development

Go 1.9 or higher is required to build the app.

[View the full development documentation](./docs/development.md)