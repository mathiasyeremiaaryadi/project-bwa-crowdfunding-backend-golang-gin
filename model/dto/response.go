package dto

type ResponseContainer struct {
	Meta MetaData    `json:"meta"`
	Data interface{} `json:"data"`
}

type MetaData struct {
	Message interface{} `json:"message"`
	Status  string      `json:"status"`
	Code    int         `json:"code"`
}

func BuildResponse(message string, status string, code int, data interface{}) *ResponseContainer {
	return &ResponseContainer{
		Meta: MetaData{
			Message: message,
			Status:  status,
			Code:    code,
		},
		Data: data,
	}
}
