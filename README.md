# AWS Lex + AWS Lambda Example

![example](./images/login.gif)

This is an AWS Lex example using an AWS Lambda with Go, and the [AWS CDK](https://cdkworkshop.com/) to manage resources.

The end goal is to communicate with an AWS Lex chatbot using HTTP.

## Installation

#### CDK

- in `CDK` run `npm install`
- Ensure you have an aws profile configured with administrator access. [(Don't use your root account)](https://cdkworkshop.com/15-prerequisites/200-account.html).
- Unfortunately, [AWS Lex does not produce CloudFormation templates](https://github.com/aws/aws-cdk/issues/4905). You will need to use the AWS Console to create an AWS Lex bot

#### Lambda

- You can build your lambda by running `make build` in any lambda/LexLambdaFunction

Once you've completed the above steps, a `cdk bootstrap` and `cdk deploy` in the `CDK` directory will deploy your infrastructure

## Usage

WIP

## License

[MIT](https://choosealicense.com/licenses/mit/)
