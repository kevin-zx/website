package extract

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	http_util "github.com/kevin-zx/http-util"
	"github.com/kevin-zx/websitetool/regexInfopaser"
	"github.com/xluohome/phonedata"
	"net/url"
	"strings"
	"time"
)

const (
	HomePageTypes    = "首页"
	ContactPageTypes = "联系我们"
	AboutPageTypes   = "关于我们"
)

func parserSiteUrl() {

}

var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var nums = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var punctuations = []string{"©", "·", "=", "-", "，", "。", "、", "；", "’", "【", "】", "、", "`", "！", "@", "#", "￥", "%", "…", "…", "&", "×", "—", "—", "《", "》", "？", "：", "”", "“", "{", "}", "‘", "|", "～", "+", ",", ".", "/", ";", "'", "[", "]", "\\", "`", "!", "@", "#", "$", "%", "^", "&", "*", "_", "+", "<", ">", "?", ":", "\"", "{", "}", "~"}
var emptybrackets = []string{"((", "（（", "()", "（）"}

func ParserPageUrl(host string, pageLinkText string) (wp *WebPage, err error) {

	hostURL := "http://"+host
	wp = &WebPage{}
	wp.PageUrl = host
	html, err := getPagePCHtml(hostURL)
	if err != nil {
		hostURL := "https://"+host
		html, err = getPagePCHtml(hostURL)
		if err != nil {
			return nil, err
		}
	}
	pu, err := url.Parse(hostURL)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	links := map[string]string{}
	doc.Find("a").Each(func(_ int, a *goquery.Selection) {
		href, ok := a.Attr("href")
		if !ok || href == "" || href == "javascript:;" ||
			strings.HasPrefix(href, "#") ||
			strings.HasPrefix(href, "mailto:") ||
			strings.HasPrefix(href, "sms:") ||
			strings.HasPrefix(href, "tel:") ||
			strings.HasPrefix(href, "#") {
			return
		}
		u, err := pu.Parse(href)
		if err != nil || !strings.Contains(u.String(), pu.Host) {
			return
		}

		links[u.String()] = a.Text()

	})
	wp.AllLinks = links
	wp.Title, wp.Description, wp.Keywords = getTDK(doc)
	wp.Type = TestPageType(host, pageLinkText, wp.Title)
	splitText, fullText := getPageContent(doc.Find("html"))
	wp.SplitText = append(splitText, clearHoleText(wp.Description))
	wp.Text = fullText + clearHoleText(wp.Description)
	//htmlText := doc.Find("html").Text()
	phones := regexInfopaser.MatchPhone(strings.Join(wp.SplitText, ","))
	for _, p := range phones {

		pr, err := phonedata.Find(p)
		if err != nil {
			continue
		}
		wp.MobilePhoneNums = append(wp.MobilePhoneNums, MobilePhoneNum{
			PhoneNum:                  pr.PhoneNum,
			TelecommunicationOperator: pr.CardType,
			City:                      pr.City,
			ZipCode:                   pr.ZipCode,
		})
	}
	tels := regexInfopaser.MatchTelephone(strings.Join(wp.SplitText, ","))
	for _, tel := range tels {
		code, area := exportAreaCode(tel)
		if code == "" {
			fmt.Println(tel)
			continue
		}
		wp.TelPhones = append(wp.TelPhones, TelPhoneNum{
			TelNum:   tel,
			AreaCode: area,
			City:     code,
		})
	}
	//phonedata.Find()
	return wp, nil

}

func getPageContent(htmlSelect *goquery.Selection) (splitText []string, wholeText string) {
	var spiltText []string
	spiltText = append(spiltText, TagText(htmlSelect, "span")...)
	spiltText = append(spiltText, TagText(htmlSelect, "a")...)
	spiltText = append(spiltText, TagText(htmlSelect, "b")...)
	spiltText = append(spiltText, TagText(htmlSelect, "em")...)
	spiltText = append(spiltText, TagText(htmlSelect, "p")...)
	spiltText = append(spiltText, TagText(htmlSelect, "h1")...)
	spiltText = append(spiltText, TagText(htmlSelect, "h2")...)
	spiltText = append(spiltText, TagText(htmlSelect, "h3")...)
	spiltText = append(spiltText, TagText(htmlSelect, "h4")...)
	spiltText = append(spiltText, TagText(htmlSelect, "h5")...)
	spiltText = append(spiltText, TagText(htmlSelect, "td")...)
	spiltText = append(spiltText, TagText(htmlSelect, "title")...)
	fullText := htmlSelect.Text()

	return spiltText, clearHoleText(fullText)
}

func TagText(s *goquery.Selection, selector string) (text []string) {
	sel := s.Find(selector)
	sel.Each(func(_ int, se *goquery.Selection) {
		selText := se.Text()
		if se.Children().Size() > 0 {
			se.Children().Each(func(_ int, s *goquery.Selection) {
				selText = strings.Replace(selText, s.Text(), "", -1)
			})
		}

		//fmt.Println(has,"-",se.TagText(),"-",selText)
		//fmt.Printf("%t `%s` `%s`\n",has,se.TagText(),selText)
		texts := clear(selText)
		if len(texts) > 0 {
			text = append(text, texts...)
		}
	})
	return text
}

func clear(text string) []string {
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\t", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	// 一个很奇怪的空白字符 0xc2, 0xa0  或者 194 160
	text = strings.Replace(text, " ", "", -1)
	//text = strings.Replace(text,"\n","",-1)
	//text = strings.Replace(text," ","",-1)
	text = strings.Replace(text, "（", "(", -1)
	text = strings.Replace(text, "）", ")", -1)
	texts := strings.Split(text, " ")
	var resultTexts []string
	for _, v := range texts {
		v = strings.Replace(v, " ", "", -1)
		if len(v) < 3 {
			continue
		}
		resultTexts = append(resultTexts, v)
	}
	return resultTexts
}

func getTDK(doc *goquery.Document) (title string, desc string, keywords []string) {
	titleELemnt := doc.Find("title")
	title = titleELemnt.Text()
	desc = doc.Find("meta[name=description]").AttrOr("content", "")
	keywordStr := doc.Find("meta[name=keywords]").AttrOr("content", "")
	keywords = splitKeywords(keywordStr)
	return
}

func splitKeywords(keywordsStr string) []string {
	keywordsStr = strings.Replace(keywordsStr, ",", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "-", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "，", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "、", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "_", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, " ", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "\t", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, ";", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "；", "|", -1)
	keywordsStr = strings.Replace(keywordsStr, "\n", "", -1)
	keywordsStr = strings.Replace(keywordsStr, "“", "", -1)
	keywordsStr = strings.Replace(keywordsStr, "”", "", -1)
	return removeRepeatedElement(strings.Split(keywordsStr, "|"))

}

//Mozilla/5.0 (Linux;u;Android 4.2.2;zh-cn;) AppleWebKit/534.46 (KHTML,like Gecko) Version/5.1 Mobile Safari/10600.6.3 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html）
func getPagePCHtml(pageUrl string) (string, error) {
	wec, err := http_util.GetWebConFromUrlWithAllArgs(pageUrl,
		map[string]string{"User-Agent": "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"},
		"GET", nil, time.Second*20)
	if err != nil {
		wec, err = http_util.GetWebConFromUrlWithAllArgs(pageUrl,
			map[string]string{"User-Agent": "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"},
			"GET", nil, time.Second*20)
	}
	if err != nil {
		return "", err
	}
	return wec, nil

}
func clearHoleText(wholeText string) string {
	for _, l := range letters {
		wholeText = strings.Replace(wholeText, l, "", -1)
	}
	for _, n := range nums {
		wholeText = strings.Replace(wholeText, n, "", -1)
	}
	for _, p := range punctuations {
		wholeText = strings.Replace(wholeText, p, "|", -1)

	}

	wholeText = strings.Replace(wholeText, " ", "", -1)
	wholeText = strings.Replace(wholeText, " ", "", -1)
	wholeText = strings.Replace(wholeText, "	", "", -1)
	wholeText = strings.Replace(wholeText, "\r", "", -1)
	wholeText = strings.Replace(wholeText, "\n", "", -1)
	wholeText = strings.Replace(wholeText, "（", "(", -1)
	wholeText = strings.Replace(wholeText, "）", ")", -1)
	//wholeText = strings.Replace(wholeText,"名称",punctuations[rand.Intn(10)]+"名称"+punctuations[rand.Intn(10)],-1)
	//wholeText = strings.Replace(wholeText,"地址",punctuations[rand.Intn(10)]+"地址"+punctuations[rand.Intn(10)],-1)
	for _, eb := range emptybrackets {
		for strings.Contains(wholeText, eb) {
			wholeText = strings.Replace(wholeText, eb, "", -1)
		}
	}

	return wholeText
}
func TestPageType(pageUrl string, pageLinkText string, title string) string {
	pageUrlR, _ := url.Parse(pageUrl)
	if pageUrlR != nil {
		host := pageUrlR.Host
		pageUrl = strings.Replace(pageUrl, "http://", "", -1)
		pageUrl = strings.Replace(pageUrl, "https://", "", -1)
		pageUrl = strings.Replace(pageUrl, "/", "", -1)
		if pageUrl == host {
			return HomePageTypes
		}
	}

	if strings.Contains(pageLinkText, "联系我们") || strings.Contains(title, "联系我们") || strings.Contains(strings.ToLower(pageUrl), "contact") ||
		strings.Contains(strings.ToLower(pageUrl), "lianxi") || strings.Contains(strings.ToLower(pageUrl), "lxwm") {
		return ContactPageTypes
	}
	if strings.Contains(title, "关于") || strings.Contains(pageLinkText, "关于") ||
		strings.Contains(strings.ToLower(pageUrl), "introduce") ||
		strings.Contains(strings.ToLower(pageUrl), "gywm") ||
		strings.Contains(strings.ToLower(pageUrl), "guanyuwomen") ||
		strings.Contains(strings.ToLower(pageUrl), "aboutus") ||
		strings.Contains(strings.ToLower(pageUrl), "about_us") ||
		strings.Contains(strings.ToLower(pageUrl), "about") {
		return AboutPageTypes
	}
	return ""
}
