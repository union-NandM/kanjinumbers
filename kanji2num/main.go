package main

import (
  "fmt"
  "net/url"
  "regexp"
  "strconv"
  "strings"

  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

var (
	figure = map[string]int{
		"壱": 1,
		"弐": 2,
		"参": 3,
		"四": 4,
		"五": 5,
		"六": 6,
		"七": 7,
		"八": 8,
		"九": 9,
	}
	
	digit_small = map[string]int{
		"拾": 10,
	  "百": 100,
		"千": 1000,
	}
	
	digit_large = map[string]int{
		"万": 10000,
		"億": 100000000,
		"兆": 1000000000000,
	}
)

/**
 *	与えられた文字列が適切な形式かを検査
 */
func checkFormat(str string) bool {

  // "零"の一文字からなる場合はtrue
  if str == "零" {
    return true
  }

  // 順番通りになっているかと不正な文字が使われていないかを検査
  re1 := regexp.MustCompile(`^(([壱弐参四五六七八九]千)?([壱弐参四五六七八九]百)?([壱弐参四五六七八九]拾)?[壱弐参四五六七八九]?兆)?(([壱弐参四五六七八九]千)?([壱弐参四五六七八九]百)?([壱弐参四五六七八九]拾)?[壱弐参四五六七八九]?億)?(([壱弐参四五六七八九]千)?([壱弐参四五六七八九]百)?([壱弐参四五六七八九]拾)?[壱弐参四五六七八九]?万)?(([壱弐参四五六七八九]千)?([壱弐参四五六七八九]百)?([壱弐参四五六七八九]拾)?[壱弐参四五六七八九]?)?$`)
  re1_match := re1.MatchString(str)
  if !re1_match {
    return false
  }

  // "兆万"などのような、`[兆億万]`のすぐ左隣に他の値がないパターンをはじく検査
  re2 := regexp.MustCompile(`[兆億万]`)
  separated_str := re2.Split(str, -1)
  for i, v := range separated_str {
    // 最後の要素は空文字列でもOK
    if i == len(separated_str) - 1 {
      break
    }
    if v == "" {
      return false
    }
  }

  return true
}

/**
 *	漢数字文字列を数値に変換する
 */
func converter(str string) int {
  // "零"のときは早期に返す
  if str == "零" {
    return 0
  }

  // 変換後の数値
  num := 0

  // 文字列を1文字ずつに分割
  slice := strings.Split(str, "")
  
  // 千単位での和
  temp_sum := 0
  // 直前の一桁の値を保持する変数
  temp := 0
  
  /*
   * 1文字ずつ見ていき、返す値に足していく
   */
  for i, v := range slice {
    switch v {
    case "壱","弐","参","四","五","六","七","八","九":
      temp = figure[v]
      
    case "拾","百","千":
      temp_sum += temp * digit_small[v]
      temp = 0
      
    case "万","億","兆":
      temp_sum += temp
      num += temp_sum * digit_large[v]
      temp = 0
      temp_sum = 0
    }
    
    if i == len(slice)-1 {
      temp_sum += temp
      num += temp_sum
    }
    
  }

  return num
}


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  // corsを有効化するためのヘッダー
  resHeaders := map[string]string{
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Credentials": "true",
    "Access-Control-Allow-Methods": "GET",
  }

  // pathパラメータの取得
  requestKanji := request.PathParameters["kanji"]

  // デコード
  kanji, err := url.QueryUnescape(requestKanji)

  if err != nil {
    return events.APIGatewayProxyResponse{
      StatusCode: 204,
      Headers: resHeaders,
    }, nil
  }

  // 形式がおかしいものをはじく
  if !checkFormat(kanji) {
    return events.APIGatewayProxyResponse{
      StatusCode: 204,
      Headers: resHeaders,
    }, nil
  }

  // 漢数字からアラビア数字に変換
  number := converter(kanji)
  // jsの MAX_SAFE_INTEGER の範囲に収まらないため文字列化する
  number_str := strconv.Itoa(number)
  

  // レスポンス
  return events.APIGatewayProxyResponse{
    Body:       fmt.Sprintf(`{"data":"%s","data_int":%d}`, number_str, number),
    StatusCode: 200,
    Headers: resHeaders,
  }, nil


}

func main() {
  lambda.Start(handler)
}
