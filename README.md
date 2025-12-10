# qbot

使用双向 HTTP 连接与 NapCat 通信。

## example

NapCat 配置：

- 设置 HTTP 服务端，监听 `http://napcat-ip:3000`
- 设置 HTTP 客户端，地址 `http://qbot-ip:3001`

*（可指定任意端口，保证二者之间能通信即可。）*

下面是一个 echo 示例。

```go
package main

import (
  "log"

  "github.com/awfufu/qbot"
)

func main() {
  bot := qbot.NewBot("qbot-ip:3001") // 填写 NapCat 的 HTTP 客户端地址
  bot.ConnectNapcat("http://napcat-ip:3000") // 填写 NapCat 的 HTTP 服务端 URL
  bot.OnMessage(func(b *qbot.Bot, msg *qbot.Message) {
    if msg.Raw == "hello" {
      b.SendGroupMsg(msg.GroupID,
        qbot.At(msg.UserID), "world")
    }
  })
  log.Fatal(bot.Run())
}
```

```text
(you) > [hello]
(bot) < [world]
```

## run

```sh
go mod init yourproject
go get github.com/awfufu/qbot@v0.1.0
go mod tidy
# edit main.go
go run main.go
```