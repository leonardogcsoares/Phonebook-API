package router

type loginResp struct {
	Token string `json:"token"`
}

type errorResp struct {
	Msg   string `json:"msg"`
	Error string `json:"error"`
}
