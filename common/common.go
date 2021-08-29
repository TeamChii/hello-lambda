package common

type HeaderResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	StatusCdSuccess = "0000"
)

var StatusDescMap = map[string]string{
	StatusCdSuccess: "SUCCESS",
}
