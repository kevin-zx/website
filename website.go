package websitetool

import "github.com/kevin-zx/websitetool/extract"

func GetWebSiteByHost(host string) (*extract.Website, error) {
	hostUrl := "http://" + host
	ws := extract.Website{}
	homePage, err := extract.ParserPageUrl(hostUrl, "")
	if err != nil {
		return nil, err
	}
	homePage.Type = extract.HomePageTypes
	var allSplitText []string
	var wholeText string
	for u, lt := range homePage.AllLinks {
		if u == hostUrl || u == hostUrl+"/" {
			continue
		}
		t := extract.TestPageType(u, lt, "")
		if t != "" && t != extract.HomePageTypes {
			page, err := extract.ParserPageUrl(hostUrl, "")
			if err != nil {
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

	cm := extract.TakeCompanyNames(allSplitText, wholeText, homePage.Title)
	ws.CompanyName = getTopCompany(cm)
	ws.SiteUrl = hostUrl
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
