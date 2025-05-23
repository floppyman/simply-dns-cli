package objects

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/floppyman/um-common/ext"
	"github.com/floppyman/um-common/jsons"

	"github.com/floppyman/simply-dns-cli/internal/styles"
)

type SimplyApiConfig struct {
	Url           string `json:"url"`
	AccountNumber string `json:"account_number"`
	AccountApiKey string `json:"account_api_key"`
}

type SimplyApiDnsRecords struct {
	Records []*SimplyDnsRecord `json:"records"`
}

type SimplyApiProducts struct {
	Products []*SimplyProduct `json:"products"`
}

type SimplyApiSuccessResponse struct {
	Record struct {
		Id int `json:"id"`
	} `json:"record"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

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
	DnsRecTypeSSHFP  DnsRecordType = "SSHFP"
	DnsRecTypeTLSA   DnsRecordType = "TLSA"
	DnsRecTypeTXT    DnsRecordType = "TXT"
)

func DnsTypeToText(dnsType DnsRecordType) string {
	switch dnsType {
	case DnsRecTypeA:
		return "A"
	case DnsRecTypeAAAA:
		return "AAAA"
	case DnsRecTypeALIAS:
		return "ALIAS"
	case DnsRecTypeCAA:
		return "CAA"
	case DnsRecTypeCNAME:
		return "CNAME"
	case DnsRecTypeDNSKEY:
		return "DNSKEY"
	case DnsRecTypeDS:
		return "DS"
	case DnsRecTypeHTTPS:
		return "HTTPS"
	case DnsRecTypeLOC:
		return "LOC"
	case DnsRecTypeMX:
		return "MX"
	case DnsRecTypeNS:
		return "NS"
	case DnsRecTypeSSHFP:
		return "SSHFP"
	case DnsRecTypeTLSA:
		return "TLSA"
	case DnsRecTypeTXT:
		return "TXT"
	}
	return "--"
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

func DnsTTLToText(ttl DnsRecordTTL) string {
	switch ttl {
	case DnsRecTTLMin10:
		return "10 Minutes"
	case DnsRecTTLHour1:
		return "1 Hour"
	case DnsRecTTLHours6:
		return "6 Hours"
	case DnsRecTTLHours12:
		return "12 Hours"
	case DnsRecTTLHours24:
		return "24 Hours"
	}
	return ""
}
func DnsTTLToNumberText(ttl DnsRecordTTL) string {
	switch ttl {
	case DnsRecTTLMin10:
		return fmt.Sprintf("%d", int(DnsRecTTLMin10))
	case DnsRecTTLHour1:
		return fmt.Sprintf("%d", int(DnsRecTTLHour1))
	case DnsRecTTLHours6:
		return fmt.Sprintf("%d", int(DnsRecTTLHours6))
	case DnsRecTTLHours12:
		return fmt.Sprintf("%d", int(DnsRecTTLHours12))
	case DnsRecTTLHours24:
		return fmt.Sprintf("%d", int(DnsRecTTLHours24))
	}
	return ""
}

//goland:noinspection SpellCheckingInspection
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
	RecordId int64            `json:"record_id,omitempty"`
	Name     string           `json:"name"`
	TTL      DnsRecordTTL     `json:"ttl"`
	Data     string           `json:"data"`
	Type     DnsRecordType    `json:"type"`
	Priority *jsons.JsonInt32 `json:"priority"`
	Comment  string           `json:"comment"`
}

func (r SimplyDnsRecord) Print(prefix string) {
	fmt.Printf(`%s%s %s
%s%s %s
%s%s %s
%s%s %s
%s%s %s
%s%s %s
`,
		prefix, styles.GraphicLight("Type:    "), styles.Value(string(r.Type)),
		prefix, styles.GraphicLight("TTL:     "), styles.Value(DnsTTLToText(r.TTL)),
		prefix, styles.GraphicLight("Name:    "), styles.Value(r.Name),
		prefix, styles.GraphicLight("Data:    "), styles.Value(r.Data),
		prefix, styles.GraphicLight("Priority:"), styles.Value(ext.Iif(r.Priority.Valid, string(r.Priority.Value), "<not applicable for Type>")),
		prefix, styles.GraphicLight("Comment: "), styles.Value(ext.Iif(len(r.Comment) > 0, r.Comment, "<no comment>")))
}

func (r SimplyDnsRecord) GetHash() string {
	var priority int32 = -1
	if r.Priority != nil {
		if r.Priority.Valid {
			priority = r.Priority.Value
		}
	}
	data := fmt.Sprintf(
		"%s||%s||%s||%s||%s||%s",
		DnsTypeToText(r.Type),
		r.Name,
		r.Data,
		DnsTTLToText(r.TTL),
		strconv.Itoa(int(priority)),
		r.Comment,
	)
	h := sha256.New()
	h.Write([]byte(data))
	res := make([]byte, 0)
	res = h.Sum(res)
	return string(res)
}
