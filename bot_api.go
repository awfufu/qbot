package qbot

import "github.com/awfufu/qbot/api"

// Bot Account APIs

func (b *Bot) SetQqProfile(nickname, company, email, college, personalNote string) error {
	return api.SetQqProfile(b, nickname, company, email, college, personalNote)
}

func (b *Bot) GetLoginInfo() (*api.LoginInfo, error) {
	return api.GetLoginInfo(b)
}

func (b *Bot) GetModelShow(model string) ([]api.ModelShow, error) {
	return api.GetModelShow(b, model)
}

func (b *Bot) SetModelShow(model, modelShow string) error {
	return api.SetModelShow(b, model, modelShow)
}

func (b *Bot) GetOnlineClients(noCache bool) ([]api.Device, error) {
	return api.GetOnlineClients(b, noCache)
}

// Friend APIs

func (b *Bot) GetStrangerInfo(userID uint64, noCache bool) (*api.StrangerInfo, error) {
	return api.GetStrangerInfo(b, userID, noCache)
}

func (b *Bot) GetFriendList() ([]api.FriendInfo, error) {
	return api.GetFriendList(b)
}

func (b *Bot) GetUnidirectionalFriendList() ([]api.UnidirectionalFriendInfo, error) {
	return api.GetUnidirectionalFriendList(b)
}

func (b *Bot) DeleteFriend(userID uint64) error {
	return api.DeleteFriend(b, userID)
}

func (b *Bot) DeleteUnidirectionalFriend(userID uint64) error {
	return api.DeleteUnidirectionalFriend(b, userID)
}

// Message APIs

func (b *Bot) SendPrivateMsg(userID uint64, message string) (uint64, error) {
	return api.SendPrivateMsg(b, userID, message, false)
}

func (b *Bot) SendPrivateTextMsg(userID uint64, message string) (uint64, error) {
	return api.SendPrivateMsg(b, userID, message, true)
}

func (b *Bot) SendGroupMsg(groupID uint64, message string) (uint64, error) {
	return api.SendGroupMsg(b, groupID, message, false)
}

func (b *Bot) SendGroupTextMsg(groupID uint64, message string) (uint64, error) {
	return api.SendGroupMsg(b, groupID, message, true)
}

func (b *Bot) SendMsg(groupID uint64, userID uint64, message string) (int32, error) {
	if groupID == 0 {
		return api.SendMsg(b, "private", userID, groupID, message, false)
	} else {
		return api.SendMsg(b, "group", userID, groupID, message, false)
	}
}

func (b *Bot) SendTextMsg(groupID uint64, userID uint64, message string) (int32, error) {
	if groupID == 0 {
		return api.SendMsg(b, "private", userID, groupID, message, true)
	} else {
		return api.SendMsg(b, "group", userID, groupID, message, true)
	}
}

func (b *Bot) GetMsg(messageID int32) (*api.MessageJson, error) {
	return api.GetMsg(b, messageID)
}

func (b *Bot) DeleteMsg(msgID uint64) error {
	return api.DeleteMsg(b, msgID)
}

func (b *Bot) MarkMsgAsRead(messageID int32) error {
	return api.MarkMsgAsRead(b, messageID)
}

func (b *Bot) GetForwardMsg(messageID string) ([]api.ForwardMsg, error) {
	return api.GetForwardMsg(b, messageID)
}

func (b *Bot) SendGroupForwardMsg(groupID uint64, messages []any) (int32, error) {
	return api.SendGroupForwardMsg(b, groupID, messages)
}

func (b *Bot) SendPrivateForwardMsg(userID uint64, messages []any) (int32, error) {
	return api.SendPrivateForwardMsg(b, userID, messages)
}

func (b *Bot) GetGroupMsgHistory(groupID uint64, messageSeq int32) ([]api.MessageJson, error) {
	return api.GetGroupMsgHistory(b, groupID, messageSeq)
}

// Image & Voice APIs

func (b *Bot) GetImage(file string) (*api.ImageInfo, error) {
	return api.GetImage(b, file)
}

func (b *Bot) CanSendImage() (bool, error) {
	return api.CanSendImage(b)
}

func (b *Bot) OcrImage(imageID string) (*api.OcrResult, error) {
	return api.OcrImage(b, imageID)
}

func (b *Bot) GetRecord(file, outFormat string) (string, error) {
	return api.GetRecord(b, file, outFormat)
}

func (b *Bot) CanSendRecord() (bool, error) {
	return api.CanSendRecord(b)
}

// Request APIs

func (b *Bot) SetFriendAddRequest(flag string, approve bool, remark string) error {
	return api.SetFriendAddRequest(b, flag, approve, remark)
}

func (b *Bot) SetGroupAddRequest(flag, subType string, approve bool, reason string) error {
	return api.SetGroupAddRequest(b, flag, subType, approve, reason)
}

// Group Info APIs

func (b *Bot) GetGroupInfo(groupID uint64, noCache bool) (*api.GroupInfo, error) {
	return api.GetGroupInfo(b, groupID, noCache)
}

func (b *Bot) GetGroupList(noCache bool) ([]api.GroupInfo, error) {
	return api.GetGroupList(b, noCache)
}

func (b *Bot) GetGroupMemberInfo(groupID uint64, userID uint64, noCache bool) (*api.GroupMemberInfo, error) {
	return api.GetGroupMemberInfo(b, groupID, userID, noCache)
}

func (b *Bot) GetGroupMemberList(groupID uint64, noCache bool) ([]api.GroupMemberInfo, error) {
	return api.GetGroupMemberList(b, groupID, noCache)
}

func (b *Bot) GetGroupHonorInfo(groupID uint64, typeStr string) (*api.GroupHonorInfo, error) {
	return api.GetGroupHonorInfo(b, groupID, typeStr)
}

func (b *Bot) GetGroupSystemMsg() (*api.GroupSystemMsg, error) {
	return api.GetGroupSystemMsg(b)
}

func (b *Bot) GetEssenceMsgList(groupID uint64) ([]api.EssenceMsg, error) {
	return api.GetEssenceMsgList(b, groupID)
}

func (b *Bot) GetGroupAtAllRemain(groupID uint64) (bool, int32, int32, error) {
	return api.GetGroupAtAllRemain(b, groupID)
}

// Group Setting APIs

func (b *Bot) SetGroupSpecialTitle(groupID uint64, userID uint64, specialTitle string) error {
	return api.SetGroupSpecialTitle(b, groupID, userID, specialTitle)
}

func (b *Bot) SetGroupName(groupID uint64, groupName string) error {
	return api.SetGroupName(b, groupID, groupName)
}

func (b *Bot) SetGroupAdmin(groupID uint64, userID uint64, enable bool) error {
	return api.SetGroupAdmin(b, groupID, userID, enable)
}

func (b *Bot) SetGroupBan(groupID uint64, userID uint64, duration int) error {
	return api.SetGroupBan(b, groupID, userID, duration)
}

func (b *Bot) SetGroupWholeBan(groupID uint64, enable bool) error {
	return api.SetGroupWholeBan(b, groupID, enable)
}

func (b *Bot) SetGroupAnonymousBan(groupID uint64, anonymous, anonymousFlag string, duration int) error {
	return api.SetGroupAnonymousBan(b, groupID, anonymous, anonymousFlag, duration)
}

func (b *Bot) SetGroupEssence(msgID uint64) error {
	return api.SetGroupEssence(b, msgID)
}

func (b *Bot) DeleteGroupEssence(msgID uint64) error {
	return api.DeleteGroupEssence(b, msgID)
}

func (b *Bot) SendGroupSign(groupID uint64) error {
	return api.SendGroupSign(b, groupID)
}

func (b *Bot) SetGroupAnonymous(groupID uint64, enable bool) error {
	return api.SetGroupAnonymous(b, groupID, enable)
}

func (b *Bot) SendGroupNotice(groupID uint64, content, image string) error {
	return api.SendGroupNotice(b, groupID, content, image)
}

func (b *Bot) GetGroupNotice(groupID uint64) ([]any, error) {
	return api.GetGroupNotice(b, groupID)
}

func (b *Bot) SetGroupKick(groupID uint64, userID uint64, rejectAddRequest bool) error {
	return api.SetGroupKick(b, groupID, userID, rejectAddRequest)
}

func (b *Bot) SetGroupLeave(groupID uint64, isDismiss bool) error {
	return api.SetGroupLeave(b, groupID, isDismiss)
}

func (b *Bot) SetGroupPortrait(groupID uint64, file string, cache int) error {
	return api.SetGroupPortrait(b, groupID, file, cache)
}

func (b *Bot) SetGroupCard(groupID uint64, userID uint64, card string) error {
	return api.SetGroupCard(b, groupID, userID, card)
}

// File APIs

func (b *Bot) UploadGroupFile(groupID uint64, file, name, folder string) error {
	return api.UploadGroupFile(b, groupID, file, name, folder)
}

func (b *Bot) DeleteGroupFile(groupID uint64, fileID string, busid int32) error {
	return api.DeleteGroupFile(b, groupID, fileID, busid)
}

func (b *Bot) CreateGroupFileFolder(groupID uint64, name, parentID string) error {
	return api.CreateGroupFileFolder(b, groupID, name, parentID)
}

func (b *Bot) DeleteGroupFileFolder(groupID uint64, folderID string) error {
	return api.DeleteGroupFileFolder(b, groupID, folderID)
}

func (b *Bot) GetGroupFileSystemInfo(groupID uint64) (*api.GroupFileSystemInfo, error) {
	return api.GetGroupFileSystemInfo(b, groupID)
}

func (b *Bot) GetGroupRootFiles(groupID uint64) (*struct {
	Files   []api.GroupFile   `json:"files"`
	Folders []api.GroupFolder `json:"folders"`
}, error) {
	return api.GetGroupRootFiles(b, groupID)
}

func (b *Bot) GetGroupFilesByFolder(groupID uint64, folderID string) (*struct {
	Files   []api.GroupFile   `json:"files"`
	Folders []api.GroupFolder `json:"folders"`
}, error) {
	return api.GetGroupFilesByFolder(b, groupID, folderID)
}

func (b *Bot) GetGroupFileUrl(groupID uint64, fileID string, busid int32) (string, error) {
	return api.GetGroupFileUrl(b, groupID, fileID, busid)
}

func (b *Bot) UploadPrivateFile(userID uint64, file, name string) error {
	return api.UploadPrivateFile(b, userID, file, name)
}

// System APIs

func (b *Bot) GetCookies(domain string) (string, error) {
	return api.GetCookies(b, domain)
}

func (b *Bot) GetCsrfToken() (int32, error) {
	return api.GetCsrfToken(b)
}

func (b *Bot) GetCredentials(domain string) (*api.Credentials, error) {
	return api.GetCredentials(b, domain)
}

func (b *Bot) GetVersionInfo() (*api.VersionInfo, error) {
	return api.GetVersionInfo(b)
}

func (b *Bot) GetStatus() (*api.StatusInfo, error) {
	return api.GetStatus(b)
}

func (b *Bot) ReloadEventFilter(file string) error {
	return api.ReloadEventFilter(b, file)
}

func (b *Bot) DownloadFile(url string, threadCount int, headers string) (string, error) {
	return api.DownloadFile(b, url, threadCount, headers)
}

func (b *Bot) CheckUrlSafely(url string) (int32, error) {
	return api.CheckUrlSafely(b, url)
}

func (b *Bot) CleanCache() error {
	return api.CleanCache(b)
}
