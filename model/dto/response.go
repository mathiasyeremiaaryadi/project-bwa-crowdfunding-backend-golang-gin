package dto

type ResponseContainer struct {
	Meta MetaData    `json:"META"`
	Data interface{} `json:"DATA"`
}

type MetaData struct {
	Message interface{} `json:"MESSAGE"`
	Status  string      `json:"STATUS"`
	Code    int         `json:"CODE"`
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
