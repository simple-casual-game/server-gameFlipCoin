# server-gameFlipCoin


---

## 介紹
這是一個使用tcp連線的博奕遊戲伺服器，遊戲規則為翻硬幣，有50%機率贏或輸，贏的話獲得2倍賭注，輸的話就歸零

## packet格式
所有server通過tcp連線時，皆使用共同的packet格式。

```
	command  uint32
	sequence uint32
	size     uint32
    body     自訂
```

## 未來方向
1. 將通用的內容放至共用repo
2. 將protobuf另外放repo
3. 將後端與db串接
4. 將測試程式額外放repo
5. 與gate server連線, 由gate開放前端連線
6. 將error額外放repo


---
以上程式均由Josh開發完成