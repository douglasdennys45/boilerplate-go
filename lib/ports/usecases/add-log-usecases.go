package usecases

type AddLogUsecases interface {
	Add(data []byte) Response
}

type Error struct {
	Code   string `json:"code"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type Response struct {
	Data  interface{} `json:"data"`
	Error *Error      `json:"error"`
}
