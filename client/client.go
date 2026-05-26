package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"

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
