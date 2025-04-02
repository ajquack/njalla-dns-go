package schema

type UdpateDomainParams struct {
	Domain         string   `json:"domain"`
	MailForwarding bool     `json:"mailforwarding"`
	DNSSEC         bool     `json:"dnssec"`
	Lock           bool     `json:"lock"`
	Nameservers    []string `json:"nameservers"`
}

type FindDomainResponse struct {
	Price  int    `json:"price"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

type FindDomainParams struct {
	Query string `json:"query"`
}

type GetDomainParams struct {
	Domain string `json:"domain"`
}

type UpdateDomainRequest struct {
	Method string             `json:"method"`
	Params UdpateDomainParams `json:"params"`
}

type ListDomainResponse struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	Expiry    string `json:"expiry"`
	Autorenew bool   `json:"autorenew"`
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
	MaxNameservers int    `json:"maxnameservers"`
	DNSSECType     string `json:"dnssec_type"`
	MaxStaticPages int    `json:"maxstaticpages"`
}

type ListDomainsRequest struct {
	Method string `json:"method"`
	Params struct {
	} `json:"params"`
}

type ListDomainsRequestResponse struct {
	Domains []ListDomainResponse `json:"domains"`
}
