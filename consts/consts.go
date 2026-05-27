package consts

const (
	GET_TOKEN_URL            = "https://sport.cqupt.edu.cn/new_wxapp/wxUnifyId/checkBinding"
	GET_USER_INFO_URL        = "https://sport.cqupt.edu.cn/new_wxapp/wxUnifyId/getUser"
	GET_NOTICE_URL           = "https://sport.cqupt.edu.cn/new_wxapp/notice/getNotice"
	GET_ALL_AREA_URL         = "https://sport.cqupt.edu.cn/new_wxapp/area/getAllArea"
	GET_TERM_URL             = "https://sport.cqupt.edu.cn/new_wxapp/yearTerm/list"
	GET_SPORTS_RESULT_URL    = "https://sport.cqupt.edu.cn/new_wxapp/sportsResult/list"
	START_RUNNING_URL        = "https://sport.cqupt.edu.cn/new_wxapp/sportRecord/sport/start2"
	GET_CURRENT_RUNNING_INFO = "https://sport.cqupt.edu.cn/new_wxapp/sportRecord/info"
	UPDATE_RUNNING_POINT     = "https://sport.cqupt.edu.cn/new_wxapp/sportRecord/point/saveListByNo"
	FINISH_RUNNING_URL       = "https://sport.cqupt.edu.cn/new_wxapp/sportRecord/sport/end"
)

const (
	PHONE_TYPE = "PJF110_Android 15"
	USER_AGENT = "Mozilla/5.0 (Linux; Android 15; PJF110 Build/UKQ1.231108.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/130.0.6723.103 Mobile Safari/537.36 XWEB/1300409 MMWEBSDK/20240404 MMWEBID/1532 MicroMessenger/8.0.49.2600(0x28003133) WeChat/arm64 Weixin NetType/5G Language/zh_CN ABI/arm64 MiniProgramEnv/android"
)

var SportsField = map[string]string{"T1001": "风华运动场", "T1005": "太极运动场", "T1014": "宁静苑"}
