# kanjinumbers

## 概要

領収書や小切手に記載するような「壱万五千」のような表記を取り扱うサービスです。  
0以上1京未満の整数について、大字の漢数字表記とアラビア数字表記の相互変換に対応しています。  


デモページはこちらです。  
[kanjinumbers](http://kanjinumbers-page.s3-website-ap-northeast-1.amazonaws.com/)  


APIのURLはこちらです。  

GET /v1/number2kanji/{変換元のアラビア数字}  
https://rxyfko3ctb.execute-api.ap-northeast-1.amazonaws.com/v1/number2kanji/{number}  

GET /v1/kanji2number/{変換元の漢数字}  
https://rxyfko3ctb.execute-api.ap-northeast-1.amazonaws.com/v1/kanji2number/{kanji}  

## 使用技術

バックエンド
- Golang
- AWS Lambda（API実行環境）
- AWS S3（デモページのホスティング）  

フロントエンド
- React
- TypeScript
- tailwind CSS
