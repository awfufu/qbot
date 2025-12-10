package qbot

type MsgType int

const (
	TextType    MsgType = 1
	AtType      MsgType = 2
	FaceType    MsgType = 3
	ImageType   MsgType = 4
	RecordType  MsgType = 5
	FileType    MsgType = 6
	ForwardType MsgType = 7
	JsonType    MsgType = 8

	OtherType MsgType = 0
)

type MsgItem interface {
	Type() MsgType

	GetTextItem() *TextItem
	GetAtItem() *AtItem
	GetFaceItem() *FaceItem
	GetImageItem() *ImageItem
	GetRecordItem() *RecordItem
	GetFileItem() *FileItem
	GetForwardItem() *ForwardItem
	GetJsonItem() *JsonItem
}

type TextItem struct {
	Content string
}

func (i *TextItem) Type() MsgType                { return TextType }
func (i *TextItem) GetTextItem() *TextItem       { return i }
func (i *TextItem) GetAtItem() *AtItem           { return nil }
func (i *TextItem) GetFaceItem() *FaceItem       { return nil }
func (i *TextItem) GetImageItem() *ImageItem     { return nil }
func (i *TextItem) GetRecordItem() *RecordItem   { return nil }
func (i *TextItem) GetFileItem() *FileItem       { return nil }
func (i *TextItem) GetForwardItem() *ForwardItem { return nil }
func (i *TextItem) GetJsonItem() *JsonItem       { return nil }

type AtItem struct {
	TargetID uint64
}

func (i *AtItem) Type() MsgType                { return AtType }
func (i *AtItem) GetTextItem() *TextItem       { return nil }
func (i *AtItem) GetAtItem() *AtItem           { return i }
func (i *AtItem) GetFaceItem() *FaceItem       { return nil }
func (i *AtItem) GetImageItem() *ImageItem     { return nil }
func (i *AtItem) GetRecordItem() *RecordItem   { return nil }
func (i *AtItem) GetFileItem() *FileItem       { return nil }
func (i *AtItem) GetForwardItem() *ForwardItem { return nil }
func (i *AtItem) GetJsonItem() *JsonItem       { return nil }

type FaceItem struct {
	ID uint64
}

func (i *FaceItem) Type() MsgType                { return FaceType }
func (i *FaceItem) GetTextItem() *TextItem       { return nil }
func (i *FaceItem) GetAtItem() *AtItem           { return nil }
func (i *FaceItem) GetFaceItem() *FaceItem       { return i }
func (i *FaceItem) GetImageItem() *ImageItem     { return nil }
func (i *FaceItem) GetRecordItem() *RecordItem   { return nil }
func (i *FaceItem) GetFileItem() *FileItem       { return nil }
func (i *FaceItem) GetForwardItem() *ForwardItem { return nil }
func (i *FaceItem) GetJsonItem() *JsonItem       { return nil }

type ImageItem struct {
	URL string
}

func (i *ImageItem) Type() MsgType                { return ImageType }
func (i *ImageItem) GetTextItem() *TextItem       { return nil }
func (i *ImageItem) GetAtItem() *AtItem           { return nil }
func (i *ImageItem) GetFaceItem() *FaceItem       { return nil }
func (i *ImageItem) GetImageItem() *ImageItem     { return i }
func (i *ImageItem) GetRecordItem() *RecordItem   { return nil }
func (i *ImageItem) GetFileItem() *FileItem       { return nil }
func (i *ImageItem) GetForwardItem() *ForwardItem { return nil }
func (i *ImageItem) GetJsonItem() *JsonItem       { return nil }

type RecordItem struct {
	Path string
}

func (i *RecordItem) Type() MsgType                { return RecordType }
func (i *RecordItem) GetTextItem() *TextItem       { return nil }
func (i *RecordItem) GetAtItem() *AtItem           { return nil }
func (i *RecordItem) GetFaceItem() *FaceItem       { return nil }
func (i *RecordItem) GetImageItem() *ImageItem     { return nil }
func (i *RecordItem) GetRecordItem() *RecordItem   { return i }
func (i *RecordItem) GetFileItem() *FileItem       { return nil }
func (i *RecordItem) GetForwardItem() *ForwardItem { return nil }
func (i *RecordItem) GetJsonItem() *JsonItem       { return nil }

type FileItem struct {
	Data string
}

func (i *FileItem) Type() MsgType                { return FileType }
func (i *FileItem) GetTextItem() *TextItem       { return nil }
func (i *FileItem) GetAtItem() *AtItem           { return nil }
func (i *FileItem) GetFaceItem() *FaceItem       { return nil }
func (i *FileItem) GetImageItem() *ImageItem     { return nil }
func (i *FileItem) GetRecordItem() *RecordItem   { return nil }
func (i *FileItem) GetFileItem() *FileItem       { return i }
func (i *FileItem) GetForwardItem() *ForwardItem { return nil }
func (i *FileItem) GetJsonItem() *JsonItem       { return nil }

type ForwardItem struct {
	Data string
}

func (i *ForwardItem) Type() MsgType                { return ForwardType }
func (i *ForwardItem) GetTextItem() *TextItem       { return nil }
func (i *ForwardItem) GetAtItem() *AtItem           { return nil }
func (i *ForwardItem) GetFaceItem() *FaceItem       { return nil }
func (i *ForwardItem) GetImageItem() *ImageItem     { return nil }
func (i *ForwardItem) GetRecordItem() *RecordItem   { return nil }
func (i *ForwardItem) GetFileItem() *FileItem       { return nil }
func (i *ForwardItem) GetForwardItem() *ForwardItem { return i }
func (i *ForwardItem) GetJsonItem() *JsonItem       { return nil }

type JsonItem struct {
	Data string
}

func (i *JsonItem) Type() MsgType                { return JsonType }
func (i *JsonItem) GetTextItem() *TextItem       { return nil }
func (i *JsonItem) GetAtItem() *AtItem           { return nil }
func (i *JsonItem) GetFaceItem() *FaceItem       { return nil }
func (i *JsonItem) GetImageItem() *ImageItem     { return nil }
func (i *JsonItem) GetRecordItem() *RecordItem   { return nil }
func (i *JsonItem) GetFileItem() *FileItem       { return nil }
func (i *JsonItem) GetForwardItem() *ForwardItem { return nil }
func (i *JsonItem) GetJsonItem() *JsonItem       { return i }

type OtherItem struct {
	Data string
}

func (i *OtherItem) Type() MsgType                { return OtherType }
func (i *OtherItem) GetTextItem() *TextItem       { return nil }
func (i *OtherItem) GetAtItem() *AtItem           { return nil }
func (i *OtherItem) GetFaceItem() *FaceItem       { return nil }
func (i *OtherItem) GetImageItem() *ImageItem     { return nil }
func (i *OtherItem) GetRecordItem() *RecordItem   { return nil }
func (i *OtherItem) GetFileItem() *FileItem       { return nil }
func (i *OtherItem) GetForwardItem() *ForwardItem { return nil }
func (i *OtherItem) GetJsonItem() *JsonItem       { return nil }

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
