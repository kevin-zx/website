package websitetool

import (
	"github.com/kevin-zx/websitetool/extract"
	"strings"
)

func GetWebSiteByHost(host string) (*extract.Website, error) {
	ws := extract.Website{}
	homePage, err := extract.ParserPageUrl("http://"+host, true, "")
	if err != nil {
		return nil, err
	}
	homePage.Type = extract.HomePageTypes
	var allSplitText []string
	var wholeText string
	for u, lt := range homePage.AllLinks {
		if strings.HasSuffix(u, host) || strings.HasSuffix(u, host+"/") {
			continue
		}
		t := extract.TestPageType(u, lt, "")
		if t != "" && t != extract.HomePageTypes {
			for _, p := range ws.Pages {
				if u == p.PageUrl {
					continue
				}
			}
			page, err := extract.ParserPageUrl(u, false, "")
			if err != nil {
				continue
			}
			page.Type = t
			if page.PageUrl == "" {
				continue
			}
			ws.Pages = append(ws.Pages, *page)
			allSplitText = append(allSplitText, page.SplitText...)
			wholeText += page.Text
			if len(ws.Pages) > 5 {
				break
			}
		}
	}
	allSplitText = append(allSplitText, homePage.SplitText...)
	wholeText += homePage.Text
	ws.Pages = append(ws.Pages, *homePage)
	cm := extract.TakeCompanyNames(allSplitText, wholeText, homePage.Title)
	ws.CompanyName = getTopCompany(cm)
	ws.SiteUrl = homePage.PageUrl
	return &ws, nil
}

func getTopCompany(m map[string]int) string {
	maxc := 0
	cn := ""
	for k, c := range m {
		if maxc < c {
			cn = k
			maxc = c
		}
	}
	return cn
}
