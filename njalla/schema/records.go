package schema

type RecordResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

type RecordCreateParams struct {
	Domain       string `json:"domain"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	Content      string `json:"content,omitempty"`
	TTL          int    `json:"ttl,omitempty"`
	Prio         int    `json:"prio,omitempty"`
	Weight       int    `json:"weight,omitempty"`
	Port         int    `json:"port,omitempty"`
	Target       string `json:"target,omitempty"`
	SSHAlgorithm int    `json:"ssh_algorithm,omitempty"`
	SSHType      int    `json:"ssh_type,omitempty"`
}

type RecordUpdateParams struct {
	ID           string `json:"id"`
	Domain       string `json:"domain"`
	Type         string `json:"type,omitempty"`
	Name         string `json:"name,omitempty"`
	Content      string `json:"content,omitempty"`
	TTL          int    `json:"ttl,omitempty"`
	Prio         int    `json:"prio,omitempty"`
	Weight       int    `json:"weight,omitempty"`
	Port         int    `json:"port,omitempty"`
	Target       string `json:"target,omitempty"`
	SSHAlgorithm int    `json:"ssh_algorithm,omitempty"`
	SSHType      int    `json:"ssh_type,omitempty"`
}

type RecordListParams struct {
	Domain string `json:"domain"`
}

type RecordDeleteParams struct {
	ID     string `json:"id"`
	Domain string `json:"domain"`
}

type RecordCreateRequest struct {
	Method string             `json:"method"`
	Params RecordCreateParams `json:"params"`
}

type RecordCreateRequestResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

type RecordUpdateRequest struct {
	Method string             `json:"method"`
	Params RecordUpdateParams `json:"params"`
}

type RecordUpdateRequestResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

type RecordsListRequest struct {
	Method string           `json:"method"`
	Params RecordListParams `json:"params"`
}

type RecordsListRequestResponse struct {
	Records []RecordResponse `json:"records"`
}

type RecordDeleteRequest struct {
	Method string             `json:"method"`
	Params RecordDeleteParams `json:"params"`
}

type RecordDeleteRequestResponse struct {
}
