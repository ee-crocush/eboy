package websocket

type RequestData struct {
	UserId int `json:"user_id"`
}

type Request struct {
	Type string      `json:"type"`
	Data RequestData `json:"data"`
}

type Action struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

type ResponseData struct {
	SessionId int      `json:"session_id"`
	UserId    int      `json:"user_id"`
	AreaId    int      `json:"area_id"`
	Actions   []Action `json:"actions"`
	Timestamp string   `json:"timestamp"`
}

type Response struct {
	Type string       `json:"type"`
	Data ResponseData `json:"data"`
}
