const path = require("path");
import * as cdk from "@aws-cdk/core";
import * as lambda from "@aws-cdk/aws-lambda";
import * as iam from "@aws-cdk/aws-iam";

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

		// Permit this lambda to have all *lex V2* conversation policies
		//
		// https://docs.aws.amazon.com/lexv2/latest/dg/security_iam_id-based-policy-examples.html
		streemBotLambda.addToRolePolicy(
			new iam.PolicyStatement({
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
			})
		);
	}
}
