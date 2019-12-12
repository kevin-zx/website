package regexInfopaser

import (
	"regexp"
	"strings"
)

//const phoneRegexStr = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4(?:[14]0\d{3}|[68]\d{4}|[579]\d{2}))\d{6}$`
const phoneRegexStr = `(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4(?:[14]0\d{3}|[68]\d{4}|[579]\d{2}))\d{6}`

var telStart = []string{"+86", "86", "010", "021", "022", "023", "0311", "0313", "0314", "0335", "0315", "0316", "0312", "0317", "0318", "0319", "0310", "0571", "0572", "0573", "0580", "0574", "0575", "0570", "0579", "0576", "0577", "0578", "024", "024", "024", "024", "0421", "0418", "0419", "0412", "0415", "0411", "0417", "0427", "0416", "0429", "0731", "0731", "0731", "0744", "0736", "0737", "0730", "0734", "0735", "0746", "0739", "0745", "0738", "0743", "025", "0516", "0518", "0527", "0517", "0515", "0514", "0523", "0513", "0511", "0519", "0510", "0512", "0471", "0472", "0473", "0476", "0475", "0470", "0477", "0474", "0478", "0482", "0479", "0483", "0791", "0792", "0798", "0701", "0790", "0799", "0797", "0793", "0794", "0795", "0796", "0351", "0352", "0349", "0353", "0355", "0356", "0350", "0354", "0357", "0359", "0358", "0931", "0935", "0935", "0943", "0938", "0937", "0937", "0936", "0934", "0933", "0932", "0939", "0930", "0941", "0531", "0635", "0534", "0546", "0533", "0536", "0535", "0631", "0532", "0633", "0539", "0632", "0537", "0538", "0634", "0543", "0530", "0451", "0452", "0456", "0459", "0458", "0468", "0454", "0469", "0464", "0467", "0453", "0455", "0457", "0591", "0599", "0598", "0594", "0595", "0592", "0596", "0597", "0593", "020", "0763", "0751", "0762", "0753", "0768", "0754", "0663", "0660", "0752", "0769", "0755", "0756", "0760", "0750", "0757", "0758", "0766", "0662", "0668", "0759", "028", "028", "028", "0839", "0816", "0838", "0817", "0826", "0825", "0832", "0833", "0813", "0830", "0831", "0812", "0827", "0818", "0835", "0837", "0836", "0834", "027", "0719", "0719", "0710", "0724", "0712", "0713", "0711", "0714", "0715", "0716", "0717", "0722", "0728", "0728", "0728", "0718", "0371", "0371", "0398", "0379", "0391", "0391", "0373", "0392", "0372", "0393", "0370", "0374", "0395", "0375", "0377", "0376", "0394", "0396", "0871", "0874", "0877", "0875", "0870", "0888", "0879", "0883", "0692", "0886", "0887", "0872", "0878", "0873", "0876", "0691", "0551", "0557", "0561", "0558", "0558", "0552", "0554", "0550", "0555", "0553", "0562", "0556", "0559", "0564", "0566", "0563", "0951", "0952", "0953", "0954", "0955", "0431", "0436", "0438", "0432", "0434", "0437", "0435", "0439", "0433", "0771", "0771", "0773", "0772", "0772", "0774", "0774", "0775", "0775", "0777", "0779", "0770", "0776", "0778", "0851", "0851", "0851", "0858", "0857", "0856", "0855", "0854", "0859", "029", "029", "0911", "0919", "0913", "0917", "0916", "0912", "0915", "0914", "0971", "0972", "0970", "0974", "0973", "0975", "0976", "0977", "0979", "0898", "0891", "0896", "0895", "0894", "0893", "0892", "0897", "0991", "0990", "0990", "0992", "0992", "0992", "0993", "0993", "0997", "0997", "0997", "0998", "0998", "0906", "0906", "0906", "0903", "0995", "0902", "0908", "0909", "0994", "0994", "0996", "0999", "0901"}

func MatchPhone(text string) []string {
	text = strings.Replace(text, "-", "", -1)
	reg, _ := regexp.Compile(phoneRegexStr)
	//b:=reg.MatchString(text)
	//fmt.Println(b)
	nums := reg.FindAllString(text, -1)
	pts := MatchtNums(text)
	r := []string{}
	for _, num := range nums {
		m := false
		for _, pt := range pts {
			if len(pt) > 11 && strings.Contains(pt, num) {
				m = true
				break
			}

		}
		if !m {
			r = append(r, num)
		}
	}
	return removeRepeatedElement(r)
}

const numsRegexStr = `\d+`

func MatchtNums(text string) []string {

	reg, _ := regexp.Compile(numsRegexStr)
	//b:=reg.MatchString(text)
	//fmt.Println(b)
	nums := reg.FindAllString(text, -1)
	return removeRepeatedElement(nums)
}

const creditCodeRegexStr = `[^_IOZSVa-z\W]{2}\d{6}[^_IOZSVa-z\W]{10}`

func MatchCreditCode(text string) []string {
	reg, _ := regexp.Compile(creditCodeRegexStr)

	//b:=reg.MatchString(text)
	//fmt.Println(b)
	ccs := reg.FindAllString(text, -1)
	return removeRepeatedElement(ccs)

}

//MA5FQFB6-1
const organizationNoRegexStr = `[A-Za-z0-9]{8}-[A-Za-z0-9]`

func MatchOrganizationNO(text string) []string {
	reg, _ := regexp.Compile(organizationNoRegexStr)

	//b:=reg.MatchString(text)
	//fmt.Println(b)
	oNs := reg.FindAllString(text, -1)
	return removeRepeatedElement(oNs)
}

//MA5FQFB6-1
const telephoneRegexStr = `(\(\d{3,4}\)|\d{3,4}-|\s)?\d{8}`

func MatchTelephone(text string) []string {
	reg, _ := regexp.Compile(telephoneRegexStr)

	//b:=reg.MatchString(text)
	//fmt.Println(b)
	ttls := []string{}
	tls := reg.FindAllString(text, -1)
	pts := MatchtNums(text)
	for _, tl := range tls {
		for _, ts := range telStart {
			if strings.HasPrefix(tl, ts) {
				m := false
				for _, pt := range pts {
					if len(pt) > 8 && strings.Contains(pt, ts) {
						m = true
						break
					}

				}
				if !m {
					ttls = append(ttls, tl)
					break
				}

			}
		}

	}

	return removeRepeatedElement(ttls)
}

// email
const emailRegexStr = `[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+`

func MatchEmail(text string) []string {
	reg, _ := regexp.Compile(emailRegexStr)

	//b:=reg.MatchString(text)
	//fmt.Println(b)
	oNs := reg.FindAllString(text, -1)
	return removeRepeatedElement(oNs)
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

//　/[1-9A-GY]{1}[1239]{1}[1-5]{1}[0-9]{5}[0-9A-Z]{10}$|[1-9A-GY]{1}[1239]{1}[1-5]{1}[0-9]{5}[0-9A-Z]{10}-[0-9]{2}$/
// 企业13位工商注册号
const icpre = "[\u4e00-\u9fa5]ICP备\\d{8}号-\\d{1,3}"

func MatchICPNO(text string) []string {
	reg, err := regexp.Compile(icpre)
	if err != nil {
		panic(err)
	}
	//b:=reg.MatchString(text)
	//fmt.Println(b)
	oNs := reg.FindAllString(text, -1)
	return removeRepeatedElement(oNs)
}
