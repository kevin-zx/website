package websitetool

import (
	"fmt"
	"github.com/kevin-zx/websitetool/companynametool"
	"github.com/kevin-zx/websitetool/extract"
	"reflect"
	"testing"
)

func TestGetWebSiteByHost(t *testing.T) {
	ss := []string{
		//"www.credithc.com",
		////"www.zgsjsws.com",
		////"www.xjbktx.com",
		////"www.yixin.com",
		////"www.credithc.com",
		////"www.zrfzsy.com",
		////"www.wangzhou.cn",
		////"www.minxinjituan.com",
		////"www.credithc.com",
		//"www.creditease.cn",
		//"www.pwccn.com",
		//"www.zhcpa.cn",
		//"www.kuaikuaidai.com",
		//"www.credithc.com",
		//"www.hebeitaihang.com",
		//"www.credithc.com",
		//"www.credithc.com",
		//"www.jieyuechina.com",
		//"www.credithc.com",
		//"www.kjcity.com",
		//"www.credithc.com",
		//"www.baoying.com",
		//"www.kjcity.com",
		//"www.shengshihuihai.com.cn",
		//"www.jxzxj.com",
		//"www.jinghuacpa.com",
		//"www.shengshihuihai.com.cn",
		//"www.shszaf.com",
		//"www.hebeitaihang.com",
		//"www.hanbangrd.com",
		//"www.jcccoffice.com",
		//"www.credithc.com",
		//"www.dashunl.com",
		//"www.minghongsws.com",
		//"www.zkbc.net",
		//"www.sunshinewh.com",
		//"www.yinde-china.com",
		//"www.51qmb.com",
		//"www.fabpo.com",
		//"www.d2capital.cn",
		//"www.credithc.com",
		//"www.deloitte.com",
		//"www.hengqijy.com",
		//"www.hengyirong.com",
		//"www.shengshihuihai.com.cn",
		//"www.axzytz.com",
		//"www.ycaxkj.com",
		//"www.credithc.com",
		//"www.vcredit.com",
		//"www.shengshihuihai.com.cn",
		//"www.rthxchina.com",
		//"www.credithc.com",
		//"www.creditease.cn",
		//"www.credithc.com.cn",
		//"www.credithc.com",
		//"www.credithc.com",
		//"www.11pro.club",
		//"www.qmx12366.cn",
		//"www.credithc.com",
		//"www.shengshihuihai.com.cn",
		//"www.jieyuechina.com",
		//"www.changhw.com",
		//"www.credithc.com",
		//"www.haihuawealth.com",
		//"www.juejindongcheng.com",
		//"www.credithc.com",
		"www.szwlskj.com",
	}
	for _, s := range ss {
		ws, err := GetWebSiteByHost(s)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		fmt.Printf("%s, %s\n", s, ws.CompanyName)
		for _, p := range ws.Pages {
			for _, tel := range p.TelPhones {
				fmt.Println(tel.TelNum)
			}
			for _, mphone := range p.MobilePhoneNums {
				fmt.Println(mphone.PhoneNum)
			}
		}
	}
}

func TestGetWebSiteByCompanyName(t *testing.T) {
	cns := []string{
		"苏州网联盛科技有限公司",
	}
	for _, cn := range cns {
		wss, err := GetWebSiteByCompanyName(cn)
		cnn := companynametool.ClearCompanyName(cn)
		if err != nil {
			panic(err)
		}
		for _, ws := range wss {

			//if strings.Contains(ws.Pages[0].Text, cn) {
			fmt.Printf("%s------------------------%s------------------------%s------------------------%s------------------------%s-------------------------%s\n", cn, cnn, ws.CompanyName, companynametool.ClearCompanyName(ws.CompanyName), ws.SiteUrl, ws.Pages[0].Title)
			//}
		}
	}
}

func TestGetWebSiteByHost1(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name    string
		args    args
		want    *extract.Website
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWebSiteByHost(tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWebSiteByHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWebSiteByHost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTopCompany(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTopCompany(tt.args.m); got != tt.want {
				t.Errorf("getTopCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}
