package CQUPT_Sports_SDK

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/Auto-CQUPT-Plan/CQUPT-Sports-SDK/consts"
	"github.com/Auto-CQUPT-Plan/CQUPT-Sports-SDK/models"
)

type TestPointData struct {
	Msg  string       `json:"msg"`
	Code string       `json:"code"`
	Data []DataEntity `json:"data"`
}

type DataEntity struct {
	Id            string      `json:"id"`
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

func TestSDK_GetToken(t *testing.T) {
	sdk, err := NewSDK(WithCookiePath("./test_cookies.json"))
	if err != nil {
		t.Fatalf("function NewSDK error: %v", err.Error())
	}

	tokenData, err := sdk.GetToken(os.Getenv("OPEN_ID"))
	if err != nil {
		t.Fatalf("function sdk.GetToken error: %v", err.Error())
	}

	jsonBytes, err := json.MarshalIndent(tokenData, "", "  ")
	if err != nil {
		t.Errorf("func json.MarshalIndent error: %v", err.Error())
	}

	t.Logf("%v", string(jsonBytes))

	err = sdk.Close()
	if err != nil {
		t.Errorf("func sdk.Close error: %v", err.Error())
	}
}

func TestSDK_GetUserInfo(t *testing.T) {
	sdk, err := NewSDK(WithCookiePath("./test_cookies.json"))
	if err != nil {
		t.Fatalf("function NewSDK error: %v", err.Error())
	}

	tokenData, err := sdk.GetUserInfo(os.Getenv("TOKEN"))
	if err != nil {
		t.Fatalf("function sdk.GetUserInfo error: %v", err.Error())
	}

	jsonBytes, err := json.MarshalIndent(tokenData, "", "  ")
	if err != nil {
		t.Errorf("func json.MarshalIndent error: %v", err.Error())
	}

	t.Logf("%v", string(jsonBytes))

	err = sdk.Close()
	if err != nil {
		t.Errorf("func sdk.Close error: %v", err.Error())
	}
}

func TestSDK_GetNotice(t *testing.T) {
	sdk, err := NewSDK(WithCookiePath("./test_cookies.json"))
	if err != nil {
		t.Fatalf("function NewSDK error: %v", err.Error())
	}

	tokenData, err := sdk.GetNotice(os.Getenv("TOKEN"))
	if err != nil {
		t.Fatalf("function sdk.GetNotice error: %v", err.Error())
	}

	jsonBytes, err := json.MarshalIndent(tokenData, "", "  ")
	if err != nil {
		t.Errorf("func json.MarshalIndent error: %v", err.Error())
	}

	t.Logf("%v", string(jsonBytes))

	err = sdk.Close()
	if err != nil {
		t.Errorf("func sdk.Close error: %v", err.Error())
	}
}

func TestSDK_GetAllArea(t *testing.T) {
	sdk, err := NewSDK(WithCookiePath("./test_cookies.json"))
	if err != nil {
		t.Fatalf("function NewSDK error: %v", err.Error())
	}

	tokenData, err := sdk.GetAllArea(os.Getenv("TOKEN"))
	if err != nil {
		t.Fatalf("function sdk.GetAllArea error: %v", err.Error())
	}

	jsonBytes, err := json.MarshalIndent(tokenData, "", "  ")
	if err != nil {
		t.Errorf("func json.MarshalIndent error: %v", err.Error())
	}

	t.Logf("%v", string(jsonBytes))

	err = sdk.Close()
	if err != nil {
		t.Errorf("func sdk.Close error: %v", err.Error())
	}
}

func TestSDK_GetTerm(t *testing.T) {
	sdk, err := NewSDK(WithCookiePath("./test_cookies.json"))
	if err != nil {
		t.Fatalf("function NewSDK error: %v", err.Error())
	}

	tokenData, err := sdk.GetTerm(os.Getenv("TOKEN"))
	if err != nil {
		t.Fatalf("function sdk.GetTerm error: %v", err.Error())
	}

	jsonBytes, err := json.MarshalIndent(tokenData, "", "  ")
	if err != nil {
		t.Errorf("func json.MarshalIndent error: %v", err.Error())
	}

	t.Logf("%v", string(jsonBytes))

	err = sdk.Close()
	if err != nil {
		t.Errorf("func sdk.Close error: %v", err.Error())
	}
}

func TestSDK_GetAllSportsResult(t *testing.T) {
	var term = "20252"

	sdk, err := NewSDK(WithCookiePath("./test_cookies.json"))
	if err != nil {
		t.Fatalf("function NewSDK error: %v", err.Error())
	}

	tokenData, err := sdk.GetAllSportsResult(os.Getenv("TOKEN"), term)
	if err != nil {
		t.Fatalf("function sdk.GetAllSportsResult error: %v", err.Error())
	}

	jsonBytes, err := json.MarshalIndent(tokenData, "", "  ")
	if err != nil {
		t.Errorf("func json.MarshalIndent error: %v", err.Error())
	}

	t.Logf("%v", string(jsonBytes))

	err = sdk.Close()
	if err != nil {
		t.Errorf("func sdk.Close error: %v", err.Error())
	}
}

func TestSDK_StartRunning(t *testing.T) {
	sdk, err := NewSDK(WithCookiePath("./test_cookies.json"))
	if err != nil {
		t.Fatalf("function NewSDK error: %v", err.Error())
	}

	tokenData, err := sdk.StartRunning(os.Getenv("TOKEN"), sdk.GetFieldID_FengHua())
	if err != nil {
		t.Fatalf("function sdk.StartRunning error: %v", err.Error())
	}

	jsonBytes, err := json.MarshalIndent(tokenData, "", "  ")
	if err != nil {
		t.Errorf("func json.MarshalIndent error: %v", err.Error())
	}

	t.Logf("%v", string(jsonBytes))

	err = sdk.Close()
	if err != nil {
		t.Errorf("func sdk.Close error: %v", err.Error())
	}
}

func TestSDK_GetCurrentRunningInfo(t *testing.T) {
	var runEventID = ""

	sdk, err := NewSDK(WithCookiePath("./test_cookies.json"))
	if err != nil {
		t.Fatalf("function NewSDK error: %v", err.Error())
	}

	tokenData, err := sdk.GetCurrentRunningInfo(os.Getenv("TOKEN"), runEventID)
	if err != nil {
		t.Fatalf("function sdk.GetCurrentRunningInfo error: %v", err.Error())
	}

	jsonBytes, err := json.MarshalIndent(tokenData, "", "  ")
	if err != nil {
		t.Errorf("func json.MarshalIndent error: %v", err.Error())
	}

	t.Logf("%v", string(jsonBytes))

	err = sdk.Close()
	if err != nil {
		t.Errorf("func sdk.Close error: %v", err.Error())
	}
}

func TestSDK_UpdateRunningPoint(t *testing.T) {
	var runEventID = ""
	var FieldID = "T1001"
	var index = 0

	var testLocationData TestPointData

	sdk, err := NewSDK(WithCookiePath("./test_cookies.json"))
	if err != nil {
		t.Fatalf("function NewSDK error: %v", err.Error())
	}

	contentBytes, err := os.ReadFile("./test_location.json")
	if err != nil {
		t.Fatalf("function os.ReadFile error: %v", err.Error())
	}

	err = json.Unmarshal(contentBytes, &testLocationData)
	if err != nil {
		t.Fatalf("function json.Unmarshal error: %v", err.Error())
	}

	data := testLocationData.Data
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	testLocationData.Data = data

	for i := 0; i < 10; i++ {

		locationPointList := make([]models.SportPointListNode, 0, 10)

		for j := 0; j < 10; j++ {
			data := testLocationData.Data[index]

			longitude, err := strconv.ParseFloat(data.Longitude, 64)
			if err != nil {
				t.Logf("function strconv.ParseFloat(data.Longitude, 64) error: %v", err.Error())
				continue
			}

			latitude, err := strconv.ParseFloat(data.Latitude, 64)
			if err != nil {
				t.Logf("function strconv.ParseFloat(data.Latitude, 64) error: %v", err.Error())
				continue
			}

			_data := models.SportPointListNode{
				SportRecordNo: "",
				Longitude:     longitude,
				Latitude:      latitude,
				PlaceName:     consts.SportsField[FieldID],
				PlaceCode:     FieldID,
				CollectTime:   time.Now().Format("2006-01-02 15:04:05"),
				IsValid:       "1",
			}

			locationPointList = append(locationPointList, _data)

			index++

			time.Sleep(time.Second)
		}

		PostResp, err := sdk.UpdateRunningPoint(os.Getenv("TOKEN"), runEventID, locationPointList)
		if err != nil {
			t.Errorf("function sdk.UpdateRunningPoint error: %v", err.Error())
		}

		jsonBytes, err := json.MarshalIndent(PostResp, "", "  ")
		if err != nil {
			t.Errorf("func json.MarshalIndent error: %v", err.Error())
		}

		t.Logf("%v", string(jsonBytes))

		time.Sleep(3 * time.Second)
	}

	err = sdk.Close()
	if err != nil {
		t.Errorf("func sdk.Close error: %v", err.Error())
	}
}

func TestSDK_FinishRunning(t *testing.T) {
	var runEventID = ""

	sdk, err := NewSDK(WithCookiePath("./test_cookies.json"))
	if err != nil {
		t.Fatalf("function NewSDK error: %v", err.Error())
	}

	tokenData, err := sdk.FinishRunning(os.Getenv("TOKEN"), runEventID)
	if err != nil {
		t.Fatalf("function sdk.FinishRunning error: %v", err.Error())
	}

	jsonBytes, err := json.MarshalIndent(tokenData, "", "  ")
	if err != nil {
		t.Errorf("func json.MarshalIndent error: %v", err.Error())
	}

	t.Logf("%v", string(jsonBytes))

	err = sdk.Close()
	if err != nil {
		t.Errorf("func sdk.Close error: %v", err.Error())
	}
}
