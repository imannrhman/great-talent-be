package model

type WebResponse struct {
	Code     int         `json:"code"`
	Status   string      `json:"status"`
	Messages string      `json:"messages"`
	Data     interface{} `json:"data"`
}
