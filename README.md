# qbot

使用双向 HTTP 连接与 NapCat 通信。

## example

NapCat 配置：

- 设置 HTTP 服务端，监听 `http://napcat-ip:3000`
- 设置 HTTP 客户端，地址 `http://qbot-ip:3001`

*可指定任意端口，保证二者之间能通信即可。*

下面是一个 echo 示例。

```go
package main

import (
	"log"
	"github.com/awfufu/qbot"
)

func main() {
	bot := qbot.NewBot("qbot-ip:3001")
	bot.ConnectNapcat("http://napcat-ip:3000")
	bot.GroupMsg(func(msg *qbot.Message) {
		if msg.UserID == 114514 {
			bot.SendGroupMsg(msg.GroupID, msg.Raw)
		}
	})
	log.Fatal(bot.Run())
}
```

## run

```sh
go mod tidy
go run main.go
```