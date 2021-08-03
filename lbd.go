package lbd

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type Client struct{}

var internatlClient = Client{}

func (Client) SimpleInvoke(functionName string, awsRegion string, payload []byte) (*lambda.InvokeOutput, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := lambda.New(sess, &aws.Config{Region: aws.String(awsRegion)})

	result, err := client.Invoke(&lambda.InvokeInput{
		FunctionName:   aws.String(functionName),
		InvocationType: aws.String("RequestResponse"),
		Payload:        payload,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Client) InvokeAuthorizer(functionName string, awsRegion string, request events.APIGatewayCustomAuthorizerRequestTypeRequest) (*events.APIGatewayCustomAuthorizerResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	result, err := internatlClient.SimpleInvoke(functionName, awsRegion, payload)
	if err != nil {
		return nil, err
	}

	var resp events.APIGatewayCustomAuthorizerResponse
	err = json.Unmarshal(result.Payload, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Context == nil {
		err = errors.New("Unauthorized")
		return nil, err
	}

	return &resp, nil
}

func (Client) Invoke(functionName string, awsRegion string, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	result, err := internatlClient.SimpleInvoke(functionName, awsRegion, payload)
	if err != nil {
		return nil, err
	}

	var resp events.APIGatewayProxyResponse

	err = json.Unmarshal(result.Payload, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
