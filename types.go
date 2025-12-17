package qbot

import "math"

type GroupID uint64
type UserID uint64
type MsgID uint64
type FaceID uint16

const (
	InvalidGroup GroupID = 0
	InvalidUser  UserID  = 0
	InvalidMsgID MsgID   = 0
	InvalidFace  FaceID  = math.MaxUint16
)

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
	At() UserID
	Face() FaceID
	Image() *ImageItem
}

type TextItem string

func (i TextItem) Type() MsgType     { return TextType }
func (i TextItem) Text() string      { return string(i) }
func (i TextItem) At() UserID        { return InvalidUser }
func (i TextItem) Face() FaceID      { return InvalidFace }
func (i TextItem) Image() *ImageItem { return nil }

func (i TextItem) String() string { return string(i) }

type AtItem UserID

const AtAll UserID = math.MaxUint64

func (i AtItem) Type() MsgType     { return AtType }
func (i AtItem) Text() string      { return "" }
func (i AtItem) At() UserID        { return UserID(i) }
func (i AtItem) Face() FaceID      { return InvalidFace }
func (i AtItem) Image() *ImageItem { return nil }

type FaceItem FaceID

func (i FaceItem) Type() MsgType     { return FaceType }
func (i FaceItem) Text() string      { return "" }
func (i FaceItem) At() UserID        { return InvalidUser }
func (i FaceItem) Face() FaceID      { return FaceID(i) }
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
func (i *ImageItem) At() UserID        { return InvalidUser }
func (i *ImageItem) Face() FaceID      { return InvalidFace }
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
	MsgID    MsgID
	ReplyID  MsgID
	UserID   UserID
	Name     string
	Time     uint64

	// group
	GroupID   GroupID   // = InvalidGroup if msg from private
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
	GroupID   GroupID
	UserID    UserID
	MessageID MsgID
	IsAdd     bool
	IsQFace   bool
	Count     int32
	FaceID    FaceID
	EmojiRune rune
}

type RecallNotice struct {
	ChatType   ChatType
	GroupID    GroupID
	UserID     UserID
	OperatorID UserID
	MessageID  MsgID
	Time       int64
}

type PokeNotify struct {
	ChatType ChatType
	GroupID  GroupID
	SenderID UserID // Sender
	TargetID UserID // Receiver
	Action   string
	Suffix   string
}
