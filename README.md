# presence

Presence is an AWS serverless application which utilizes IoT devices to communicate BLE presence to Slack.

```bash
.
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── presence                    <-- Source code for a lambda function
│   ├── main.go                 <-- Lambda function code
└── template.yaml
```

## Requirements

- AWS CLI already configured with Administrator permission
- [Docker installed](https://www.docker.com/community-edition)
- [Golang](https://golang.org)

## Setup process

### Installing dependencies

In this example we use the built-in `go get`:

```shell
go get -u github.com/aws/aws-lambda-go/...
go get -u github.com/nlopes/slack
```

### Building

Golang is a staticly compiled language, meaning that in order to run it you have to build the executeable target.

You can issue the following command in a shell to build it:

```shell
GOOS=linux GOARCH=amd64 go build -o presence/presence ./presence
```

**NOTE**: If you're not building the function on a Linux machine, you will need to specify the `GOOS` and `GOARCH` environment variables, this allows Golang to build your function for another system architecture and ensure compatability.

## Packaging and deployment

First and foremost, we need a `S3 bucket` where we can upload our Lambda functions packaged as ZIP before we deploy anything - If you don't have a S3 bucket to store code artifacts then this is a good time to create one:

```bash
aws s3 mb s3://presence-function
```

Next, run the following command to package our Lambda function to S3:

```bash
sam package \
    --template-file template.yaml \
    --output-template-file packaged.yaml \
    --s3-bucket presence-function
```

Next, the following command will create a Cloudformation Stack and deploy your SAM resources.

```bash
sam deploy \
    --template-file packaged.yaml \
    --stack-name presence \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides PresenceSlackToken=MySlackOAuthToken
```

> **See [Serverless Application Model (SAM) HOWTO Guide](https://github.com/awslabs/serverless-application-model/blob/master/HOWTO.md) for more details in how to get started.**

After deployment is complete you can run the following command to retrieve the API Gateway Endpoint URL:

```bash
aws cloudformation describe-stacks \
    --stack-name presence \
    --query 'Stacks[].Outputs'
```

### Testing (not implemented yet)

We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests:

```shell
go test -v ./presence/
```

# Appendix

## Golang installation

Please ensure Go 1.x (where 'x' is the latest version) is installed as per the instructions on the official golang website: <https://golang.org/doc/install>

A quickstart way would be to use Homebrew, chocolatey or your linux package manager.

### Homebrew (Mac)

Issue the following command from the terminal:

```shell
brew install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
brew update
brew upgrade golang
```

### Chocolatey (Windows)

Issue the following command from the powershell:

```shell
choco install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
choco upgrade golang
```

## AWS CLI commands

AWS CLI commands to package, deploy and describe outputs defined within the cloudformation stack:

```bash
sam package \
    --template-file template.yaml \
    --output-template-file packaged.yaml \
    --s3-bucket presence

sam deploy \
    --template-file packaged.yaml \
    --stack-name presence \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides PresenceSlackToken=MySlackOAuthToken

aws cloudformation describe-stacks \
    --stack-name presence --query 'Stacks[].Outputs'
```
