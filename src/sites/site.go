package sites

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Site struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Code  int
}

type Result struct {
	Site   *Site
	Errors []string
}

func NewResult(site *Site) *Result {
	return &Result{Site: site}
}

func (result *Result) HasError() bool {
	return len(result.Errors) > 0
}

func (result *Result) AddError(message string) {
	result.Errors = append(result.Errors, message)
}

func (site *Site) Process() *Result {

	result := NewResult(site)

	response, err := http.Get(site.Url)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != site.Code {
		result.AddError("Response code test: expected(" + strconv.Itoa(response.StatusCode) + ") but given(" + strconv.Itoa(site.Code) + ")")
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}

	title := document.Find("title").Text()
	if bytes.Compare([]byte(site.Title), []byte(title)) != 0 {
		result.AddError("Title compare: expected(" + site.Title + ") but given(" + title + ")")
	}

	return result
}