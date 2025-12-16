# qbot

使用 Golang 实现的 NapCat QQ 机器人框架，通过双向 HTTP 连接与 NapCat 通信。

## example

NapCat 配置：

- NapCat HTTP 服务端，对应 `qbot.HttpClient`
- NapCat HTTP 客户端，对应 `qbot.HttpServer`

> 可指定任意端口，保证二者之间能通信即可。

下面是一个 echo 示例。

```go
package main

import (
  "log"
  "strings"

  "github.com/awfufu/qbot"
)

func main() {
  receiver := qbot.HttpServer(":3002")
  sender := qbot.HttpClient("http://napcat:3000")

  for {
    select {
    case msg := <-receiver.OnMessage():
      if msg.ChatType != qbot.Group {
        continue // Not a group message
      }

      if len(msg.Array) > 0 && msg.Array[0].Type() == qbot.TextType {
        if after, ok := strings.CutPrefix(msg.Array[0].GetTextItem(), "/echo "); ok {
          sender.SendGroupMsg(msg.GroupID, after, msg.Array[1:])
        }
      }

    case err := <-receiver.Error():
      // Handle HTTP server error
      log.Printf("http server error: %v", err)
      receiver.Close()
      return
    }
  }
}
```

```text
(you) > /echo helloworld
(bot) < helloworld
```

## run

```sh
go mod init yourproject
# edit main.go
go get github.com/awfufu/qbot
go mod tidy
go run main.go
```