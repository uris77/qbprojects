package projects

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type ProjectDoc struct {
	XMLName  xml.Name  `xml:"qdbapi"`
	ErrCode  int       `xml:"errcode"`
	Projects []Project `xml:"record"`
}

type Project struct {
	ProjectId             int64  `xml:"project_id"`
	ProjectStatus         string `xml:"project_status"`
	Partner               string `xml:"partner"`
	Client                string `xml:"client"`
	ProjectName           string `xml:"project_name"`
	Product               string `xml:"product"`
	RelatedClient         int64  `xml:"related_client"`
	SDate                 string `xml:"start_date"`
	EDate                 string `xml:"end_date"`
	ContractStatus        string `xml:"contract_status"`
	ClientBillable        bool   `xml:"client_billable_"`
	AssociatedContractIds string `xml:"associated_contract_ids"`
	Vertical              string `xml:"vertical"`
	HsProduct             string `xml:"hs_product"`
	PrebillMedia          string `xml:"prebill_media"`
	PrebillManagement     string `xml:"prebill_management"`
	ContractType          string `xml:"contract_type"`
	DealType              string `xml:"deal_type"`
}

func (b *Project) EndDate() pq.NullTime {
	if len(b.EDate) == 0 {
		return pq.NullTime{Valid: false}
	}
	return pq.NullTime{Time: EpochToDate(b.EDate), Valid: true}
}

func (b *Project) StartDate() pq.NullTime {
	if len(b.SDate) == 0 {
		return pq.NullTime{Valid: false}
	}
	return pq.NullTime{Time: EpochToDate(b.SDate), Valid: true}
}

func UnmarshallProject(raw string) ProjectDoc {
	var proj ProjectDoc
	if err := xml.Unmarshal([]byte(raw), &proj); err != nil {
		Logger.Fatalw("An error occurred unmarshalling the projects", "error", err)
	}
	return proj
}

type AuthResult struct {
	XMLName xml.Name `xml:"qdbapi"`
	Errcode int      `xml:"errcode"`
	Errtxt  string   `xml:"errtext"`
	Ticket  string   `xml:"ticket"`
}

func Auth(username string, password string) AuthResult {
	client := &http.Client{}
	body := authBody(username, password)
	url := "https://befoundonline.quickbase.com/db/main"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		Logger.Fatalw("An error occurred when creating a request to retrieve a ticket", "error", err)
	}

	req.Header.Add("content-type", "application/xml")
	req.Header.Add("accept", "application/xml")
	req.Header.Add("QUICKBASE-ACTION", "API_Authenticate")

	resp, err := client.Do(req)
	if err != nil {
		Logger.Fatalw("Auth Post failed", "error", err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			Logger.Fatalw("Crashed while closing the body for the auth response", "error", err)
		}
	}()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	var result AuthResult
	err = xml.Unmarshal([]byte(bodyString), &result)
	if err != nil {
		Logger.Fatalw("Failed ot unmarshal the auth body", "error", err)
	}
	return result
}

type ApiAuth struct {
	XMLName xml.Name `xml:"qdbapi"`
	Ticket  string   `xml:"ticket"`
}

func authBody(username string, password string) string {
	return fmt.Sprintf("<qdbapi><username>%s</username><password>%s</password></qdbapi>", username, password)
}

type ApiCount struct {
	XMLName xml.Name `xml:qdbapi`
	Total   int      `xml:"numMatches"`
}

func UnmarshalApiCount(raw string) ApiCount {
	var data ApiCount
	xml.Unmarshal([]byte(raw), &data)
	return data
}

func CntQuery(ticket string, token string) string {
	return fmt.Sprintf("<qdbapi><ticket>%s</ticket><apptoken>%s</apptoken><query>{1.IR.'last 5 y '}OR{2.IR. 'this y'}</query></qdbapi>", ticket, token)
}

func CountQbProjects(ticket string, table string) ApiCount {
	client := &http.Client{}
	pcQuery := CntQuery(ticket, "bue6purcydb2dwcggn43udiycga7")
	url := fmt.Sprintf("https://befoundonline.quickbase.com/db/%s", table)
	Logger.Debugw("Counting Projects", "countProjectsQuery", bytes.NewBuffer([]byte(pcQuery)))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(pcQuery)))
	if err != nil {
		Logger.Fatalw("Creating Count Projects Request Failed", "error", err)
	}

	req.Header.Add("content-type", "application/xml")
	req.Header.Add("accept", "application/xml")
	req.Header.Add("QUICKBASE-ACTION", "API_DoQueryCount")

	resp, err := client.Do(req)
	if err != nil {
		Logger.Fatalw("Count Projects Post Failed", "error", err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	return UnmarshalApiCount(bodyString)
}

func EpochToDate(e string) time.Time {
	dsec := e[:10]
	dnsec := e[10:]
	sec, _ := strconv.ParseInt(dsec, 10, 64)
	nsec, _ := strconv.ParseInt(dnsec, 10, 64)
	return time.Unix(sec, nsec)
}

func ProjectsQuery(ticket string, appToken string, size int, offset int) string {
	return fmt.Sprintf("<qdbapi><ticket>%s</ticket><apptoken>%s</apptoken><query>{1.IR.'last 5 y '}OR{2.IR. 'this y'}</query><includeRids>1</includeRids><clist>a</clist><slist>3</slist><options>num-%d.skp-%d.sortorder-A</options></qdbapi>", ticket, appToken, size, offset)
}

func GetProjects(ticket string, table string, appToken string, size int, offset int) ProjectDoc {
	client := &http.Client{}
	q := ProjectsQuery(ticket, appToken, size, offset)
	baseUrl := fmt.Sprintf("https://befoundonline.quickbase.com/db/%s", table)
	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer([]byte(q)))
	if err != nil {
		Logger.Fatalw("Failed to create request for retrieving budgets", "error", err)
	}
	req.Header.Add("content-type", "application/xml")
	req.Header.Add("accept", "application/xml")
	req.Header.Add("QUICKBASE-ACTION", "API_DoQuery")

	resp, err := client.Do(req)
	if err != nil {
		Logger.Fatalw("Posting Request for GetProjectsFailed Failed", "error", err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	return UnmarshallProject(bodyString)
}
