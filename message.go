package qbot

import "math"

type MsgType int

const (
	// normal types
	TextType  MsgType = 1
	AtType    MsgType = 2
	FaceType  MsgType = 3
	ImageType MsgType = 4

	// // private types
	// recordType  MsgType = 5
	// fileType    MsgType = 6
	// forwardType MsgType = 7
	// jsonType    MsgType = 8

	// // undefined
	// otherType MsgType = 0
)

type MsgItem interface {
	Type() MsgType

	GetTextItem() string
	GetAtItem() int64
	GetFaceItem() int16
	GetImageItem() *ImageItem
}

type TextItem string

func (i TextItem) Type() MsgType            { return TextType }
func (i TextItem) GetTextItem() string      { return string(i) }
func (i TextItem) GetAtItem() int64         { return -1 }
func (i TextItem) GetFaceItem() int16       { return -1 }
func (i TextItem) GetImageItem() *ImageItem { return nil }

func (i TextItem) String() string { return string(i) }

type AtItem int64

const AtAll AtItem = math.MaxInt64

func (i AtItem) Type() MsgType            { return AtType }
func (i AtItem) GetTextItem() string      { return "" }
func (i AtItem) GetAtItem() int64         { return int64(i) }
func (i AtItem) GetFaceItem() int16       { return -1 }
func (i AtItem) GetImageItem() *ImageItem { return nil }

type FaceItem int16

func (i FaceItem) Type() MsgType            { return FaceType }
func (i FaceItem) GetTextItem() string      { return "" }
func (i FaceItem) GetAtItem() int64         { return -1 }
func (i FaceItem) GetFaceItem() int16       { return int16(i) }
func (i FaceItem) GetImageItem() *ImageItem { return nil }

type ImageItem struct {
	Url string
}

func (i *ImageItem) Type() MsgType            { return ImageType }
func (i *ImageItem) GetTextItem() string      { return "" }
func (i *ImageItem) GetAtItem() int64         { return -1 }
func (i *ImageItem) GetFaceItem() int16       { return -1 }
func (i *ImageItem) GetImageItem() *ImageItem { return i }

// type recordItem struct {
// 	Path string
// }
// type fileItem struct {
// 	Data string
// }

// type forwardItem struct {
// 	Data string
// }
// type jsonItem struct {
// 	Data string
// }

// type otherItem struct {
// 	Data string
// }

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
