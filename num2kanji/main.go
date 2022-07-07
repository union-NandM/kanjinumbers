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

/*
 *  千ごと（4桁ごと）に文字列化する関数
 */
func miniConverter(num int) string {
	str := ""

	for i := 1000; i >= 1; i/=10 {
		if num / i > 0 {
			str += figure[int(num / i)] + digit_small[i]
		}
		num %= i;
	}

	return str
}

/*
 *  アラビア数字の数値を漢数字表記の文字列に変換する関数
 */
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

/*
 *  ハンドラ
 */
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// corsを有効化するためのヘッダー
	resHeaders := map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Access-Control-Allow-Credentials": "true",
		"Access-Control-Allow-Methods": "GET",
	}

	// pathパラメータの取得
	requestNum := request.PathParameters["num"]

	// 文字列を数値に変換
	num, err := strconv.Atoi(requestNum);

	// 変換できなかったら204で返す
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 204,
			Headers: resHeaders,
		}, nil
	}
	
	// 扱える数値の範囲外の値が来たら204で返す
	if num < 0 || num >= 10000000000000000 {
		return events.APIGatewayProxyResponse{
			StatusCode: 204,
			Headers: resHeaders,
		}, nil
	}

	// アラビア数字表記を漢数字表記に変換
	kanji := converter(num);

	// レスポンス
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf(`{"data":"%s"}`, kanji),
		StatusCode: 200,
		Headers: resHeaders,
	}, nil
}

func main() {
	lambda.Start(handler)
}
