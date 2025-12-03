package qbot

type MsgType int

const (
	Text    MsgType = 0
	At      MsgType = 1
	Face    MsgType = 2
	Image   MsgType = 3
	Record  MsgType = 4
	File    MsgType = 5
	Forward MsgType = 6
	Json    MsgType = 7

	Other MsgType = -1
)

type MsgItem interface {
	Type() MsgType
}

type TextItem struct {
	Content string
}

func (i *TextItem) Type() MsgType { return Text }

type AtItem struct {
	TargetID uint64
}

func (i *AtItem) Type() MsgType { return At }

type FaceItem struct {
	ID uint64
}

func (i *FaceItem) Type() MsgType { return Face }

type ImageItem struct {
	URL string
}

func (i *ImageItem) Type() MsgType { return Image }

type RecordItem struct {
	Path string
}

func (i *RecordItem) Type() MsgType { return Record }

type FileItem struct {
	Data string
}

func (i *FileItem) Type() MsgType { return File }

type ForwardItem struct {
	Data string
}

func (i *ForwardItem) Type() MsgType { return Forward }

type JsonItem struct {
	Data string
}

func (i *JsonItem) Type() MsgType { return Json }

type OtherItem struct {
	Data string
}

func (i *OtherItem) Type() MsgType { return Other }

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
