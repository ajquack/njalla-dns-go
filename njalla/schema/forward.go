package schema

type ForwardResponse struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type ForwardCreateParams struct {
	Domain string `json:"domain"`
	From   string `json:"from"`
	To     string `json:"to"`
}

type ForwardListParams struct {
	Domain string `json:"domain"`
}

type ForwardDeleteParams struct {
	Domain string `json:"domain"`
	From   string `json:"from"`
	To     string `json:"to"`
}

type ForwardCreateRequest struct {
	Method string              `json:"method"`
	Params ForwardCreateParams `json:"params"`
}

type ForwardCreateRequestResponse struct {
	Domain string `json:"domain"`
	From   string `json:"from"`
	To     string `json:"to"`
}

type ForwardListRequest struct {
	Method string            `json:"method"`
	Params ForwardListParams `json:"params"`
}

type ForwardListRequestResponse struct {
	Forward []ForwardResponse `json:"forwards"`
}

type ForwardDeleteRequest struct {
	Method string              `json:"method"`
	Params ForwardDeleteParams `json:"params"`
}

type ForwardDeleteRequestResponse struct {
}
