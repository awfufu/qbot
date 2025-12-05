package qbot

import "fmt"

// Deprecated: Use qbot.At instead
func CQAt(userId uint64) string {
	return fmt.Sprintf("[CQ:at,qq=%d]", userId)
}

// Deprecated: Use qbot.SendGroupReply or qbot.SendPrivateReply instead
func CQReply(msgId uint64) string {
	return fmt.Sprintf("[CQ:reply,id=%d]", msgId)
}

// Deprecated: Use qbot.Poke instead (if implemented) or raw segment
func CQPoke(userId uint64) string {
	return fmt.Sprintf("[CQ:poke,qq=%d]", userId)
}

// Deprecated: Use qbot.Image instead
func CQImage(file string) string {
	return fmt.Sprintf("[CQ:image,file=%s]", file)
}

// Deprecated: Use qbot.File instead (if implemented) or raw segment
func CQFile(file string) string {
	return fmt.Sprintf("[CQ:file,file=%s]", file)
}

// Deprecated: Use qbot.SendGroupRecord or qbot.SendPrivateRecord instead (if implemented) or raw segment.
func CQRecord(file string) string {
	return fmt.Sprintf("[CQ:record,file=%s]", file)
}

// Deprecated: Use qbot.SendGroupRps or qbot.SendPrivateRps instead (if implemented) or raw segment.
func CQRps() string {
	return "[CQ:rps]"
}

// Deprecated: Use qbot.SendGroupDice or qbot.SendPrivateDice instead (if implemented) or raw segment.
func CQDice() string {
	return "[CQ:dice]"
}
