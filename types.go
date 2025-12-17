package qbot

import "math"

type MsgType uint8

const (
	TextType  MsgType = 1
	AtType    MsgType = 2
	FaceType  MsgType = 3
	ImageType MsgType = 4

	UnknownType MsgType = 0
)

type MsgItem interface {
	Type() MsgType
	Text() string
	At() int64
	Face() int16
	Image() *ImageItem
}

type TextItem string

func (i TextItem) Type() MsgType     { return TextType }
func (i TextItem) Text() string      { return string(i) }
func (i TextItem) At() int64         { return -1 }
func (i TextItem) Face() int16       { return -1 }
func (i TextItem) Image() *ImageItem { return nil }

func (i TextItem) String() string { return string(i) }

type AtItem int64

const AtAll AtItem = math.MaxInt64

func (i AtItem) Type() MsgType     { return AtType }
func (i AtItem) Text() string      { return "" }
func (i AtItem) At() int64         { return int64(i) }
func (i AtItem) Face() int16       { return -1 }
func (i AtItem) Image() *ImageItem { return nil }

type FaceItem int16

func (i FaceItem) Type() MsgType     { return FaceType }
func (i FaceItem) Text() string      { return "" }
func (i FaceItem) At() int64         { return -1 }
func (i FaceItem) Face() int16       { return int16(i) }
func (i FaceItem) Image() *ImageItem { return nil }

func (i FaceItem) String() string {
	if found, exists := qfaceMap[int(i)]; exists {
		return found
	}
	return ""
}

type ImageItem struct {
	Url string
}

func (i *ImageItem) Type() MsgType     { return ImageType }
func (i *ImageItem) Text() string      { return "" }
func (i *ImageItem) At() int64         { return -1 }
func (i *ImageItem) Face() int16       { return -1 }
func (i *ImageItem) Image() *ImageItem { return i }

type ChatType int8

const (
	Private ChatType = 1
	Group   ChatType = 2

	OtherChat ChatType = 0
)

type GroupRole int8

const (
	GroupMember GroupRole = 1
	GroupAdmin  GroupRole = 2
	GroupOwner  GroupRole = 3

	NotAGroup GroupRole = 0
)

type Message struct {
	ChatType ChatType // enum: Private, Group
	MsgID    uint64
	ReplyID  uint64
	UserID   uint64
	Name     string
	Time     uint64

	// group
	GroupID   uint64    // = 0  if msg from private
	GroupCard string    // = "" if msg from private
	GroupRole GroupRole // = NotAGroup if msg from private

	// content
	Raw   string
	Array []MsgItem
}

type EmojiLikeItem struct {
	Count   int32
	EmojiID uint64
}

type EmojiReaction struct {
	GroupID   uint64
	UserID    uint64
	MessageID uint64
	IsAdd     bool
	IsQFace   bool
	Count     int32
	FaceID    uint64
	EmojiRune rune
}

type RecallNotice struct {
	ChatType   ChatType
	GroupID    uint64
	UserID     uint64
	OperatorID uint64
	MessageID  uint64
	Time       int64
}

type PokeNotify struct {
	ChatType ChatType
	GroupID  uint64
	SenderID uint64 // Sender
	TargetID uint64 // Receiver
	Action   string
	Suffix   string
}
