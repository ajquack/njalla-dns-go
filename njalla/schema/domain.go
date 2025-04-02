package schema

type UpdateDomainParams struct {
	Domain         string `json:"domain"`
	MailForwarding bool   `json:"mailforwarding"`
	DNSSEC         bool   `json:"dnssec"`
	Lock           bool   `json:"lock"`
}

type FindDomainParams struct {
	Query string `json:"query"`
}

type GetDomainParams struct {
	Domain string `json:"domain"`
}

type ListDomainParams struct {
}

type UpdateDomainRequest struct {
	Method string             `json:"method"`
	Params UpdateDomainParams `json:"params"`
}

type UpdateDomainRequestResponse struct {
	Name           string `json:"name"`
	Status         string `json:"status"`
	Expiry         string `json:"expiry"`
	Autorenew      bool   `json:"autorenew"`
	Locked         bool   `json:"locked"`
	Mailforwarding bool   `json:"mailforwarding"`
	MaxNameservers int    `json:"maxnameservers"`
	DNSSECType     string `json:"dnssec_type"`
	MaxStaticPages int    `json:"maxstaticpages"`
}

type FindDomainRequest struct {
	Method string           `json:"method"`
	Params FindDomainParams `json:"params"`
}

type FindDomainRequestResponse struct {
	Domains []FindDomainResponse `json:"domains"`
}

type FindDomainResponse struct {
	Price  int    `json:"price"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

type GetDomainRequest struct {
	Method string          `json:"method"`
	Params GetDomainParams `json:"params"`
}

type GetDomainRequestResponse struct {
	Name           string `json:"name"`
	Status         string `json:"status"`
	Expiry         string `json:"expiry"`
	Autorenew      bool   `json:"autorenew"`
	Locked         bool   `json:"locked"`
	Mailforwarding bool   `json:"mailforwarding"`
	MaxNameservers int    `json:"max_nameservers"`
	DNSSECType     string `json:"dnssec_type"`
	MaxStaticPages int    `json:"max_static_pages"`
}

type ListDomainsRequest struct {
	Method string           `json:"method"`
	Params ListDomainParams `json:"params"`
}

type ListDomainsRequestResponse struct {
	Domains []ListDomainResponse `json:"domains"`
}

type ListDomainResponse struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	Expiry    string `json:"expiry"`
	Autorenew bool   `json:"autorenew"`
}
