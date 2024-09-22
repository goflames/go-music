package common

// Response 统一返回数据结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// Success 成功响应，不带数据
func Success(message string) Response {
	return Response{
		Code:    200,
		Message: message,
		Type:    "success",
		Success: true,
		Data:    nil,
	}
}

// SuccessWithData 成功响应，带数据
func SuccessWithData(message string, data interface{}) Response {
	return Response{
		Code:    200,
		Message: message,
		Type:    "success",
		Success: true,
		Data:    data,
	}
}

// Warning 警告响应
func Warning(message string) Response {
	return Response{
		Code:    200,
		Message: message,
		Type:    "warning",
		Success: false,
		Data:    nil,
	}
}

// Error 错误响应
func Error(message string) Response {
	return Response{
		Code:    400,
		Message: message,
		Type:    "error",
		Success: false,
		Data:    nil,
	}
}

// Fatal 致命错误响应
func Fatal(message string) Response {
	return Response{
		Code:    500,
		Message: message,
		Type:    "fatal",
		Success: false,
		Data:    nil,
	}
}
