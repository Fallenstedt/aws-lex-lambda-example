const path = require("path");
import * as cdk from "@aws-cdk/core";
import * as lambda from "@aws-cdk/aws-lambda";
import * as iam from "@aws-cdk/aws-iam";
import * as apigateway from "@aws-cdk/aws-apigateway";

export class LexStack extends cdk.Stack {
	constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
		super(scope, id, props);

		const streemBotLambda = new lambda.Function(this, "LexLambdaFunction", {
			runtime: lambda.Runtime.GO_1_X,
			handler: "main",
			code: lambda.Code.fromAsset(
				path.resolve("..", "lambda", "LexLambdaFunction", "main.zip")
			),
		});

		const streemBotFulfillmentLambda = new lambda.Function(
			this,
			"LexLambdaFulifllmentFunction",
			{
				runtime: lambda.Runtime.GO_1_X,
				handler: "main",
				code: lambda.Code.fromAsset(
					path.resolve("..", "lambda", "LexLambdaFulfillment", "main.zip")
				),
			}
		);

		// Permit this lambda to have all *lex V2* conversation policies
		//
		// https://docs.aws.amazon.com/lexv2/latest/dg/security_iam_id-based-policy-examples.html
		streemBotLambda.addToRolePolicy(this.lexPolicy());
		streemBotFulfillmentLambda.addToRolePolicy(this.lexPolicy());

		// API
		const api = new apigateway.RestApi(this, "LexApi");
		this.addCorsOptions(api.root);

		const integration = new apigateway.LambdaIntegration(streemBotLambda);
		api.root.addMethod("POST", integration, {
			operationName: "SendMessage",
		});
	}

	private lexPolicy() {
		return new iam.PolicyStatement({
			actions: [
				"lex:StartConversation",
				"lex:RecognizeText",
				"lex:RecognizeUtterance",
				"lex:GetSession",
				"lex:PutSession",
				"lex:DeleteSession",
			],
			effect: iam.Effect.ALLOW,
			resources: ["*"],
		});
	}

	private addCorsOptions(apiResource: apigateway.IResource) {
		apiResource.addMethod(
			"OPTIONS",
			new apigateway.MockIntegration({
				integrationResponses: [
					{
						statusCode: "200",
						responseParameters: {
							"method.response.header.Access-Control-Allow-Headers":
								"'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent'",
							"method.response.header.Access-Control-Allow-Origin": "'*'",
							"method.response.header.Access-Control-Allow-Credentials":
								"'false'",
							"method.response.header.Access-Control-Allow-Methods":
								"'OPTIONS,GET,PUT,POST,DELETE'",
						},
					},
				],
				passthroughBehavior: apigateway.PassthroughBehavior.NEVER,
				requestTemplates: {
					"application/json": '{"statusCode": 200}',
				},
			}),
			{
				methodResponses: [
					{
						statusCode: "200",
						responseParameters: {
							"method.response.header.Access-Control-Allow-Headers": true,
							"method.response.header.Access-Control-Allow-Methods": true,
							"method.response.header.Access-Control-Allow-Credentials": true,
							"method.response.header.Access-Control-Allow-Origin": true,
						},
					},
				],
			}
		);
	}
}
