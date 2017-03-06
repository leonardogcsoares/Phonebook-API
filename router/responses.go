package router

// LoginResp TODO
type LoginResp struct {
	Token string `json:"token"`
}

// ErrorResp TODO
type ErrorResp struct {
	Msg   string `json:"msg"`
	Error string `json:"error"`
}
