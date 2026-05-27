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
	Duration        float64     `json:"duration"`
	Distance        float64     `json:"distance"`
	IsAppeal        string      `json:"isAppeal"`
	Reason          string      `json:"reason"`
}

type StartRunningPostData struct {
	PlaceName string `json:"placeName"`
	PlaceCode string `json:"placeCode"`
}

type StartRunningResp struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data string `json:"data"`
}

type CurrentRunningInfo struct {
	Msg  string             `json:"msg"`
	Code string             `json:"code"`
	Data CurrentRunningData `json:"data"`
}

type CurrentRunningData struct {
	SportRecordNo               string         `json:"sportRecordNo"`
	UnifyId                     string         `json:"unifyId"`
	StudentName                 string         `json:"studentName"`
	Mileage                     float64        `json:"mileage"`
	TimeConsuming               int64          `json:"timeConsuming"`
	ExpiredCountInForbiddenArea int64          `json:"expiredCountInForbiddenArea"`
	Points                      []PointsEntity `json:"points"`
	RunningCount                string         `json:"runningCount"`
}

type PointsEntity struct {
	Id            interface{} `json:"id"`
	SportPointNo  string      `json:"sportPointNo"`
	SportRecordNo string      `json:"sportRecordNo"`
	Longitude     string      `json:"longitude"`
	Latitude      string      `json:"latitude"`
	CollectTime   string      `json:"collectTime"`
	IsValid       string      `json:"isValid"`
	AfterPoints   interface{} `json:"afterPoints"`
	CreateBy      string      `json:"createBy"`
	CreateName    string      `json:"createName"`
	CreateTime    string      `json:"createTime"`
	UpdateBy      interface{} `json:"updateBy"`
	UpdateName    interface{} `json:"updateName"`
	UpdateTime    interface{} `json:"updateTime"`
	Remark        string      `json:"remark"`
}

type RunningLocationPointPostData struct {
	SportRecordNo  string               `json:"sportRecordNo"`
	SportPointList []SportPointListNode `json:"sportPointList"`
}

type SportPointListNode struct {
	SportRecordNo string  `json:"sportRecordNo"`
	Longitude     float64 `json:"longitude"`
	Latitude      float64 `json:"latitude"`
	PlaceName     string  `json:"placeName"`
	PlaceCode     string  `json:"placeCode"`
	CollectTime   string  `json:"collectTime"`
	IsValid       string  `json:"isValid"`
}

type RunningLocationPointResp struct {
	Msg  string                       `json:"msg"`
	Code string                       `json:"code"`
	Data RunningLocationPointRespData `json:"data"`
}

type RunningLocationPointRespData struct {
	SportRecordNo               string      `json:"sportRecordNo"`
	UnifyId                     string      `json:"unifyId"`
	StudentName                 string      `json:"studentName"`
	Mileage                     float64     `json:"mileage"`
	TimeConsuming               int64       `json:"timeConsuming"`
	ExpiredCountInForbiddenArea int64       `json:"expiredCountInForbiddenArea"`
	Points                      interface{} `json:"points"`
	RunningCount                string      `json:"runningCount"`
}

type FinishRunningResp struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data bool   `json:"data"`
}
