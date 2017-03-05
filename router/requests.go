package router

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
