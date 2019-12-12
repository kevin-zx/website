package extract

import (
	"github.com/PuerkitoBio/goquery"
	httpUtil "github.com/kevin-zx/http-util"
	"github.com/kevin-zx/websitetool/regexInfopaser"
	"net/url"
	"strings"
)

type PageInfo struct {
	HomePagePhones    []string
	HomePageTels      []string
	ContactPagePhones []string
	ContactPageTels   []string
	ContactUrl        string
	IcpNo             string
	HomeSiteUrl       string
	Title             string
}

func ParserDomain(domain string) (p PageInfo, err error) {
	p = PageInfo{}
	curl := "http://www." + domain
	wec, err := httpUtil.GetWebConFromUrl(curl)
	p.HomeSiteUrl = "www." + domain
	if err != nil {
		wec, err = httpUtil.GetWebConFromUrl("http://" + domain)
		if err != nil {
			return
		}
		p.HomeSiteUrl = domain
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(wec))
	if err != nil {
		return
	}
	ccurl, _ := url.Parse(curl)

	p.Title = doc.Find("title").Text()
	pageContent := doc.Text()
	icpNos := regexInfopaser.MatchICPNO(pageContent)
	if len(icpNos) > 0 {
		p.IcpNo = icpNos[0]
	}

	p.HomePageTels = regexInfopaser.MatchTelephone(pageContent)
	p.HomePagePhones = regexInfopaser.MatchPhone(pageContent)
	getContactF := false
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		if getContactF {
			return
		}
		aText := s.Text()
		href, _ := s.Attr("href")
		hrefUrl, err := ccurl.Parse(href)
		if err != nil {
			return
		}
		//fmt.Println(hrefUrl.String(),href)
		if strings.Contains(aText, "联系我们") || strings.Contains(strings.ToLower(href), "contact") ||
			strings.Contains(strings.ToLower(href), "lianxi") || strings.Contains(strings.ToLower(href), "lxwm") {
			getContactF = true
			p.ContactPagePhones, p.ContactPageTels, _ = getContactPage(hrefUrl.String())
			p.ContactUrl = hrefUrl.String()
			return
		}
	})

	if !getContactF {
		doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			if getContactF {
				return
			}
			aText := s.Text()
			href, _ := s.Attr("href")
			hrefUrl, err := ccurl.Parse(href)
			if err != nil {
				return
			}
			//fmt.Println(hrefUrl.String(),href)
			if strings.Contains(aText, "关于") ||
				strings.Contains(strings.ToLower(href), "introduce") ||
				strings.Contains(strings.ToLower(href), "about") {
				getContactF = true
				p.ContactPagePhones, p.ContactPageTels, _ = getContactPage(hrefUrl.String())
				p.ContactUrl = hrefUrl.String()
				return
			}
		})
	}
	p.ContactPageTels = removeRepeatedElement(p.ContactPageTels)
	p.ContactPagePhones = removeRepeatedElement(p.ContactPagePhones)
	p.HomePagePhones = removeRepeatedElement(p.HomePagePhones)
	p.HomePageTels = removeRepeatedElement(p.HomePageTels)

	return

}

func getContactPage(contactpageurl string) (phones []string, tels []string, err error) {
	wec, err := httpUtil.GetWebConFromUrl(contactpageurl)
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(wec))
	if err != nil {
		return
	}
	text := doc.Text()
	phones = regexInfopaser.MatchPhone(text)
	tels = regexInfopaser.MatchTelephone(text)
	return
}

func removeRepeatedElement(arr []string) (ret []string) {
	var keywordCount = make(map[string]int)
	a_len := len(arr)
	for i := 0; i < a_len; i++ {
		duFlag := false
		for _, re := range ret {
			if len(arr[i]) == 0 {
				duFlag = true
				break
			}
			if re == arr[i] {
				if _, ok := keywordCount[re]; !ok {
					keywordCount[re] = 1
				}
				duFlag = true
				break
			}
		}
		if !duFlag {
			ret = append(ret, arr[i])
		}
	}
	return
}
