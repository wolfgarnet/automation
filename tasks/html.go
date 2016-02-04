package tasks

import (
	"github.com/wolfgarnet/pageparser"
	"fmt"
	"github.com/wolfgarnet/automation/system"
	"github.com/wolfgarnet/typeutils"
	"reflect"
	"net/url"
)

type URL struct {
	crawler html2.Crawler
	URL string
	URLExtensions []string
	FileTypes []string
}

func (d *URL) Execute(tr *automation.TaskRunner, cache automation.Cache) error {
	result := d.crawler.Crawl(d.URL)
	//result.Process()
	switch r := result.(type) {
	case *html2.PageResult:
		for _, p := range r.Links {
			fmt.Printf("TEST: %v\n", p.URL.Host)
		}
	}
	return nil
}

func isIn(value string, values []string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}

func (d *URL) Finalize(failed bool) error {
	return nil
}

func (d URL) String() string {
	return "URL"
}

func init() {
	//registry.AddFactory("display", NewDisplay)
	typeutils.RegisterType("url", reflect.TypeOf(Display{}))
}
