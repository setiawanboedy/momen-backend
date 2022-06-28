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

type ResponseTrans struct {
	Meta Meta `json:"meta"`
	TotalTransaction int `json:"total_transaction"`
	Data interface{} `json:"data"`
}

func APIResponseTransactions(meta Meta, totalTrans int, data interface{}) ResponseTrans  {
	metaData := Meta{
		Message: meta.Message,
		Code: meta.Code,
		Status: meta.Status,
	}

	responsData := ResponseTrans{
		Meta: metaData,
		TotalTransaction: totalTrans,
		Data: data,
	}
	return responsData
}