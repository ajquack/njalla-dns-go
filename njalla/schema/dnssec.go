package schema

type DNSSECCreateParams struct {
	Domain     string `json:"domain"`
	Algorithm  int    `json:"algorithm"`
	Digest     string `json:"digest"`
	DigestType int    `json:"digest_type"`
	KeyTag     int    `json:"key_tag"`
	PublicKey  string `json:"public_key"`
}

type DNSSECListParams struct {
	Domain string `json:"domain"`
}

type DNSSECDeleteParams struct {
	Domain string `json:"domain"`
	ID     string `json:"id"`
}

type DNSSECCreateRequest struct {
	Method string             `json:"method"`
	Params DNSSECCreateParams `json:"params"`
}

type DNSSECCreateRequestResponse struct {
}

type DNSSECListRequest struct {
	Method string           `json:"method"`
	Params DNSSECListParams `json:"params"`
}

type DNSSECListRequestResponse struct {
	DNSSec []struct {
	} `json:"dnssec"`
}

type DNSSECDeleteRequest struct {
	Method string             `json:"method"`
	Params DNSSECDeleteParams `json:"params"`
}
