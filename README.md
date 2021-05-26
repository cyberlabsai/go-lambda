![Gopher](https://miro.medium.com/max/3200/1*G4kD68gYM1J1Fu-J7qvSSA.png)
# go-lambda
GoLambda is a lambda function "summoner". Could invoke two type: simple request or authorizer request.
## How it works
Simple request
```
package main

import lbd "github.com/cyberlabsai/go-lambda"

func main() {
    dataRequest := events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"gopher-name": "atila",
		},
		HTTPMethod: "GET",
	}
	userDataResponse, err := lbd.Invoke("gopher-function", "gopher-planet", dataRequest)
}
```
Authorizer
```
package main

import lbd "github.com/cyberlabsai/go-lambda"

func main() {
	authorizerRequest := events.APIGatewayCustomAuthorizerRequestTypeRequest{
		Type: "REQUEST",
		Headers: map[string]string{
			"Authorization": "Bearer GOPHER-TOKEN",
		},
	}
	jwtAuthorizerResponse, _ := lbd.InvokeAuthorizer("gopher-function", "gopher-planet", authorizerRequest)
}
```