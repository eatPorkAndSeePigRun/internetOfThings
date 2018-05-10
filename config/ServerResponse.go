package config

type ServerResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func CreateBySuccess() ServerResponse {
	return ServerResponse{Code: SUCCESS}
}

func CreateByError() ServerResponse {
	return ServerResponse{Code: ERROR, Msg: "ERROR"}
}

func CreateBySuccessMessage1(successMessage string) ServerResponse {
	return ServerResponse{Code: SUCCESS, Msg: successMessage}
}

func CreateBySuccessMessage2(data interface{}) ServerResponse {
	return ServerResponse{Code: SUCCESS, Data: data}
}

func CreateBySuccessMessage3(successMessage string, data interface{}) ServerResponse {
	return ServerResponse{Code: SUCCESS, Msg: successMessage, Data: data}
}

func CreateByErrorMessage(errorMessage string) ServerResponse {
	return ServerResponse{Code: ERROR, Msg: errorMessage}
}

func CreateByErrorCodeMessage(errorCode int, errorMessage string) ServerResponse {
	return ServerResponse{Code: errorCode, Msg: errorMessage}
}
