# CQUPT-Sports-SDK

重庆邮电大学体育系统 Go SDK，封装了智慧体育平台 (`sport.cqupt.edu.cn`) 的 API 接口。

## 项目结构

```
CQUPT-Sports-SDK/
├── client/
│   ├── client.go       # HTTP 客户端初始化、Cookie 持久化、关闭
│   └── root.go         # 底层 HTTP 请求实现（令牌获取、用户信息、跑步记录等）
├── consts/
│   └── consts.go       # API 地址常量、User-Agent、运动场地映射
├── models/
│   └── models.go       # 请求/响应数据结构体定义
├── handler.go          # SDK 公开 API 层（封装 client 方法）
├── root.go             # SDK 入口：Options 模式初始化
├── unity_test.go       # 单元测试
├── test_cookies.json   # 测试用 Cookie 持久化文件
├── test_location.json  # 测试用跑步轨迹坐标数据
├── go.mod / go.sum
└── LICENSE             # MIT License
```

## 安装

```bash
go get github.com/Auto-CQUPT-Plan/CQUPT-Sports-SDK
```

## 初始化

```go
import sdk "github.com/Auto-CQUPT-Plan/CQUPT-Sports-SDK"

// 使用默认 Cookie 路径（./CookieJar.json）
s, err := sdk.NewSDK()

// 自定义 Cookie 持久化路径
s, err := sdk.NewSDK(sdk.WithCookiePath("./my_cookies.json"))

// 使用完毕后关闭（保存 Cookie）
defer s.Close()
```

### `NewSDK(opts ...Options) (*SDK, error)`

| 参数 | 类型 | 说明 |
|------|------|------|
| opts | `...Options` | 可选配置项，目前支持 `WithCookiePath` |

| 返回值 | 类型 | 说明 |
|--------|------|------|
| sdk | `*SDK` | SDK 实例 |
| err | `error` | 初始化错误 |

### `WithCookiePath(path string) Options`

设置 Cookie Jar 持久化文件路径。

---

## SDK 函数

### 1. GetToken — 获取 Token

```go
tokenInfo, err := s.GetToken("your_openid")
```

| 参数 | 类型 | 说明 |
|------|------|------|
| openid | `string` | 微信 OpenID |

| 返回值 | 类型 | 说明 |
|--------|------|------|
| tokenInfo | `*models.TokenInfo` | 包含 isBind、publicKey、token |
| err | `error` | 请求错误 |

---

### 2. GetUserInfo — 获取用户信息

```go
userInfo, err := s.GetUserInfo("your_token")
```

| 参数 | 类型 | 说明 |
|------|------|------|
| token | `string` | 用户 Token |

| 返回值 | 类型 | 说明 |
|--------|------|------|
| userInfo | `*models.UserInfo` | 包含学号、姓名、性别、年级、院系等 |
| err | `error` | 请求错误 |

---

### 3. GetNotice — 获取通知

```go
noticeInfo, err := s.GetNotice("your_token")
```

| 参数 | 类型 | 说明 |
|------|------|------|
| token | `string` | 用户 Token |

| 返回值 | 类型 | 说明 |
|--------|------|------|
| noticeInfo | `*models.NoticeInfo` | 通知列表（标题、内容） |
| err | `error` | 请求错误 |

---

### 4. GetAllArea — 获取所有跑步场地

```go
areaInfo, err := s.GetAllArea("your_token")
```

| 参数 | 类型 | 说明 |
|------|------|------|
| token | `string` | 用户 Token |

| 返回值 | 类型 | 说明 |
|--------|------|------|
| areaInfo | `*models.AreaInfo` | 场地列表（含区域编号、名称、坐标点） |
| err | `error` | 请求错误 |

---

### 5. GetTerm — 获取学期信息

```go
termInfo, err := s.GetTerm("your_token")
```

| 参数 | 类型 | 说明 |
|------|------|------|
| token | `string` | 用户 Token |

| 返回值 | 类型 | 说明 |
|--------|------|------|
| termInfo | `*models.TermInfo` | 学期列表（如 20251、20252，含起止日期、是否当前学期） |
| err | `error` | 请求错误 |

---

### 6. GetAllSportsResult — 获取所有锻炼记录

```go
results, err := s.GetAllSportsResult("your_token", "20252")
```

| 参数 | 类型 | 说明 |
|------|------|------|
| token | `string` | 用户 Token |
| term | `string` | 学期编号，如 `"20252"`（2025-2026 第二学期） |

| 返回值 | 类型 | 说明 |
|--------|------|------|
| results | `*models.AllSportsResults` | 锻炼记录列表（运动类型、是否有效、时长、距离等） |
| err | `error` | 请求错误 |

---

### 7. StartRunning — 开始跑步

```go
resp, err := s.StartRunning("your_token", s.GetFieldID_FengHua())
```

| 参数 | 类型 | 说明 |
|------|------|------|
| token | `string` | 用户 Token |
| fieldID | `string` | 场地编号，如 `"T1001"` |

| 返回值 | 类型 | 说明 |
|--------|------|------|
| resp | `*models.StartRunningResp` | 开始跑步响应（data 为跑步记录编号） |
| err | `error` | 请求错误 |

---

### 8. GetCurrentRunningInfo — 获取当前跑步详细信息

```go
info, err := s.GetCurrentRunningInfo("your_token", "running_event_id")
```

| 参数 | 类型 | 说明 |
|------|------|------|
| token | `string` | 用户 Token |
| runningEventID | `string` | 跑步事件编号（由 StartRunning 返回） |

| 返回值 | 类型 | 说明 |
|--------|------|------|
| info | `*models.CurrentRunningInfo` | 当前跑步详情（里程、耗时、轨迹点等） |
| err | `error` | 请求错误 |

---

### 9. UpdateRunningPoint — 上报跑步位置点

```go
points := []models.SportPointListNode{
    {
        SportRecordNo: "",
        Longitude:     106.60525,
        Latitude:      29.53286,
        PlaceName:     "风华运动场",
        PlaceCode:     "T1001",
        CollectTime:   "2026-05-27 15:30:00",
        IsValid:       "1",
    },
}
resp, err := s.UpdateRunningPoint("your_token", "running_event_id", points)
```

| 参数 | 类型 | 说明 |
|------|------|------|
| token | `string` | 用户 Token |
| runningEventID | `string` | 跑步事件编号 |
| pointData | `[]models.SportPointListNode` | 位置点数组（经纬度、场地、采集时间等） |

| 返回值 | 类型 | 说明 |
|--------|------|------|
| resp | `*models.RunningLocationPointResp` | 上报响应（含当前里程、耗时） |
| err | `error` | 请求错误 |

---

### 10. FinishRunning — 结束跑步

```go
resp, err := s.FinishRunning("your_token", "running_event_id")
```

| 参数 | 类型 | 说明 |
|------|------|------|
| token | `string` | 用户 Token |
| runningEventID | `string` | 跑步事件编号 |

| 返回值 | 类型 | 说明 |
|--------|------|------|
| resp | `*models.FinishRunningResp` | 结束跑步响应（data 为 bool 表示是否成功） |
| err | `error` | 请求错误 |

---

### 11. Close — 关闭 SDK（持久化 Cookie）

```go
err := s.Close()
```

| 返回值 | 说明 |
|--------|------|
| `error` | 保存 Cookie 到文件时的错误 |

---

### 12. 场地编号辅助方法

```go
s.GetFieldID_FengHua()   // "T1001" — 风华运动场
s.GetFieldID_TaiJi()     // "T1005" — 太极运动场
s.GetFieldID_NingJing()  // "T1014" — 宁静苑
```

---

## 数据结构

所有返回值的 `Msg` 和 `Code` 为通用响应字段，以下仅列出各结构体特有的 `Data` 字段。

### TokenInfo — GetToken 返回值

```go
type TokenInfo struct {
    Msg  string        `json:"msg"`
    Code string        `json:"code"`
    Data TokenInfoData `json:"data"`
}
type TokenInfoData struct {
    IsBind    string `json:"isBind"`    // 是否已绑定
    PublicKey string `json:"publicKey"` // 公钥
    Token     string `json:"token"`     // 用户 Token
}
```

### UserInfo — GetUserInfo 返回值

```go
type UserInfo struct {
    Msg  string       `json:"msg"`
    Code string       `json:"code"`
    Data UserInfoData `json:"data"`
}
type UserInfoData struct {
    Id        int64       `json:"id"`
    Avatar    interface{} `json:"avatar"`
    Username  string      `json:"username"`
    UnifyId   string      `json:"unifyId"`   // 统一 ID
    StudentNo string      `json:"studentNo"` // 学号
    Sex       string      `json:"sex"`       // 性别
    Grade     string      `json:"grade"`     // 年级
    Major     interface{} `json:"major"`     // 专业
    DeptName  string      `json:"deptName"`  // 院系
    RoleCode  string      `json:"roleCode"`  // 角色
    PublicKey string      `json:"publicKey"`
}
```

### NoticeInfo — GetNotice 返回值

```go
type NoticeInfo struct {
    Msg  string           `json:"msg"`
    Code string           `json:"code"`
    Data []NoticeInfoData `json:"data"`
}
type NoticeInfoData struct {
    Id         int64  `json:"id"`
    NoticeName string `json:"noticeName"` // 通知标题
    Content    string `json:"content"`    // 通知内容
}
```

### AreaInfo — GetAllArea 返回值

```go
type AreaInfo struct {
    Msg  string         `json:"msg"`
    Code string         `json:"code"`
    Data []AreaInfoData `json:"data"`
}
type AreaInfoData struct {
    Id            int64           `json:"id"`
    AreaNo        string          `json:"areaNo"`        // 区域编号
    PlaceName     string          `json:"placeName"`     // 场地名称
    PlaceCode     string          `json:"placeCode"`     // 场地代码
    AreaFuncName  string          `json:"areaFuncName"`  // 区域功能名
    AreaFuncCode  string          `json:"areaFuncCode"`  // 区域功能代码
    CreateBy      string          `json:"createBy"`
    CreateName    string          `json:"createName"`
    CreateTime    string          `json:"createTime"`
    AreaPointList []AreaPointList `json:"areaPointList"` // 区域坐标点列表
}
type AreaPointList struct {
    Id        int64  `json:"id"`
    Longitude string `json:"longitude"` // 经度
    Latitude  string `json:"latitude"`  // 纬度
}
```

### TermInfo — GetTerm 返回值

```go
type TermInfo struct {
    Msg  string         `json:"msg"`
    Code string         `json:"code"`
    Data []TermInfoData `json:"data"`
}
type TermInfoData struct {
    YearTerm    string `json:"yearTerm"`    // 学期编号（如 "20251"）
    YearTermDes string `json:"yearTermDes"` // 学期描述（如 "2025-2026第一学期"）
    StartDate   string `json:"startDate"`   // 开始日期
    EndDate     string `json:"endDate"`     // 结束日期
    IsCurrent   string `json:"isCurrent"`   // 是否当前学期
}
```

### AllSportsResults — GetAllSportsResult 返回值

```go
type AllSportsResults struct {
    Msg  string                 `json:"msg"`
    Code string                 `json:"code"`
    Data []AllSportsResultsData `json:"data"`
}
type AllSportsResultsData struct {
    SportsResultNo  string      `json:"sportsResultNo"`  // 记录编号
    SportsType      string      `json:"sportsType"`      // 运动类型
    IsValid         string      `json:"isValid"`         // 是否有效
    SportsStartTime string      `json:"sportsStartTime"` // 开始时间
    SportsEndTime   string      `json:"sportsEndTime"`   // 结束时间
    PlaceName       string      `json:"placeName"`       // 场地名称
    ReckonType      interface{} `json:"reckonType"`      // 结算类型
    Duration        float64     `json:"duration"`        // 时长（秒）
    Distance        float64     `json:"distance"`        // 距离（公里）
    IsAppeal        string      `json:"isAppeal"`        // 是否申诉
    Reason          string      `json:"reason"`          // 原因
}
```

### StartRunningResp — StartRunning 返回值

```go
type StartRunningResp struct {
    Msg  string `json:"msg"`
    Code string `json:"code"`
    Data string `json:"data"` // 跑步记录编号（runningEventID）
}
```

### CurrentRunningInfo — GetCurrentRunningInfo 返回值

```go
type CurrentRunningInfo struct {
    Msg  string             `json:"msg"`
    Code string             `json:"code"`
    Data CurrentRunningData `json:"data"`
}
type CurrentRunningData struct {
    SportRecordNo               string         `json:"sportRecordNo"`               // 跑步记录编号
    UnifyId                     string         `json:"unifyId"`                     // 统一 ID
    StudentName                 string         `json:"studentName"`                 // 学生姓名
    Mileage                     float64        `json:"mileage"`                     // 当前里程（公里）
    TimeConsuming               int64          `json:"timeConsuming"`               // 耗时（秒）
    ExpiredCountInForbiddenArea int64          `json:"expiredCountInForbiddenArea"` // 禁区停留次数
    Points                      []PointsEntity `json:"points"`                      // 轨迹点列表
    RunningCount                string         `json:"runningCount"`                // 跑步次数
}
type PointsEntity struct {
    Id            interface{} `json:"id"`
    SportPointNo  string      `json:"sportPointNo"`  // 点位编号
    SportRecordNo string      `json:"sportRecordNo"` // 跑步记录编号
    Longitude     string      `json:"longitude"`     // 经度
    Latitude      string      `json:"latitude"`      // 纬度
    CollectTime   string      `json:"collectTime"`   // 采集时间
    IsValid       string      `json:"isValid"`       // 是否有效
    AfterPoints   interface{} `json:"afterPoints"`
    CreateBy      string      `json:"createBy"`
    CreateName    string      `json:"createName"`
    CreateTime    string      `json:"createTime"`
    UpdateBy      interface{} `json:"updateBy"`
    UpdateName    interface{} `json:"updateName"`
    UpdateTime    interface{} `json:"updateTime"`
    Remark        string      `json:"remark"`
}
```

### SportPointListNode — UpdateRunningPoint 请求参数

```go
type SportPointListNode struct {
    SportRecordNo string  `json:"sportRecordNo"` // 跑步记录编号（可空）
    Longitude     float64 `json:"longitude"`     // 经度
    Latitude      float64 `json:"latitude"`      // 纬度
    PlaceName     string  `json:"placeName"`     // 场地名称
    PlaceCode     string  `json:"placeCode"`     // 场地代码
    CollectTime   string  `json:"collectTime"`   // 采集时间 "2006-01-02 15:04:05" 格式
    IsValid       string  `json:"isValid"`       // 是否有效 ("1"=有效)
}
```

### RunningLocationPointResp — UpdateRunningPoint 返回值

```go
type RunningLocationPointResp struct {
    Msg  string                       `json:"msg"`
    Code string                       `json:"code"`
    Data RunningLocationPointRespData `json:"data"`
}
type RunningLocationPointRespData struct {
    SportRecordNo               string      `json:"sportRecordNo"`               // 跑步记录编号
    UnifyId                     string      `json:"unifyId"`                     // 统一 ID
    StudentName                 string      `json:"studentName"`                 // 学生姓名
    Mileage                     float64     `json:"mileage"`                     // 当前里程（公里）
    TimeConsuming               int64       `json:"timeConsuming"`               // 耗时（秒）
    ExpiredCountInForbiddenArea int64       `json:"expiredCountInForbiddenArea"` // 禁区停留次数
    Points                      interface{} `json:"points"`
    RunningCount                string      `json:"runningCount"`                // 跑步次数
}
```

### FinishRunningResp — FinishRunning 返回值

```go
type FinishRunningResp struct {
    Msg  string `json:"msg"`
    Code string `json:"code"`
    Data bool   `json:"data"` // true=成功, false=失败
}
```

---

## 典型使用流程

```go
func main() {
    s, _ := sdk.NewSDK(sdk.WithCookiePath("./cookies.json"))
    defer s.Close()

    // 1. 获取 Token
    tokenInfo, _ := s.GetToken("your_openid")
    token := tokenInfo.Data.Token

    // 2. 获取学期
    terms, _ := s.GetTerm(token)

    // 3. 开始跑步
    startResp, _ := s.StartRunning(token, s.GetFieldID_FengHua())
    runID := startResp.Data

    // 4. 循环上报位置点
    points := []models.SportPointListNode{{
        Longitude: 106.60525, Latitude: 29.53286,
        PlaceName: "风华运动场", PlaceCode: "T1001",
        CollectTime: time.Now().Format("2006-01-02 15:04:05"),
        IsValid: "1",
    }}
    s.UpdateRunningPoint(token, runID, points)

    // 5. 结束跑步
    s.FinishRunning(token, runID)
}
```

## 单元测试

测试文件 `unity_test.go` 包含 10 个测试用例，覆盖所有 SDK 函数：

| 测试函数 | 测试内容 |
|----------|----------|
| `TestSDK_GetToken` | 通过环境变量 `OPEN_ID` 获取 Token |
| `TestSDK_GetUserInfo` | 通过环境变量 `TOKEN` 获取用户信息 |
| `TestSDK_GetNotice` | 通过环境变量 `TOKEN` 获取通知 |
| `TestSDK_GetAllArea` | 通过环境变量 `TOKEN` 获取所有运动场地 |
| `TestSDK_GetTerm` | 通过环境变量 `TOKEN` 获取学期列表 |
| `TestSDK_GetAllSportsResult` | 通过环境变量 `TOKEN` + 学期号获取锻炼记录 |
| `TestSDK_StartRunning` | 通过环境变量 `TOKEN` 开始跑步 |
| `TestSDK_GetCurrentRunningInfo` | 获取当前跑步详情 |
| `TestSDK_UpdateRunningPoint` | 从 `test_location.json` 读取坐标数据，循环上报轨迹点 |
| `TestSDK_FinishRunning` | 结束跑步 |

### 运行测试

```bash
# 运行所有测试
go test -v

# 运行单个测试
go test -v -run TestSDK_GetToken
```

> 部分测试需要设置环境变量 `OPEN_ID` 和 `TOKEN`，且需要有效的 Cookie 文件 `test_cookies.json`。

## License

MIT License — 详见 [LICENSE](./LICENSE)
