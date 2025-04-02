package schema

type GlueResponse struct {
	Name     string `json:"name"`
	Address4 string `json:"address4"`
	Address6 string `json:"address6"`
}

type GlueParams struct {
	Domain   string `json:"domain"`
	Name     string `json:"name"`
	Address4 string `json:"address4,omitempty"`
	Address6 string `json:"address6,omitempty"`
}

type GlueListParams struct {
	Domain string `json:"domain"`
}

type GlueDeleteParams struct {
	Domain string `json:"domain"`
	Name   string `json:"name"`
}

type GlueCreateRequest struct {
	Method string     `json:"method"`
	Params GlueParams `json:"params"`
}

type GlueCreateRequestResponse struct {
}

type GlueListRequest struct {
	Method string         `json:"method"`
	Params GlueListParams `json:"params"`
}

type GlueListRequestResponse struct {
	Glue []GlueResponse `json:"glue"`
}

type GlueUpdateRequest struct {
	Method string     `json:"method"`
	Params GlueParams `json:"params"`
}

type GlueUpdateRequestResponse struct {
}

type GlueDeleteRequest struct {
	Method string           `json:"method"`
	Params GlueDeleteParams `json:"params"`
}

type GlueDeleteRequestResponse struct {
}
