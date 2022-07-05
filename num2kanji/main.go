package main

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	figure = map[int]string{
		1: "壱",
		2: "弐",
		3: "参",
		4: "四",
		5: "五",
		6: "六",
		7: "七",
		8: "八",
		9: "九",
	}
	
	digit_small = map[int]string{
		1:    "",
		10:   "拾",
		100:  "百",
		1000: "千",
	}
	
	digit_large = map[int]string{
		1:						 "",
		10000:         "万",
		100000000:     "億",
		1000000000000: "兆",
	}
)

func miniConverter(num int) string {
	str := ""

	var i int
	for i = 1000; i >= 1; i/=10 {
		if num / i > 0 {
			str += figure[int(num / i)] + digit_small[i]
		}
		num %= i;
	}

	return str
}

func converter(num int) string {
	if (num == 0) {
		return "零"
	}

	str := ""
	var i int
	for i = 1000000000000 ; i >= 1; i/=10000 {
		if num / i > 0 {
			str += miniConverter(num / i) + digit_large[i]
		}
		num %= i;
	}

	return str
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	resHeaders := map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Access-Control-Allow-Credentials": "true",
		"Access-Control-Allow-Methods": "GET",
	}

	requestNum := request.PathParameters["num"]

	num, err := strconv.Atoi(requestNum);

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 204,
			Headers: resHeaders,
		}, nil
	}
	
	if num < 0 || num > 10000000000000000 {
		return events.APIGatewayProxyResponse{
			StatusCode: 204,
			Headers: resHeaders,
		}, nil
	}

	kanji := converter(num);

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf(`{"data":"%s"}`, kanji),
		StatusCode: 200,
		Headers: resHeaders,
	}, nil
}

func main() {
	lambda.Start(handler)
}
