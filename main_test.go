package main

import (
	"testing"
	"fmt"
)

func TestGetServiceLocation(t *testing.T) {
	inputURL := "http://arcserver1.htgis.com:6080/arcgis/rest/services/LiShuiCacheTest/ls_tile_management/MapServer"
	output := "LiShuiCacheTest/ls_tile_management:MapServer"
	result := getServiceLocation(inputURL)
	if output != result {
		t.Fatal("解析不成功", result)
	}
}

func TestReadConfig(t *testing.T) {
	conf, _ := readConfig()
	confResult := AgsConfig{
		AgsName: "siteadmin",
	}
	if conf.AgsName != confResult.AgsName {
		t.Fatal("配置读取失败", conf.AgsName)
	}
}

func TestGetToken(t *testing.T) {
	expectedExpire := 43200
	tokeninfo, err := getToken()
	if err != nil {
		t.Fatal(err)
	}
	if tokeninfo.Expires < expectedExpire {
		t.Fatal("token获取失败")
	}
}

func TestInputAOI(t *testing.T){
	eg:=&exampleAOI{
		filename:"aoi.json",
		aoijson:"",
	}
	result,err:=eg.inputAOI()
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Printf("%v",result)
}