package helper

type Response struct {
	Meta Meta `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Status string `json:"status"`
}

func APIResponse(meta Meta, data interface{}) Response {
	metaData := Meta{
		Message: meta.Message,
		Code: meta.Code,
		Status: meta.Status,
	}

	responsData := Response{
		Meta: metaData,
		Data: data,
	}
	return responsData
}