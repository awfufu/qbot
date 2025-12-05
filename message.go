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

type Message struct {
	GroupID  uint64
	UserID   uint64
	ReplyID  uint64
	Nickname string
	Card     string
	Role     string
	Time     uint64
	MsgID    uint64
	Raw      string
	Array    []MsgItem
}
