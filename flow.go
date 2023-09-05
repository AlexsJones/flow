package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"

	golambda "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FlowStackProps struct {
	awscdk.StackProps
}

func NewFlowStack(scope constructs.Construct, id string, props *FlowStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Create SNS topic
	topic := awssns.NewTopic(stack, jsii.String("FlowTopic"), &awssns.TopicProps{
		DisplayName: jsii.String("FlowTopic"),
	})

	publisher := golambda.NewGoFunction(stack, jsii.String("PublisherLambda"), &golambda.GoFunctionProps{
		Entry: jsii.String("lambda/publisher/main.go"),
		Environment: &map[string]*string{
			"TOPIC_ARN": topic.TopicArn(),
			"REGION":    stack.Region(),
		},
	})
	// Allow the publisher to publish to the topic
	topic.GrantPublish(publisher)

	consumer := golambda.NewGoFunction(stack, jsii.String("ConsumerLambda"), &golambda.GoFunctionProps{
		Entry:       jsii.String("lambda/consumer/main.go"),
		Environment: &map[string]*string{},
	})
	// subscriber the consumer to the topic
	topic.AddSubscription(awssnssubscriptions.NewLambdaSubscription(consumer,
		&awssnssubscriptions.LambdaSubscriptionProps{}))

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewFlowStack(app, "FlowStack", &FlowStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
