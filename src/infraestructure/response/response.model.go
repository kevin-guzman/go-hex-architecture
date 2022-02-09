package response

type ResponseModel struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Err     string      `json:"error"`
	Message string      `json:"message"`
}
