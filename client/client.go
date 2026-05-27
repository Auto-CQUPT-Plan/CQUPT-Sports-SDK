package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"
	"path"

	"github.com/Auto-CQUPT-Plan/CQUPT-Sports-SDK/consts"
	"github.com/Auto-CQUPT-Plan/CQUPT-Sports-SDK/models"
)

func (r *Client) GetToken(openid string) (*models.TokenInfo, error) {
	var tokenInfo models.TokenInfo

	url, err := url2.Parse(consts.GET_TOKEN_URL)
	if err != nil {
		return nil, err
	}

	query := url.Query()
	{
		query.Add("wxCode", "")
		query.Add("openid", openid)
		query.Add("phoneType", consts.PHONE_TYPE)
	}
	url.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", consts.USER_AGENT)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&tokenInfo)
	if err != nil {
		return nil, err
	}

	return &tokenInfo, nil
}

func (r *Client) GetUserInfo(token string) (*models.UserInfo, error) {
	var userInfo models.UserInfo

	url, err := url2.Parse(consts.GET_USER_INFO_URL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", consts.USER_AGENT)
	req.Header.Add("token", token)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func (r *Client) GetNotice(token string) (*models.NoticeInfo, error) {
	var noticeInfo models.NoticeInfo

	url, err := url2.Parse(consts.GET_NOTICE_URL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", consts.USER_AGENT)
	req.Header.Add("token", token)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&noticeInfo)
	if err != nil {
		return nil, err
	}

	return &noticeInfo, nil
}

func (r *Client) GetAllArea(token string) (*models.AreaInfo, error) {
	var areaInfo models.AreaInfo

	url, err := url2.Parse(consts.GET_ALL_AREA_URL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", consts.USER_AGENT)
	req.Header.Add("token", token)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&areaInfo)
	if err != nil {
		return nil, err
	}

	return &areaInfo, nil
}

func (r *Client) GetTerm(token string) (*models.TermInfo, error) {
	var termInfo models.TermInfo

	url, err := url2.Parse(consts.GET_TERM_URL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", consts.USER_AGENT)
	req.Header.Add("token", token)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&termInfo)
	if err != nil {
		return nil, err
	}

	return &termInfo, nil
}

func (r *Client) GetAllSportsResult(token string, term string) (*models.AllSportsResults, error) {
	var allSportsResults models.AllSportsResults

	url, err := url2.Parse(consts.GET_SPORTS_RESULT_URL)
	if err != nil {
		return nil, err
	}

	query := url.Query()
	{
		query.Add("yearTerm", term)
		query.Add("weekly", "")
		query.Add("sportsType", "")
		query.Add("isValid", "")
		query.Add("reckonType", "")
	}
	url.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", consts.USER_AGENT)
	req.Header.Add("token", token)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&allSportsResults)
	if err != nil {
		return nil, err
	}

	return &allSportsResults, nil
}

func (r *Client) StartRunning(token string, fieldID string) (*models.StartRunningResp, error) {
	var startRunningResp models.StartRunningResp

	url, err := url2.Parse(consts.START_RUNNING_URL)
	if err != nil {
		return nil, err
	}

	postBody, err := json.Marshal(models.StartRunningPostData{PlaceName: consts.SportsField[fieldID], PlaceCode: fieldID})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", consts.USER_AGENT)
	req.Header.Add("token", token)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&startRunningResp)
	if err != nil {
		return nil, err
	}

	return &startRunningResp, nil
}

func (r *Client) GetCurrentRunningInfo(token string, runningEventID string) (*models.CurrentRunningInfo, error) {
	var currentRunningInfo models.CurrentRunningInfo

	url, err := url2.Parse(consts.GET_CURRENT_RUNNING_INFO)
	if err != nil {
		return nil, err
	}

	url.Path = path.Join(url.Path, runningEventID)

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", consts.USER_AGENT)
	req.Header.Add("token", token)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&currentRunningInfo)
	if err != nil {
		return nil, err
	}

	return &currentRunningInfo, nil
}

func (r *Client) UpdateRunningPoint(token string, runningEventID string, pointData []models.SportPointListNode) (*models.RunningLocationPointResp, error) {
	var runningLocationPointResp models.RunningLocationPointResp

	url, err := url2.Parse(consts.UPDATE_RUNNING_POINT)
	if err != nil {
		return nil, err
	}

	postData, err := json.Marshal(models.RunningLocationPointPostData{SportRecordNo: runningEventID, SportPointList: pointData})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", consts.USER_AGENT)
	req.Header.Add("token", token)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&runningLocationPointResp)
	if err != nil {
		return nil, err
	}

	return &runningLocationPointResp, nil
}

func (r *Client) FinishRunning(token string, runningEventID string) (*models.FinishRunningResp, error) {
	var finishRunningResp models.FinishRunningResp

	url, err := url2.Parse(consts.FINISH_RUNNING_URL)
	if err != nil {
		return nil, err
	}

	url.Path = path.Join(url.Path, runningEventID)

	req, err := http.NewRequest("POST", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", consts.USER_AGENT)
	req.Header.Add("token", token)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&finishRunningResp)
	if err != nil {
		return nil, err
	}

	return &finishRunningResp, nil
}
