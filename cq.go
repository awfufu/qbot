package qbot

import "fmt"

func CQAt(userId uint64) string {
	return fmt.Sprintf("[CQ:at,qq=%d]", userId)
}

func CQReply(msgId uint64) string {
	return fmt.Sprintf("[CQ:reply,id=%d]", msgId)
}

func CQPoke(userId uint64) string {
	return fmt.Sprintf("[CQ:poke,qq=%d]", userId)
}

func CQImageFromUrl(url string) string {
	return fmt.Sprintf("[CQ:image,sub_type=0,url=%s]", url)
}

func CQImage(file string) string {
	return fmt.Sprintf("[CQ:image,file=file://data/%s]", file)
}

func CQFile(file string) string {
	return fmt.Sprintf("[CQ:file,file=file://data/%s]", file)
}

func CQRecord(file string) string {
	return fmt.Sprintf("[CQ:record,file=file://data/%s]", file)
}

func CQRps() string {
	return "[CQ:rps]"
}

func CQDice() string {
	return "[CQ:dice]"
}
