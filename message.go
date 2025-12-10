package qbot

type MsgType int

const (
	TextType    MsgType = 0
	AtType      MsgType = 1
	FaceType    MsgType = 2
	ImageType   MsgType = 3
	RecordType  MsgType = 4
	FileType    MsgType = 5
	ForwardType MsgType = 6
	JsonType    MsgType = 7

	OtherType MsgType = -1
)

type MsgItem interface {
	Type() MsgType
}

type TextItem struct {
	Content string
}

func (i *TextItem) Type() MsgType { return TextType }

type AtItem struct {
	TargetID uint64
}

func (i *AtItem) Type() MsgType { return AtType }

type FaceItem struct {
	ID uint64
}

func (i *FaceItem) Type() MsgType { return FaceType }

type ImageItem struct {
	URL string
}

func (i *ImageItem) Type() MsgType { return ImageType }

type RecordItem struct {
	Path string
}

func (i *RecordItem) Type() MsgType { return RecordType }

type FileItem struct {
	Data string
}

func (i *FileItem) Type() MsgType { return FileType }

type ForwardItem struct {
	Data string
}

func (i *ForwardItem) Type() MsgType { return ForwardType }

type JsonItem struct {
	Data string
}

func (i *JsonItem) Type() MsgType { return JsonType }

type OtherItem struct {
	Data string
}

func (i *OtherItem) Type() MsgType { return OtherType }

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
