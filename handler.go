package CQUPT_Sports_SDK

import (
	"github.com/Auto-CQUPT-Plan/CQUPT-Sports-SDK/models"
)

// GetToken 获取Token
func (r *SDK) GetToken(openid string) (*models.TokenInfo, error) {
	return r.client.GetToken(openid)
}

// GetUserInfo 获取用户信息
func (r *SDK) GetUserInfo(token string) (*models.UserInfo, error) {
	return r.client.GetUserInfo(token)
}

// GetNotice 获取通知
func (r *SDK) GetNotice(token string) (*models.NoticeInfo, error) {
	return r.client.GetNotice(token)
}

// GetAllArea 获取所有跑步场地
func (r *SDK) GetAllArea(token string) (*models.AreaInfo, error) {
	return r.client.GetAllArea(token)
}

// GetTerm 获取学期
func (r *SDK) GetTerm(token string) (*models.TermInfo, error) {
	return r.client.GetTerm(token)
}

// GetAllSportsResult 获取所有锻炼记录
func (r *SDK) GetAllSportsResult(token string, term string) (*models.AllSportsResults, error) {
	return r.client.GetAllSportsResult(token, term)
}

// StartRunning 开始跑步
func (r *SDK) StartRunning(token string, fieldID string) (*models.StartRunningResp, error) {
	return r.client.StartRunning(token, fieldID)
}

// GetCurrentRunningInfo 获取当前跑步详细信息
func (r *SDK) GetCurrentRunningInfo(token string, runningEventID string) (*models.CurrentRunningInfo, error) {
	return r.client.GetCurrentRunningInfo(token, runningEventID)
}

// UpdateRunningPoint 上报当前位置经纬度
func (r *SDK) UpdateRunningPoint(token string, runningEventID string, pointData []models.SportPointListNode) (*models.RunningLocationPointResp, error) {
	return r.client.UpdateRunningPoint(token, runningEventID, pointData)
}

// FinishRunning 结束跑步
func (r *SDK) FinishRunning(token string, runningEventID string) (*models.FinishRunningResp, error) {
	return r.client.FinishRunning(token, runningEventID)
}

func (r *SDK) Close() error {
	return r.client.ShutdownClient()
}

// GetFieldID_FengHua 风华运动场ID
func (r *SDK) GetFieldID_FengHua() string {
	return "T1001"
}

// GetFieldID_TaiJi 风华运动场ID
func (r *SDK) GetFieldID_TaiJi() string {
	return "T1005"
}

// GetFieldID_NingJing 风华运动场ID
func (r *SDK) GetFieldID_NingJing() string {
	return "T1014"
}
