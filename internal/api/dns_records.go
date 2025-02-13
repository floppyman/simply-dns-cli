package api

type DnsRecordType string

//goland:noinspection SpellCheckingInspection
const (
	DnsRecTypeA      DnsRecordType = "A"
	DnsRecTypeAAAA   DnsRecordType = "AAAA"
	DnsRecTypeALIAS  DnsRecordType = "ALIAS"
	DnsRecTypeCAA    DnsRecordType = "CAA"
	DnsRecTypeCNAME  DnsRecordType = "CNAME"
	DnsRecTypeDNSKEY DnsRecordType = "DNSKEY"
	DnsRecTypeDS     DnsRecordType = "DS"
	DnsRecTypeHTTPS  DnsRecordType = "HTTPS"
	DnsRecTypeLOC    DnsRecordType = "LOC"
	DnsRecTypeMX     DnsRecordType = "MX"
	DnsRecTypeNS     DnsRecordType = "NS"
	DnsRecTypeSRV    DnsRecordType = "SRV"
	DnsRecTypeSSHFP  DnsRecordType = "SSHFP"
	DnsRecTypeTLSA   DnsRecordType = "TLSA"
	DnsRecTypeTXT    DnsRecordType = "TXT"
)

var DnsRecordTypes = []DnsRecordType{
	DnsRecTypeA,
	DnsRecTypeAAAA,
	DnsRecTypeALIAS,
	DnsRecTypeCAA,
	DnsRecTypeCNAME,
	DnsRecTypeDNSKEY,
	DnsRecTypeDS,
	DnsRecTypeHTTPS,
	DnsRecTypeLOC,
	DnsRecTypeMX,
	DnsRecTypeNS,
	DnsRecTypeSRV,
	DnsRecTypeSSHFP,
	DnsRecTypeTLSA,
	DnsRecTypeTXT,
}

type DnsRecordTTL int

const (
	min1             = DnsRecordTTL(60)
	DnsRecTTLMin10   = min1 * 10
	DnsRecTTLHour1   = min1 * 60
	DnsRecTTLHours6  = DnsRecTTLHour1 * 6
	DnsRecTTLHours12 = DnsRecTTLHour1 * 12
	DnsRecTTLHours24 = DnsRecTTLHour1 * 24
)

var DnsRecordTTLs = []DnsRecordTTL{
	DnsRecTTLMin10,
	DnsRecTTLHour1,
	DnsRecTTLHours6,
	DnsRecTTLHours12,
	DnsRecTTLHours24,
}

type SimplyProduct struct {
	Object    string `json:"object"`
	Name      string `json:"name"`
	AutoRenew bool   `json:"autorenew"`
	Cancelled bool   `json:"cancelled"`
	Domain    struct {
		Name          string `json:"name"`
		NameIdn       string `json:"name_idn"`
		Managed       bool   `json:"managed"`
		DateRenewDate int    `json:"date_renewdate"`
	} `json:"domain"`
	Product struct {
		Id          int         `json:"id"`
		Name        string      `json:"name"`
		DateCreated int         `json:"date_created"`
		DateExpire  interface{} `json:"date_expire"`
	} `json:"product"`

	DnsRecords []*SimplyDnsRecord `json:"dns_records"`
}

type SimplyDnsRecord struct {
	RecordId int64         `json:"record_id"`
	Name     string        `json:"name"`
	Ttl      DnsRecordTTL  `json:"ttl"`
	Data     string        `json:"data"`
	Type     DnsRecordType `json:"type"`
	Priority *int          `json:"priority"`
	Comment  string        `json:"comment"`
}
