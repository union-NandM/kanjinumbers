# kanjinumbers

## 概要

領収書や小切手に記載するような「壱万五千」のような表記を取り扱うサービスです。

実際の API はこちらです。

[GET /v1/kanji2number/{kanji}](https://rxyfko3ctb.execute-api.ap-northeast-1.amazonaws.com/v1/kanji2number/壱千弐百参拾四兆五千六百七拾八億九千壱拾弐万参千四百五拾六)  
[GET /v1/number2kanji/{num}](https://rxyfko3ctb.execute-api.ap-northeast-1.amazonaws.com/v1/number2kanji/1234567890123456)

## 使用技術

- Golang
- AWS Lambda
