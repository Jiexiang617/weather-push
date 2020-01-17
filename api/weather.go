package api

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"weather-push/model"
	"weather-push/util"
)

/// 根据城市名字，查询天气（模糊），若地址有多个重复，默认取第一个
func QueryWithCityName(cityName string) model.Weather {
	addressList := queryAddress(cityName)
	return queryWeather(addressList[0].CityCode)
}

/// 根据城市名字，查询匹配的城市结果
func queryAddress(keyWord string) [] model.Address {
	// 把url中的汉字编码
	keyWord = url.PathEscape(keyWord)
	url := fmt.Sprintf("http://toy1.weather.com.cn/search?cityname=%s", keyWord)
	res, err := http.Get(url)
	if err != nil {
		util.Log().Error("查询城市失败", err)
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		util.Log().Error("查询城市失败", err)
	}
	jsonStr := string(bytes)
	// 返回结果进行处理
	jsonStr = jsonStr[1:len(jsonStr) - 1]
	bytes = []byte(jsonStr)

	var addressDictList []map[string]string
	err = json.Unmarshal(bytes, &addressDictList)
	if err != nil {
		util.Log().Error("查询城市失败", err)
	}

	if len(addressDictList) == 0 {
		log.Fatal("没有查询到这个地址信息")
	}

	var addressList []model.Address
	for _, value := range addressDictList {
		addressList = append(addressList, addressStringToModel(value["ref"]))
	}

	return addressList
}

/// 根据城市id，查询天气
func queryWeather(cityCode string) model.Weather {
	if len(cityCode) > 9 { // 如果城市编码比较长，代表是乡镇，就只能调用天气预测接口
		return forecastWeather(cityCode)
	} else { // 如果是城市，直接调用天气接口
		return queryCityWeather(cityCode)
	}
}

func queryCityWeather(cityCode string) model.Weather {
	url := fmt.Sprintf("http://www.weather.com.cn/weather1d/%s.shtml", cityCode)

	res, err := http.Get(url)
	if err != nil {
		util.Log().Error("查询城市天气失败", err)
	}

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		util.Log().Error("查询城市天气失败", err)
	}

	detail := doc.Find("div.t")

	maxT := detail.Find("p.tem span").First().Text()
	minT := detail.Find("p.tem span").Last().Text()

	dayTime := detail.Find("p.wea").First().Text()
	night := detail.Find("p.wea").Last().Text()
	describe := fmt.Sprintf("白天：%s，夜间：%s", dayTime, night)

	return model.Weather{
		Info:    describe,
		MaxTemp: maxT,
		MinTemp: minT,
	}
}

func forecastWeather(cityCode string) model.Weather {
	url := fmt.Sprintf("http://forecast.weather.com.cn/town/weather1dn/%s.shtml", cityCode)

	res, err := http.Get(url)
	if err != nil {
		util.Log().Error("查询城市天气失败", err)
	}

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		util.Log().Error("查询城市天气失败", err)
	}

	detail := doc.Find("div.todayLeft")
	describe, _ := doc.Find("ul#weatherALL li").First().Find("i.weather").Attr("title")
	maxT := detail.Find("div.minMax").Find("div#maxTempDiv span").Text()
	minT := detail.Find("div.minMax").Find("div#minTempDiv span").Text()

	maxT = strings.Replace(maxT, "℃", "", 1)
	minT = strings.Replace(minT, "℃", "", 1)

	return model.Weather{
		Info:    describe,
		MaxTemp: maxT,
		MinTemp: minT,
	}
}

/// 将字符串转换为模型
/// 101271003008~sichuan~立石镇~lishizhen~泸县~luxian~0830~646100~sichuan~四川
func addressStringToModel(addressStr string) model.Address {
	dataInfo := strings.Split(addressStr, "~")
	if len(dataInfo) < 5 {
		log.Fatalln("返回的地址信息异常长度太短：", addressStr)
		panic("返回数据异常")
	}
	return model.Address{
		CityCode: dataInfo[0],
		Province: dataInfo[len(dataInfo) - 1],
		County:   dataInfo[0],
	}
}