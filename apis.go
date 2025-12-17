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

func (b *Sender) GetStrangerInfo(userID UserID, noCache bool) (*api.StrangerInfo, error) {
	return api.GetStrangerInfo(b, uint64(userID), noCache)
}

func (b *Sender) GetFriendList() ([]api.FriendInfo, error) {
	return api.GetFriendList(b)
}

func (b *Sender) GetUnidirectionalFriendList() ([]api.UnidirectionalFriendInfo, error) {
	return api.GetUnidirectionalFriendList(b)
}

func (b *Sender) DeleteFriend(userID UserID) error {
	return api.DeleteFriend(b, uint64(userID))
}

func (b *Sender) DeleteUnidirectionalFriend(userID UserID) error {
	return api.DeleteUnidirectionalFriend(b, uint64(userID))
}

// Private Message APIs

func (b *Sender) SendPrivateMsg(userID UserID, message ...any) (uint64, error) {
	return api.SendPrivateMsg(b, uint64(userID), toSegments(message...), false)
}

func (b *Sender) SendPrivateReplyMsg(userID UserID, msgID MsgID, message ...any) (uint64, error) {
	fullMessage := append([]any{replySegment(msgID)}, message...)
	return api.SendPrivateMsg(b, uint64(userID), toSegments(fullMessage...), false)
}

func (b *Sender) SendPrivateText(userID UserID, message string) (uint64, error) {
	return api.SendPrivateMsg(b, uint64(userID), []Segment{textSegment(message)}, true)
}

func (b *Sender) SendPrivateJson(userID UserID, data string) (uint64, error) {
	return api.SendPrivateMsg(b, uint64(userID), []Segment{jsonSegment(data)}, false)
}

func (b *Sender) SendPrivateVoice(userID UserID, file string) (uint64, error) {
	return api.SendPrivateMsg(b, uint64(userID), []Segment{recordSegment(file)}, false)
}

func (b *Sender) SendPrivateVideo(userID UserID, file string) (uint64, error) {
	return api.SendPrivateMsg(b, uint64(userID), []Segment{videoSegment(file)}, false)
}

func (b *Sender) SendPrivateMusic(userID UserID, typeStr, id string) (uint64, error) {
	return api.SendPrivateMsg(b, uint64(userID), []Segment{musicSegment(typeStr, id)}, false)
}

func (b *Sender) SendPrivateCustomMusic(userID UserID, url, audio, title, content, image string) (uint64, error) {
	return api.SendPrivateMsg(b, uint64(userID), []Segment{customMusicSegment(url, audio, title, content, image)}, false)
}

func (b *Sender) SendPrivateDice(userID UserID) (uint64, error) {
	return api.SendPrivateMsg(b, uint64(userID), []Segment{diceSegment()}, false)
}

func (b *Sender) SendPrivateRps(userID UserID) (uint64, error) {
	return api.SendPrivateMsg(b, uint64(userID), []Segment{rpsSegment()}, false)
}

func (b *Sender) SendPrivateFile(userID UserID, file string) (uint64, error) {
	return api.SendPrivateMsg(b, uint64(userID), []Segment{fileSegment(file)}, false)
}

func (b *Sender) SendPrivateForward(userID UserID, block ForwardBlock) (int32, string, error) {
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

	return api.SendPrivateForwardMsg(b, uint64(userID), messages, news, block.Prompt, block.Summary, block.Title)
}

func (b *Sender) SendPrivatePoke(userID UserID) error {
	return api.FriendPoke(b, uint64(userID))
}

func (b *Sender) ForwardMsgToPrivate(userID UserID, messageID string) (uint64, error) {
	return api.ForwardFriendSingleMsg(b, uint64(userID), messageID)
}

// Group Message APIs

func (b *Sender) SendGroupMsg(groupID GroupID, message ...any) (uint64, error) {
	return api.SendGroupMsg(b, uint64(groupID), toSegments(message...), false)
}

func (b *Sender) SendGroupReplyMsg(groupID GroupID, msgID MsgID, message ...any) (uint64, error) {
	fullMessage := append([]any{replySegment(msgID)}, message...)
	return api.SendGroupMsg(b, uint64(groupID), toSegments(fullMessage...), false)
}

func (b *Sender) SendGroupText(groupID GroupID, message string) (uint64, error) {
	return api.SendGroupMsg(b, uint64(groupID), []Segment{textSegment(message)}, true)
}

func (b *Sender) SendGroupJson(groupID GroupID, data string) (uint64, error) {
	return api.SendGroupMsg(b, uint64(groupID), []Segment{jsonSegment(data)}, false)
}

func (b *Sender) SendGroupVoice(groupID GroupID, file string) (uint64, error) {
	return api.SendGroupMsg(b, uint64(groupID), []Segment{recordSegment(file)}, false)
}

func (b *Sender) SendGroupVideo(groupID GroupID, file string) (uint64, error) {
	return api.SendGroupMsg(b, uint64(groupID), []Segment{videoSegment(file)}, false)
}

func (b *Sender) SendGroupMusic(groupID GroupID, typeStr, id string) (uint64, error) {
	return api.SendGroupMsg(b, uint64(groupID), []Segment{musicSegment(typeStr, id)}, false)
}

func (b *Sender) SendGroupCustomMusic(groupID GroupID, url, audio, title, content, image string) (uint64, error) {
	return api.SendGroupMsg(b, uint64(groupID), []Segment{customMusicSegment(url, audio, title, content, image)}, false)
}

func (b *Sender) SendGroupDice(groupID GroupID) (uint64, error) {
	return api.SendGroupMsg(b, uint64(groupID), []Segment{diceSegment()}, false)
}

func (b *Sender) SendGroupRps(groupID GroupID) (uint64, error) {
	return api.SendGroupMsg(b, uint64(groupID), []Segment{rpsSegment()}, false)
}

func (b *Sender) SendGroupFile(groupID GroupID, file string) (uint64, error) {
	return api.SendGroupMsg(b, uint64(groupID), []Segment{fileSegment(file)}, false)
}

func (b *Sender) SendGroupForward(groupID GroupID, block ForwardBlock) (int32, string, error) {
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

	return api.SendGroupForwardMsg(b, uint64(groupID), messages, news, block.Prompt, block.Summary, block.Title)
}

func (b *Sender) SendGroupPoke(groupID GroupID, userID UserID) error {
	return api.GroupPoke(b, uint64(groupID), uint64(userID))
}

func (b *Sender) ForwardMsgToGroup(messageID string, groupID GroupID) (uint64, error) {
	return api.ForwardGroupSingleMsg(b, uint64(groupID), messageID)
}

// Message APIs

func (b *Sender) GetMsg(messageID int32) (*api.MessageJson, error) {
	return api.GetMsg(b, messageID)
}

func (b *Sender) DeleteMsg(msgID MsgID) error {
	return api.DeleteMsg(b, uint64(msgID))
}

func (b *Sender) MarkMsgAsRead(messageID int32) error {
	return api.MarkMsgAsRead(b, messageID)
}

func (b *Sender) GetForwardMsg(messageID string) ([]api.ForwardMsg, error) {
	return api.GetForwardMsg(b, messageID)
}

func (b *Sender) GetGroupMsgHistory(groupID GroupID, messageSeq int32) ([]api.MessageJson, error) {
	return api.GetGroupMsgHistory(b, uint64(groupID), messageSeq)
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

func (b *Sender) GetGroupInfo(groupID GroupID, noCache bool) (*api.GroupInfo, error) {
	return api.GetGroupInfo(b, uint64(groupID), noCache)
}

func (b *Sender) GetGroupList(noCache bool) ([]api.GroupInfo, error) {
	return api.GetGroupList(b, noCache)
}

func (b *Sender) GetGroupMemberInfo(groupID GroupID, userID UserID, noCache bool) (*api.GroupMemberInfo, error) {
	return api.GetGroupMemberInfo(b, uint64(groupID), uint64(userID), noCache)
}

func (b *Sender) GetGroupMemberList(groupID GroupID, noCache bool) ([]api.GroupMemberInfo, error) {
	return api.GetGroupMemberList(b, uint64(groupID), noCache)
}

func (b *Sender) GetGroupHonorInfo(groupID GroupID, typeStr string) (*api.GroupHonorInfo, error) {
	return api.GetGroupHonorInfo(b, uint64(groupID), typeStr)
}

func (b *Sender) GetGroupSystemMsg() (*api.GroupSystemMsg, error) {
	return api.GetGroupSystemMsg(b)
}

func (b *Sender) GetEssenceMsgList(groupID GroupID) ([]api.EssenceMsg, error) {
	return api.GetEssenceMsgList(b, uint64(groupID))
}

func (b *Sender) GetGroupAtAllRemain(groupID GroupID) (bool, int32, int32, error) {
	return api.GetGroupAtAllRemain(b, uint64(groupID))
}

// Group Setting APIs

func (b *Sender) SetGroupSpecialTitle(groupID GroupID, userID UserID, specialTitle string) error {
	return api.SetGroupSpecialTitle(b, uint64(groupID), uint64(userID), specialTitle)
}

func (b *Sender) SetGroupName(groupID GroupID, groupName string) error {
	return api.SetGroupName(b, uint64(groupID), groupName)
}

func (b *Sender) SetGroupAdmin(groupID GroupID, userID UserID, enable bool) error {
	return api.SetGroupAdmin(b, uint64(groupID), uint64(userID), enable)
}

func (b *Sender) SetGroupBan(groupID GroupID, userID UserID, duration int) error {
	return api.SetGroupBan(b, uint64(groupID), uint64(userID), duration)
}

func (b *Sender) SetGroupWholeBan(groupID GroupID, enable bool) error {
	return api.SetGroupWholeBan(b, uint64(groupID), enable)
}

func (b *Sender) SetGroupAnonymousBan(groupID GroupID, anonymous, anonymousFlag string, duration int) error {
	return api.SetGroupAnonymousBan(b, uint64(groupID), anonymous, anonymousFlag, duration)
}

func (b *Sender) SetGroupEssence(msgID MsgID) error {
	return api.SetGroupEssence(b, uint64(msgID))
}

func (b *Sender) DeleteGroupEssence(msgID MsgID) error {
	return api.DeleteGroupEssence(b, uint64(msgID))
}

func (b *Sender) SendGroupSign(groupID GroupID) error {
	return api.SendGroupSign(b, uint64(groupID))
}

func (b *Sender) SetGroupAnonymous(groupID GroupID, enable bool) error {
	return api.SetGroupAnonymous(b, uint64(groupID), enable)
}

func (b *Sender) SendGroupNotice(groupID GroupID, content, image string) error {
	return api.SendGroupNotice(b, uint64(groupID), content, image)
}

func (b *Sender) GetGroupNotice(groupID GroupID) ([]any, error) {
	return api.GetGroupNotice(b, uint64(groupID))
}

func (b *Sender) SetGroupKick(groupID GroupID, userID UserID, rejectAddRequest bool) error {
	return api.SetGroupKick(b, uint64(groupID), uint64(userID), rejectAddRequest)
}

func (b *Sender) SetGroupLeave(groupID GroupID, isDismiss bool) error {
	return api.SetGroupLeave(b, uint64(groupID), isDismiss)
}

func (b *Sender) SetGroupPortrait(groupID GroupID, file string, cache int) error {
	return api.SetGroupPortrait(b, uint64(groupID), file, cache)
}

func (b *Sender) SetGroupCard(groupID GroupID, userID UserID, card string) error {
	return api.SetGroupCard(b, uint64(groupID), uint64(userID), card)
}

// File APIs

func (b *Sender) UploadGroupFile(groupID GroupID, file, name, folder string) error {
	return api.UploadGroupFile(b, uint64(groupID), file, name, folder)
}

func (b *Sender) DeleteGroupFile(groupID GroupID, fileID string, busid int32) error {
	return api.DeleteGroupFile(b, uint64(groupID), fileID, busid)
}

func (b *Sender) CreateGroupFileFolder(groupID GroupID, name, parentID string) error {
	return api.CreateGroupFileFolder(b, uint64(groupID), name, parentID)
}

func (b *Sender) DeleteGroupFileFolder(groupID GroupID, folderID string) error {
	return api.DeleteGroupFileFolder(b, uint64(groupID), folderID)
}

func (b *Sender) GetGroupFileSystemInfo(groupID GroupID) (*api.GroupFileSystemInfo, error) {
	return api.GetGroupFileSystemInfo(b, uint64(groupID))
}

func (b *Sender) GetGroupRootFiles(groupID GroupID) (*struct {
	Files   []api.GroupFile   `json:"files"`
	Folders []api.GroupFolder `json:"folders"`
}, error) {
	return api.GetGroupRootFiles(b, uint64(groupID))
}

func (b *Sender) GetGroupFilesByFolder(groupID GroupID, folderID string) (*struct {
	Files   []api.GroupFile   `json:"files"`
	Folders []api.GroupFolder `json:"folders"`
}, error) {
	return api.GetGroupFilesByFolder(b, uint64(groupID), folderID)
}

func (b *Sender) GetGroupFileUrl(groupID GroupID, fileID string, busid int32) (string, error) {
	return api.GetGroupFileUrl(b, uint64(groupID), fileID, busid)
}

func (b *Sender) UploadPrivateFile(userID UserID, file, name string) error {
	return api.UploadPrivateFile(b, uint64(userID), file, name)
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
