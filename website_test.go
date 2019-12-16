package websitetool

import (
	"fmt"
	"testing"
)

func TestGetWebSiteByHost(t *testing.T) {
	ss := []string{
		"www.credithc.com",
		//"www.zgsjsws.com",
		//"www.xjbktx.com",
		//"www.yixin.com",
		//"www.credithc.com",
		//"www.zrfzsy.com",
		//"www.wangzhou.cn",
		//"www.minxinjituan.com",
		//"www.credithc.com",
		"www.creditease.cn",
		"www.pwccn.com",
		"www.zhcpa.cn",
		"www.kuaikuaidai.com",
		"www.credithc.com",
		"www.hebeitaihang.com",
		"www.credithc.com",
		"www.credithc.com",
		"www.jieyuechina.com",
		"www.credithc.com",
		"www.kjcity.com",
		"www.credithc.com",
		"www.baoying.com",
		"www.kjcity.com",
		"www.shengshihuihai.com.cn",
		"www.jxzxj.com",
		"www.jinghuacpa.com",
		"www.shengshihuihai.com.cn",
		"www.shszaf.com",
		"www.hebeitaihang.com",
		"www.hanbangrd.com",
		"www.jcccoffice.com",
		"www.credithc.com",
		"www.dashunl.com",
		"www.minghongsws.com",
		"www.zkbc.net",
		"www.sunshinewh.com",
		"www.yinde-china.com",
		"www.51qmb.com",
		"www.fabpo.com",
		"www.d2capital.cn",
		"www.credithc.com",
		"www.deloitte.com",
		"www.hengqijy.com",
		"www.hengyirong.com",
		"www.shengshihuihai.com.cn",
		"www.axzytz.com",
		"www.ycaxkj.com",
		"www.credithc.com",
		"www.vcredit.com",
		"www.shengshihuihai.com.cn",
		"www.rthxchina.com",
		"www.credithc.com",
		"www.creditease.cn",
		"www.credithc.com.cn",
		"www.credithc.com",
		"www.credithc.com",
		"www.11pro.club",
		"www.qmx12366.cn",
		"www.credithc.com",
		"www.shengshihuihai.com.cn",
		"www.jieyuechina.com",
		"www.changhw.com",
		"www.credithc.com",
		"www.haihuawealth.com",
		"www.juejindongcheng.com",
		"www.credithc.com",
		"www.hebeitaihang.com"}
	for _, s := range ss {
		ws, err := GetWebSiteByHost(s)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		fmt.Printf("%s, %s\n", s, ws.CompanyName)
	}
}
