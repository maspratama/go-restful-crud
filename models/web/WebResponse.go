package web

type WebResponse struct {
	ResponseCode    int
	ResponseMessage string
	Data            interface{}
}
