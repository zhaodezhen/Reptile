package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents,err :=fetcher.Fetch("hts://www.zhenai.com/zhenghun")	//用来测试
	contents,err :=ioutil.ReadFile("cityList_test_data.html")
	
	if err!=nil {
		panic(err)
	}
	//fmt.Printf("%s \n",contents)
	result := ParseCityList(contents)

	expectedUrls   :=[]string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCitys  :=[]string{
		"City:阿坝",
		"City:阿克苏",
		"City:阿拉善盟",
	}
	const resultSize  = 470

	if len(result.Requests) != resultSize {
 		t.Errorf("result should have %d" + "requests;but had %d",resultSize,len(result.Requests))
	}
	for i,url :=range expectedUrls{
		if result.Requests[i].Url != url{
			t.Errorf("expected url #%d: %s;but "+"was %s",i,url,result.Requests[i].Url)
		}
		
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d" + "item;but had %d",resultSize,len(result.Items))
	}
	for i,city :=range expectedCitys{
		if result.Items[i].(string) != city{
			t.Errorf("expected city #%d: %s;but "+"was %s",i,city,result.Items[i].(string))
		}

	}
}