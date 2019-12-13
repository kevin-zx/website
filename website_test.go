package websitetool

import (
	"fmt"
	"log"
	"testing"
)

func TestGetWebSiteByHost(t *testing.T) {
	ws, err := GetWebSiteByHost("www.020jinyou.com")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Printf("%v\n", ws)
}
