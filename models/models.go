package models

type TokenInfo struct {
	Msg  string        `json:"msg"`
	Code string        `json:"code"`
	Data TokenInfoData `json:"data"`
}

type TokenInfoData struct {
	IsBind    string `json:"isBind"`
	PublicKey string `json:"publicKey"`
	Token     string `json:"token"`
}

type UserInfo struct {
	Msg  string       `json:"msg"`
	Code string       `json:"code"`
	Data UserInfoData `json:"data"`
}

type UserInfoData struct {
	Id        int64       `json:"id"`
	Avatar    interface{} `json:"avatar"`
	Username  string      `json:"username"`
	UnifyId   string      `json:"unifyId"`
	StudentNo string      `json:"studentNo"`
	Sex       string      `json:"sex"`
	Grade     string      `json:"grade"`
	Major     interface{} `json:"major"`
	DeptName  string      `json:"deptName"`
	RoleCode  string      `json:"roleCode"`
	PublicKey string      `json:"publicKey"`
}

type NoticeInfo struct {
	Msg  string           `json:"msg"`
	Code string           `json:"code"`
	Data []NoticeInfoData `json:"data"`
}

type NoticeInfoData struct {
	Id         int64  `json:"id"`
	NoticeName string `json:"noticeName"`
	Content    string `json:"content"`
}

type AreaInfo struct {
	Msg  string         `json:"msg"`
	Code string         `json:"code"`
	Data []AreaInfoData `json:"data"`
}

type AreaInfoData struct {
	Id            int64           `json:"id"`
	AreaNo        string          `json:"areaNo"`
	PlaceName     string          `json:"placeName"`
	PlaceCode     string          `json:"placeCode"`
	AreaFuncName  string          `json:"areaFuncName"`
	AreaFuncCode  string          `json:"areaFuncCode"`
	CreateBy      string          `json:"createBy"`
	CreateName    string          `json:"createName"`
	CreateTime    string          `json:"createTime"`
	AreaPointList []AreaPointList `json:"areaPointList"`
}

type AreaPointList struct {
	Id        int64  `json:"id"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type TermInfo struct {
	Msg  string         `json:"msg"`
	Code string         `json:"code"`
	Data []TermInfoData `json:"data"`
}

type TermInfoData struct {
	YearTerm    string `json:"yearTerm"`
	YearTermDes string `json:"yearTermDes"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	IsCurrent   string `json:"isCurrent"`
}

type AllSportsResults struct {
	Msg  string                 `json:"msg"`
	Code string                 `json:"code"`
	Data []AllSportsResultsData `json:"data"`
}

type AllSportsResultsData struct {
	SportsResultNo  string      `json:"sportsResultNo"`
	SportsType      string      `json:"sportsType"`
	IsValid         string      `json:"isValid"`
	SportsStartTime string      `json:"sportsStartTime"`
	SportsEndTime   string      `json:"sportsEndTime"`
	PlaceName       string      `json:"placeName"`
	ReckonType      interface{} `json:"reckonType"`
	Duration        int64       `json:"duration"`
	Distance        int64       `json:"distance"`
	IsAppeal        string      `json:"isAppeal"`
	Reason          string      `json:"reason"`
}
