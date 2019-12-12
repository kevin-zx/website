package extract

type WebPage struct {
	PageUrl         string            `json:"page_url,omitempty"`
	Type            string            `json:"type,omitempty"`              //主页 联系我们 地图 资讯
	Title           string            `json:"title,omitempty"`             // 标题
	Keywords        []string          `json:"keywords,omitempty"`          // keywords
	Description     string            `json:"description,omitempty"`       // 描述
	TelPhones       []TelPhoneNum     `json:"tel_phones,omitempty"`        // 电话
	MobilePhoneNums []MobilePhoneNum  `json:"mobile_phone_nums,omitempty"` // 手机
	Text            string            `json:"-"`
	SplitText       []string          `json:"-"`
	AllLinks        map[string]string `json:"-"`
}

type Website struct {
	SiteUrl     string    `json:"site_url,omitempty"`     // 主页链接
	Pages       []WebPage `json:"pages"`                  // 页面
	CompanyName string    `json:"company_name,omitempty"` //公司名称
}

type TelPhoneNum struct {
	TelNum   string `json:"tel_num,omitempty"`
	AreaCode string `json:"area_code,omitempty"`
	City     string `json:"city,omitempty"`
}

type MobilePhoneNum struct {
	PhoneNum                  string `json:"phone_num,omitempty"`
	TelecommunicationOperator string `json:"telecommunication_operator,omitempty"`
	City                      string `json:"city,omitempty"`
	ZipCode                   string `json:"zip_code,omitempty"` //邮政编码
}
