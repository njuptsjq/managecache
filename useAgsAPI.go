package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type CacheParams struct {
	Token           string `json:"token"`
	ServiceLocation string `json:"service_url"`
	Scales          string `json:"levels"`
	Instances       int    `json:"thread_count"`
	UpdateMode      string `json:"update_mode"`
	AOI             string `json:"area_of_interest"`
}

type TokenGen struct {
	F          string `json:"f"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Clinet     string `json:"clinet"`
	Expiration int    `json:"expiration"`
}

type TokenInfo struct {
	Token   string `json:"token"`
	Expires int    `json:"expires"`
}

type aoiInputer interface {
	inputAOI() (string, error)
}

type exampleAOI struct {
	filename string
	aoijson  string
}

func getToken() (TokenInfo, error) {
	conf, err := readConfig()
	getTokenAPI := conf.AgsURL + "tokens" + "/generateToken"
	tokengen := TokenGen{
		UserName:   conf.AgsName,
		Password:   conf.AgsPassword,
		F:          "json",
		Clinet:     "requestip",
		Expiration: 43200,
	}
	tokeninfo := TokenInfo{}
	if err != nil {
		fmt.Println(err)
		return tokeninfo, err
	}
	resp, err := http.Get(
		getTokenAPI + "?" + "f=json&" +
			"username=" + tokengen.UserName + "&" +
			"password=" + tokengen.Password + "&" +
			"clinet=" + tokengen.Clinet + "&" +
			"expiration=" + "60")
	if err != nil {
		log.Fatal(err)
		return tokeninfo, err
	}
	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return tokeninfo, err
	}
	if err := json.Unmarshal(result, &tokeninfo); err != nil {
		log.Fatal(err)
		return tokeninfo, err
	}
	return tokeninfo, nil
}

func getServiceLocation(serviceURL string) string {
	locationInfo := strings.Split(serviceURL, "/")
	if len(locationInfo) < 6 {
		return ""
	}
	locationInfo = locationInfo[6 : len(locationInfo)-1]
	location := ""
	for i := 0; i < len(locationInfo); i++ {
		location = location + "/" + locationInfo[i]
	}
	location = location[1:]
	return location + ":" + "MapServer"
}

// func getScales([]]string) (string, error) {
// 	requestURL := serviceURL
// 	resp, err := http.Get(requestURL)
// 	if err != nil {
// 		return "", nil
// 	}
// 	result, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", nil
// 	}
// 	return string(result), nil

// }

func getparam(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("参数解析错误")
	}
	valus := r.Form
	if _, URLok := valus["serviceURL"]; !URLok {
		log.Println("url参数解析错误")
	}
	if _, aoiok := valus["aoi"]; !aoiok {
		log.Println("aoi参数解析错误")
	}
	if _, scalesok := valus["scales"]; !scalesok {
		log.Println("scales参数解析错误")
	}
	log.Println(valus["serviceURL"], valus["scales"], valus["aoi"])

	tokenInfo, err := getToken()
	if err != nil {
		log.Println("获取token错误")
	}
	// 测试用
	eg := &exampleAOI{
		filename: "aoi.json",
		aoijson:  "",
	}
	eg.inputAOI()
	if err != nil {
		log.Println("获取result错误")
	}

	cacheManage := &CacheParams{
		Token:           tokenInfo.Token,
		ServiceLocation: getServiceLocation(valus["serviceURL"][0]),
		Scales:          valus["scales"][0],
		Instances:       2,
		UpdateMode:      "RECREATE_ALL_TILES",
		AOI:             eg.aoijson,
	}
}

func (eg *exampleAOI) inputAOI() (string, error) {
	aoifile, err := ioutil.ReadFile("aoi.json")
	if err != nil {
		return "", nil
	}
	eg.aoijson = string(aoifile)
	return string(aoifile), nil
}
