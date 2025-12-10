package api

import "encoding/json"

type MessageJson struct {
	GroupID     uint64 `json:"group_id"`
	Time        uint64 `json:"time"`
	MessageID   uint64 `json:"message_id"`
	MessageType string `json:"message_type"`
	Sender      struct {
		UserID   uint64 `json:"user_id"`
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
		Role     string `json:"role"`
	} `json:"sender"`
	RawMessage string `json:"raw_message"`
	Message    []struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	} `json:"message"`
}

type GroupMemberInfo struct {
	GroupID         uint64 `json:"group_id"`
	UserID          uint64 `json:"user_id"`
	Nickname        string `json:"nickname"`
	Card            string `json:"card"`
	Sex             string `json:"sex"`
	Age             int32  `json:"age"`
	Area            string `json:"area"`
	JoinTime        int32  `json:"join_time"`
	LastSentTime    int32  `json:"last_sent_time"`
	Level           string `json:"level"`
	Role            string `json:"role"`
	Unfriendly      bool   `json:"unfriendly"`
	Title           string `json:"title"`
	TitleExpireTime int64  `json:"title_expire_time"`
	CardChangeable  bool   `json:"card_changeable"`
	ShutUpTimestamp int64  `json:"shut_up_timestamp"`
}

type LoginInfo struct {
	UserID   uint64 `json:"user_id"`
	Nickname string `json:"nickname"`
}

type FriendInfo struct {
	UserID   uint64 `json:"user_id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}

type UnidirectionalFriendInfo struct {
	UserID   uint64 `json:"user_id"`
	Nickname string `json:"nickname"`
	Source   string `json:"source"`
}

type StrangerInfo struct {
	UserID    uint64 `json:"user_id"`
	Nickname  string `json:"nickname"`
	Sex       string `json:"sex"`
	Age       int32  `json:"age"`
	Qid       string `json:"qid"`
	Level     int32  `json:"level"`
	LoginDays int32  `json:"login_days"`
}

type GroupInfo struct {
	GroupID         uint64 `json:"group_id"`
	GroupName       string `json:"group_name"`
	GroupMemo       string `json:"group_memo"`
	GroupCreateTime uint32 `json:"group_create_time"`
	GroupLevel      uint32 `json:"group_level"`
	MemberCount     int32  `json:"member_count"`
	MaxMemberCount  int32  `json:"max_member_count"`
}

type GroupHonorInfo struct {
	GroupID          uint64 `json:"group_id"`
	CurrentTalkative struct {
		UserID   uint64 `json:"user_id"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
		DayCount int32  `json:"day_count"`
	} `json:"current_talkative"`
	TalkativeList []struct {
		UserID      uint64 `json:"user_id"`
		Nickname    string `json:"nickname"`
		Avatar      string `json:"avatar"`
		Description string `json:"description"`
	} `json:"talkative_list"`
	PerformerList []struct {
		UserID      uint64 `json:"user_id"`
		Nickname    string `json:"nickname"`
		Avatar      string `json:"avatar"`
		Description string `json:"description"`
	} `json:"performer_list"`
	LegendList []struct {
		UserID      uint64 `json:"user_id"`
		Nickname    string `json:"nickname"`
		Avatar      string `json:"avatar"`
		Description string `json:"description"`
	} `json:"legend_list"`
	StrongNewbieList []struct {
		UserID      uint64 `json:"user_id"`
		Nickname    string `json:"nickname"`
		Avatar      string `json:"avatar"`
		Description string `json:"description"`
	} `json:"strong_newbie_list"`
	EmotionList []struct {
		UserID      uint64 `json:"user_id"`
		Nickname    string `json:"nickname"`
		Avatar      string `json:"avatar"`
		Description string `json:"description"`
	} `json:"emotion_list"`
}

type VersionInfo struct {
	AppName         string `json:"app_name"`
	AppVersion      string `json:"app_version"`
	ProtocolVersion string `json:"protocol_version"`
}

type StatusInfo struct {
	AppInitialized bool `json:"app_initialized"`
	AppEnabled     bool `json:"app_enabled"`
	PluginsGood    bool `json:"plugins_good"`
	AppGood        bool `json:"app_good"`
	Online         bool `json:"online"`
	Good           bool `json:"good"`
	Stat           struct {
		PacketReceived  uint64 `json:"packet_received"`
		PacketSent      uint64 `json:"packet_sent"`
		PacketLost      uint64 `json:"packet_lost"`
		MessageReceived uint64 `json:"message_received"`
		MessageSent     uint64 `json:"message_sent"`
		DisconnectTimes uint32 `json:"disconnect_times"`
		LostTimes       uint32 `json:"lost_times"`
		LastMessageTime int64  `json:"last_message_time"`
	} `json:"stat"`
}

type FileInfo struct {
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	Url      string `json:"url"`
}

type GroupFileSystemInfo struct {
	FileCount  int32 `json:"file_count"`
	LimitCount int32 `json:"limit_count"`
	UsedSpace  int64 `json:"used_space"`
	TotalSpace int64 `json:"total_space"`
}

type GroupFile struct {
	GroupID       uint64 `json:"group_id"`
	FileID        string `json:"file_id"`
	FileName      string `json:"file_name"`
	BusID         int32  `json:"busid"`
	FileSize      int64  `json:"file_size"`
	UploadTime    int64  `json:"upload_time"`
	DeadTime      int64  `json:"dead_time"`
	ModifyTime    int64  `json:"modify_time"`
	DownloadTimes int32  `json:"download_times"`
	Uploader      uint64 `json:"uploader"`
	UploaderName  string `json:"uploader_name"`
}

type GroupFolder struct {
	GroupID        uint64 `json:"group_id"`
	FolderID       string `json:"folder_id"`
	FolderName     string `json:"folder_name"`
	CreateTime     int64  `json:"create_time"`
	Creator        uint64 `json:"creator"`
	CreatorName    string `json:"creator_name"`
	TotalFileCount int32  `json:"total_file_count"`
}

type EssenceMsg struct {
	SenderID     uint64 `json:"sender_id"`
	SenderNick   string `json:"sender_nick"`
	SenderTime   int64  `json:"sender_time"`
	OperatorID   uint64 `json:"operator_id"`
	OperatorNick string `json:"operator_nick"`
	OperatorTime int64  `json:"operator_time"`
	MessageID    int32  `json:"message_id"`
}

type ModelShow struct {
	Model     string `json:"model"`
	ModelShow string `json:"model_show"`
	NeedPay   bool   `json:"need_pay"`
}

type Device struct {
	AppID      int64  `json:"app_id"`
	DeviceName string `json:"device_name"`
	DeviceKind string `json:"device_kind"`
}

type OcrResult struct {
	Texts []struct {
		Text        string `json:"text"`
		Confidence  int32  `json:"confidence"`
		Coordinates any    `json:"coordinates"` // simplified
	} `json:"texts"`
	Language string `json:"language"`
}

type GroupSystemMsg struct {
	InvitedRequests []struct {
		RequestID   int64  `json:"request_id"`
		InvitorUin  uint64 `json:"invitor_uin"`
		InvitorNick string `json:"invitor_nick"`
		GroupID     uint64 `json:"group_id"`
		GroupName   string `json:"group_name"`
		Checked     bool   `json:"checked"`
		Actor       uint64 `json:"actor"`
	} `json:"invited_requests"`
	JoinRequests []struct {
		RequestID     int64  `json:"request_id"`
		RequesterUin  uint64 `json:"requester_uin"`
		RequesterNick string `json:"requester_nick"`
		Message       string `json:"message"`
		GroupID       uint64 `json:"group_id"`
		GroupName     string `json:"group_name"`
		Checked       bool   `json:"checked"`
		Actor         uint64 `json:"actor"`
	} `json:"join_requests"`
}

type Credentials struct {
	Cookies   string `json:"cookies"`
	CsrfToken int32  `json:"csrf_token"`
}

type ImageInfo struct {
	Size     int32  `json:"size"`
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

type ForwardMsg struct {
	Content string `json:"content"`
	Sender  struct {
		Nickname string `json:"nickname"`
		UserID   uint64 `json:"user_id"`
	} `json:"sender"`
	Time int64 `json:"time"`
}

type QiDianAccountInfo struct {
	MasterID   uint64 `json:"master_id"`
	MasterNick string `json:"master_nick"`
	Account    uint64 `json:"account"`
	Nickname   string `json:"nickname"`
}

type Segment struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

type News struct {
	Text string `json:"text"`
}

type EmojiLikeNotice struct {
	GroupID   uint64 `json:"group_id"`
	UserID    uint64 `json:"user_id"`
	MessageID uint64 `json:"message_id"`
	IsAdd     bool   `json:"is_add"`
	Likes     []struct {
		Count   int32  `json:"count"`
		EmojiID string `json:"emoji_id"`
	} `json:"likes"`
}

type GroupRecallNotice struct {
	GroupID    uint64 `json:"group_id"`
	UserID     uint64 `json:"user_id"`
	OperatorID uint64 `json:"operator_id"`
	MessageID  uint64 `json:"message_id"`
	Time       int64  `json:"time"`
}

type FriendRecallNotice struct {
	UserID    uint64 `json:"user_id"`
	MessageID uint64 `json:"message_id"`
	Time      int64  `json:"time"`
}

type PokeNotify struct {
	GroupID  uint64 `json:"group_id"`
	UserID   uint64 `json:"user_id"`   // Sender
	TargetID uint64 `json:"target_id"` // Receiver
	SubType  string `json:"sub_type"`
	RawInfo  any    `json:"raw_info"`
}
