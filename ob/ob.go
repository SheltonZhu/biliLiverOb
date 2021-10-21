package ob

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	OFFLINE = iota
	LIVING
)

type responseData struct {
	Code     int      `json:"code"`
	Message  string   `json:"message"`
	UserInfo UserInfo `json:"data"`
}
type UserInfo struct {
	Mid      int      `json:"mid"`
	Name     string   `json:"name"`
	Face     string   `json:"face"`
	LiveRoom LiveRoom `json:"live_room"`
}
type LiveRoom struct {
	LiveStatus int    `json:"liveStatus"`
	Url        string `json:"url"`
	Title      string `json:"title"`
	Cover      string `json:"cover"`
	RoomId     int    `json:"roomid"`
}

func NewOb(mid int) (*UserInfo, error) {
	u := UserInfo{Mid: mid}
	if err := u.Fill(); err != nil {
		return nil, err
	}
	return &u, nil
}
func (u *UserInfo) Fill() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bilibili.com/x/space/acc/info?mid=%d&jsonp=jsonp", u.Mid), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Origin", "https://space.bilibili.com")
	req.Header.Add("Referer", "https://space.bilibili.com")
	req.Header.Add("Host", "api.bilibili.com")
	clit := http.Client{}
	resp, err := clit.Do(req)
	if err != nil {
		return err
	}
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var r responseData
	err = json.Unmarshal(respByte, &r)
	if err != nil {
		return err
	}
	if r.Code == 0 {
		user := r.UserInfo
		u.Name = user.Name
		u.Face = user.Face
		u.LiveRoom = user.LiveRoom
		return nil
	}
	return errors.New(r.Message)
}
func (u *UserInfo) IsLiving() bool {
	return u.LiveRoom.LiveStatus == LIVING
}
