package qbot

import (
	"github.com/awfufu/qbot/api"
)

// Bot Account APIs

func (b *Sender) SetProfile(nickname, company, email, college, personalNote string) error {
	return api.SetProfile(b, nickname, company, email, college, personalNote)
}

func (b *Sender) GetLoginInfo() (*api.LoginInfo, error) {
	return api.GetLoginInfo(b)
}

func (b *Sender) GetModelShow(model string) ([]api.ModelShow, error) {
	return api.GetModelShow(b, model)
}

func (b *Sender) SetModelShow(model, modelShow string) error {
	return api.SetModelShow(b, model, modelShow)
}

func (b *Sender) GetOnlineClients(noCache bool) ([]api.Device, error) {
	return api.GetOnlineClients(b, noCache)
}

// Friend APIs

func (b *Sender) GetStrangerInfo(userID uint64, noCache bool) (*api.StrangerInfo, error) {
	return api.GetStrangerInfo(b, userID, noCache)
}

func (b *Sender) GetFriendList() ([]api.FriendInfo, error) {
	return api.GetFriendList(b)
}

func (b *Sender) GetUnidirectionalFriendList() ([]api.UnidirectionalFriendInfo, error) {
	return api.GetUnidirectionalFriendList(b)
}

func (b *Sender) DeleteFriend(userID uint64) error {
	return api.DeleteFriend(b, userID)
}

func (b *Sender) DeleteUnidirectionalFriend(userID uint64) error {
	return api.DeleteUnidirectionalFriend(b, userID)
}

// Private Message APIs

func (b *Sender) SendPrivateMsg(userID uint64, message ...any) (uint64, error) {
	return api.SendPrivateMsg(b, userID, toSegments(message...), false)
}

func (b *Sender) SendPrivateReplyMsg(userID uint64, msgID uint64, message ...any) (uint64, error) {
	fullMessage := append([]any{replySegment(msgID)}, message...)
	return api.SendPrivateMsg(b, userID, toSegments(fullMessage...), false)
}

func (b *Sender) SendPrivateText(userID uint64, message string) (uint64, error) {
	return api.SendPrivateMsg(b, userID, []Segment{textSegment(message)}, true)
}

func (b *Sender) SendPrivateJson(userID uint64, data string) (uint64, error) {
	return api.SendPrivateMsg(b, userID, []Segment{jsonSegment(data)}, false)
}

func (b *Sender) SendPrivateVoice(userID uint64, file string) (uint64, error) {
	return api.SendPrivateMsg(b, userID, []Segment{recordSegment(file)}, false)
}

func (b *Sender) SendPrivateVideo(userID uint64, file string) (uint64, error) {
	return api.SendPrivateMsg(b, userID, []Segment{videoSegment(file)}, false)
}

func (b *Sender) SendPrivateMusic(userID uint64, typeStr, id string) (uint64, error) {
	return api.SendPrivateMsg(b, userID, []Segment{musicSegment(typeStr, id)}, false)
}

func (b *Sender) SendPrivateCustomMusic(userID uint64, url, audio, title, content, image string) (uint64, error) {
	return api.SendPrivateMsg(b, userID, []Segment{customMusicSegment(url, audio, title, content, image)}, false)
}

func (b *Sender) SendPrivateDice(userID uint64) (uint64, error) {
	return api.SendPrivateMsg(b, userID, []Segment{diceSegment()}, false)
}

func (b *Sender) SendPrivateRps(userID uint64) (uint64, error) {
	return api.SendPrivateMsg(b, userID, []Segment{rpsSegment()}, false)
}

func (b *Sender) SendPrivateFile(userID uint64, file string) (uint64, error) {
	return api.SendPrivateMsg(b, userID, []Segment{fileSegment(file)}, false)
}

func (b *Sender) SendPrivateForward(userID uint64, block ForwardBlock) (int32, string, error) {
	var messages []Segment
	for _, item := range block.Content {
		content := make([]any, len(item.Content))
		for i, s := range item.Content {
			content[i] = s
		}
		messages = append(messages, Segment(customNodeSegment(item.Name, item.UserID, content...)))
	}

	var news []api.News
	if block.Preview != "" {
		news = append(news, api.News{Text: block.Preview})
	}

	return api.SendPrivateForwardMsg(b, userID, messages, news, block.Prompt, block.Summary, block.Title)
}

func (b *Sender) SendPrivatePoke(userID uint64) error {
	return api.FriendPoke(b, userID)
}

func (b *Sender) ForwardMsgToPrivate(userID uint64, messageID string) (uint64, error) {
	return api.ForwardFriendSingleMsg(b, userID, messageID)
}

// Group Message APIs

func (b *Sender) SendGroupMsg(groupID uint64, message ...any) (uint64, error) {
	return api.SendGroupMsg(b, groupID, toSegments(message...), false)
}

func (b *Sender) SendGroupReplyMsg(groupID uint64, msgID uint64, message ...any) (uint64, error) {
	fullMessage := append([]any{replySegment(msgID)}, message...)
	return api.SendGroupMsg(b, groupID, toSegments(fullMessage...), false)
}

func (b *Sender) SendGroupText(groupID uint64, message string) (uint64, error) {
	return api.SendGroupMsg(b, groupID, []Segment{textSegment(message)}, true)
}

func (b *Sender) SendGroupJson(groupID uint64, data string) (uint64, error) {
	return api.SendGroupMsg(b, groupID, []Segment{jsonSegment(data)}, false)
}

func (b *Sender) SendGroupVoice(groupID uint64, file string) (uint64, error) {
	return api.SendGroupMsg(b, groupID, []Segment{recordSegment(file)}, false)
}

func (b *Sender) SendGroupVideo(groupID uint64, file string) (uint64, error) {
	return api.SendGroupMsg(b, groupID, []Segment{videoSegment(file)}, false)
}

func (b *Sender) SendGroupMusic(groupID uint64, typeStr, id string) (uint64, error) {
	return api.SendGroupMsg(b, groupID, []Segment{musicSegment(typeStr, id)}, false)
}

func (b *Sender) SendGroupCustomMusic(groupID uint64, url, audio, title, content, image string) (uint64, error) {
	return api.SendGroupMsg(b, groupID, []Segment{customMusicSegment(url, audio, title, content, image)}, false)
}

func (b *Sender) SendGroupDice(groupID uint64) (uint64, error) {
	return api.SendGroupMsg(b, groupID, []Segment{diceSegment()}, false)
}

func (b *Sender) SendGroupRps(groupID uint64) (uint64, error) {
	return api.SendGroupMsg(b, groupID, []Segment{rpsSegment()}, false)
}

func (b *Sender) SendGroupFile(groupID uint64, file string) (uint64, error) {
	return api.SendGroupMsg(b, groupID, []Segment{fileSegment(file)}, false)
}

func (b *Sender) SendGroupForward(groupID uint64, block ForwardBlock) (int32, string, error) {
	var messages []Segment
	for _, item := range block.Content {
		content := make([]any, len(item.Content))
		for i, s := range item.Content {
			content[i] = s
		}
		messages = append(messages, Segment(customNodeSegment(item.Name, item.UserID, content...)))
	}

	var news []api.News
	if block.Preview != "" {
		news = append(news, api.News{Text: block.Preview})
	}

	return api.SendGroupForwardMsg(b, groupID, messages, news, block.Prompt, block.Summary, block.Title)
}

func (b *Sender) SendGroupPoke(groupID uint64, userID uint64) error {
	return api.GroupPoke(b, groupID, userID)
}

func (b *Sender) ForwardMsgToGroup(messageID string, groupID uint64) (uint64, error) {
	return api.ForwardGroupSingleMsg(b, groupID, messageID)
}

// Message APIs

func (b *Sender) GetMsg(messageID int32) (*api.MessageJson, error) {
	return api.GetMsg(b, messageID)
}

func (b *Sender) DeleteMsg(msgID uint64) error {
	return api.DeleteMsg(b, msgID)
}

func (b *Sender) MarkMsgAsRead(messageID int32) error {
	return api.MarkMsgAsRead(b, messageID)
}

func (b *Sender) GetForwardMsg(messageID string) ([]api.ForwardMsg, error) {
	return api.GetForwardMsg(b, messageID)
}

func (b *Sender) GetGroupMsgHistory(groupID uint64, messageSeq int32) ([]api.MessageJson, error) {
	return api.GetGroupMsgHistory(b, groupID, messageSeq)
}

// Image & Voice APIs

func (b *Sender) GetImage(file string) (*api.ImageInfo, error) {
	return api.GetImage(b, file)
}

func (b *Sender) CanSendImage() (bool, error) {
	return api.CanSendImage(b)
}

func (b *Sender) OcrImage(imageID string) (*api.OcrResult, error) {
	return api.OcrImage(b, imageID)
}

func (b *Sender) GetRecord(file, outFormat string) (string, error) {
	return api.GetRecord(b, file, outFormat)
}

func (b *Sender) CanSendRecord() (bool, error) {
	return api.CanSendRecord(b)
}

func (b *Sender) SendEmojiReaction(messageID uint64, emojiID uint64, set bool) error {
	return api.SetMsgEmojiLike(b, messageID, emojiID, set)
}

// Request APIs

func (b *Sender) SetFriendAddRequest(flag string, approve bool, remark string) error {
	return api.SetFriendAddRequest(b, flag, approve, remark)
}

func (b *Sender) SetGroupAddRequest(flag, subType string, approve bool, reason string) error {
	return api.SetGroupAddRequest(b, flag, subType, approve, reason)
}

// Group Info APIs

func (b *Sender) GetGroupInfo(groupID uint64, noCache bool) (*api.GroupInfo, error) {
	return api.GetGroupInfo(b, groupID, noCache)
}

func (b *Sender) GetGroupList(noCache bool) ([]api.GroupInfo, error) {
	return api.GetGroupList(b, noCache)
}

func (b *Sender) GetGroupMemberInfo(groupID uint64, userID uint64, noCache bool) (*api.GroupMemberInfo, error) {
	return api.GetGroupMemberInfo(b, groupID, userID, noCache)
}

func (b *Sender) GetGroupMemberList(groupID uint64, noCache bool) ([]api.GroupMemberInfo, error) {
	return api.GetGroupMemberList(b, groupID, noCache)
}

func (b *Sender) GetGroupHonorInfo(groupID uint64, typeStr string) (*api.GroupHonorInfo, error) {
	return api.GetGroupHonorInfo(b, groupID, typeStr)
}

func (b *Sender) GetGroupSystemMsg() (*api.GroupSystemMsg, error) {
	return api.GetGroupSystemMsg(b)
}

func (b *Sender) GetEssenceMsgList(groupID uint64) ([]api.EssenceMsg, error) {
	return api.GetEssenceMsgList(b, groupID)
}

func (b *Sender) GetGroupAtAllRemain(groupID uint64) (bool, int32, int32, error) {
	return api.GetGroupAtAllRemain(b, groupID)
}

// Group Setting APIs

func (b *Sender) SetGroupSpecialTitle(groupID uint64, userID uint64, specialTitle string) error {
	return api.SetGroupSpecialTitle(b, groupID, userID, specialTitle)
}

func (b *Sender) SetGroupName(groupID uint64, groupName string) error {
	return api.SetGroupName(b, groupID, groupName)
}

func (b *Sender) SetGroupAdmin(groupID uint64, userID uint64, enable bool) error {
	return api.SetGroupAdmin(b, groupID, userID, enable)
}

func (b *Sender) SetGroupBan(groupID uint64, userID uint64, duration int) error {
	return api.SetGroupBan(b, groupID, userID, duration)
}

func (b *Sender) SetGroupWholeBan(groupID uint64, enable bool) error {
	return api.SetGroupWholeBan(b, groupID, enable)
}

func (b *Sender) SetGroupAnonymousBan(groupID uint64, anonymous, anonymousFlag string, duration int) error {
	return api.SetGroupAnonymousBan(b, groupID, anonymous, anonymousFlag, duration)
}

func (b *Sender) SetGroupEssence(msgID uint64) error {
	return api.SetGroupEssence(b, msgID)
}

func (b *Sender) DeleteGroupEssence(msgID uint64) error {
	return api.DeleteGroupEssence(b, msgID)
}

func (b *Sender) SendGroupSign(groupID uint64) error {
	return api.SendGroupSign(b, groupID)
}

func (b *Sender) SetGroupAnonymous(groupID uint64, enable bool) error {
	return api.SetGroupAnonymous(b, groupID, enable)
}

func (b *Sender) SendGroupNotice(groupID uint64, content, image string) error {
	return api.SendGroupNotice(b, groupID, content, image)
}

func (b *Sender) GetGroupNotice(groupID uint64) ([]any, error) {
	return api.GetGroupNotice(b, groupID)
}

func (b *Sender) SetGroupKick(groupID uint64, userID uint64, rejectAddRequest bool) error {
	return api.SetGroupKick(b, groupID, userID, rejectAddRequest)
}

func (b *Sender) SetGroupLeave(groupID uint64, isDismiss bool) error {
	return api.SetGroupLeave(b, groupID, isDismiss)
}

func (b *Sender) SetGroupPortrait(groupID uint64, file string, cache int) error {
	return api.SetGroupPortrait(b, groupID, file, cache)
}

func (b *Sender) SetGroupCard(groupID uint64, userID uint64, card string) error {
	return api.SetGroupCard(b, groupID, userID, card)
}

// File APIs

func (b *Sender) UploadGroupFile(groupID uint64, file, name, folder string) error {
	return api.UploadGroupFile(b, groupID, file, name, folder)
}

func (b *Sender) DeleteGroupFile(groupID uint64, fileID string, busid int32) error {
	return api.DeleteGroupFile(b, groupID, fileID, busid)
}

func (b *Sender) CreateGroupFileFolder(groupID uint64, name, parentID string) error {
	return api.CreateGroupFileFolder(b, groupID, name, parentID)
}

func (b *Sender) DeleteGroupFileFolder(groupID uint64, folderID string) error {
	return api.DeleteGroupFileFolder(b, groupID, folderID)
}

func (b *Sender) GetGroupFileSystemInfo(groupID uint64) (*api.GroupFileSystemInfo, error) {
	return api.GetGroupFileSystemInfo(b, groupID)
}

func (b *Sender) GetGroupRootFiles(groupID uint64) (*struct {
	Files   []api.GroupFile   `json:"files"`
	Folders []api.GroupFolder `json:"folders"`
}, error) {
	return api.GetGroupRootFiles(b, groupID)
}

func (b *Sender) GetGroupFilesByFolder(groupID uint64, folderID string) (*struct {
	Files   []api.GroupFile   `json:"files"`
	Folders []api.GroupFolder `json:"folders"`
}, error) {
	return api.GetGroupFilesByFolder(b, groupID, folderID)
}

func (b *Sender) GetGroupFileUrl(groupID uint64, fileID string, busid int32) (string, error) {
	return api.GetGroupFileUrl(b, groupID, fileID, busid)
}

func (b *Sender) UploadPrivateFile(userID uint64, file, name string) error {
	return api.UploadPrivateFile(b, userID, file, name)
}

// System APIs

func (b *Sender) GetCookies(domain string) (string, error) {
	return api.GetCookies(b, domain)
}

func (b *Sender) GetCsrfToken() (int32, error) {
	return api.GetCsrfToken(b)
}

func (b *Sender) GetCredentials(domain string) (*api.Credentials, error) {
	return api.GetCredentials(b, domain)
}

func (b *Sender) GetVersionInfo() (*api.VersionInfo, error) {
	return api.GetVersionInfo(b)
}

func (b *Sender) GetStatus() (*api.StatusInfo, error) {
	return api.GetStatus(b)
}

func (b *Sender) ReloadEventFilter(file string) error {
	return api.ReloadEventFilter(b, file)
}

func (b *Sender) DownloadFile(url string, threadCount int, headers string) (string, error) {
	return api.DownloadFile(b, url, threadCount, headers)
}

func (b *Sender) CheckUrlSafely(url string) (int32, error) {
	return api.CheckUrlSafely(b, url)
}

func (b *Sender) CleanCache() error {
	return api.CleanCache(b)
}
