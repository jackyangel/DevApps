package response

type ShuntingRequest struct {
	Infix   string `json:"infix"`
	Postfix string `json:"postfix"`
	Result  int    `json:"result"`
}
