package main

import (
	"fmt"
	"log"

	"netgov/spider/spider"
)

// Send order

var Dealorder []interface{}

func Manager(cookie string, url1 string, url2 string) bool {
	result := spider.GetList(cookie, url1)
	for _, i := range result.Rows {
		result := spider.GetOrderDetails(i.WfID, cookie, url2)
		result = spider.DealOrder(result)
		b := spider.SendOrder(result)
		if b == true {
			log.Printf("1")
		}
	}
	return true
}

func main() {
	cookie := "JSESSIONID=6B6706D22EFB3A50FA6763F4D8FAE4FD"
	// db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	log.Print("Please input Cookit:")

	fmt.Scanln(&cookie)
	url1 := "http://171.221.172.75:8890/mTelWF/openManagerListBL.ajx?act=6"
	url2 := "http://171.221.172.75:8890/mTelWF/cmnFile/showWorkForm.do?wfid="
	Manager(cookie, url1, url2)

	// url1 = "http://171.221.172.75:8890/mMailWF/mailWFSendBackListBL.ajx?lCol=300&wc=-30&urlStt=10,11&act=7"
	// url2 = "http://171.221.172.75:8890/mMailWF/mailWFSendBackListBL.ajx?lCol=300&wc=-30&urlStt=10,11&act=7"
	// Manager(cookie, url1, url2)
}
