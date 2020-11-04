package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

//func main() {
//	//fmt.Print("hykj")
//
//
//
//	//nonce := "ssasdkkhhs1233sr"
//
//
//	key := "192006250b4c09247ec02edce69f6a2d"
//	besign := "appid=wxd930ea5d5a258f4f&body=test&device_info=1000&mch_id=10000100&nonce=ibuaiVcKdpRxkhJA&key=192006250b4c09247ec02edce69f6a2d"
////	//
//	sign := HmacSHA256(key,besign)
//	fmt.Println(strings.ToUpper(sign))
//	//uploadImg()
//	//printll()
//	Deci()
////fmt.Print(	GetNumOfDays(time.Now().AddDate(1,0,0),time.Now().AddDate(1,0,0)))
////	calcul()
////	fmt.Println()
//}
//
func HmacSHA256(key, s string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}
//
//
////func ComputeHmacSha256(message string, secret string) string {
////	key := []byte(secret)
////	h := hmac.New(sha256.New, key)
////	h.Write([]byte(message))
////	//	fmt.Println(h.Sum(nil))
////	sha := hex.EncodeToString(h.Sum(nil))
////	//	fmt.Println(sha)
////
////	//	hex.EncodeToString(h.Sum(nil))
////	return base64.StdEncoding.EncodeToString([]byte(sha))
////}
//
//func uploadImg()  {
//	url := "http://47.106.229.244:9466/api/base/uploadImage"
//	bFile,_ :=ioutil.ReadFile("./132401681.jpg")
//	req, _ := http.NewRequest("POST", url, strings.NewReader(string(bFile)))
//	req.Header.Add("authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDI4MTUxNjQsImlhdCI6MTYwMjcyODc2NCwibmJmIjoxNjAyNzI4NzY0LCJ1c2VySWQiOiIyMjgwMDAwNzUyMjM3NzcyODAwIiwic3lzVHlwZSI6MH0.L-3yO8GCihWijhml3KXiyjPyYF_BH9ERdxJ-ewA-Uxc")
//	res, _ := http.DefaultClient.Do(req)
//	defer res.Body.Close()
//	b, _ := ioutil.ReadAll(res.Body)
//	fmt.Println(res)
//	fmt.Println(string(b))
//}
//
//func printll()  {
//	nowDate := time.Now().UTC()
//	fmt.Println(nowDate)
//}



func Deci (){

	dd ,_ := decimal.NewFromString("100.0000")

	fmt.Print(dd.Floor())
}


func calcul (){
	x := 1342
	ints := make([]int,0)
	i := GetTimeArr("2020-10-18","2021-10-18")
	for index := int(i) ; index > 0 ; index --  {
	ints = append(ints,x / index)
	x = x- x/index
	}
	fmt.Println(len(ints))
}

func GetNumOfDays(start, end time.Time) int64 {
	startTime := start.Unix()
	endTime := end.Unix()
	// 求相差天数
	date := (endTime - startTime) / 86400
	return date
}

func GetTimeArr(start, end string) int64{
	timeLayout  := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	// 转成时间戳
	startUnix, _ := time.ParseInLocation(timeLayout,  start,  loc)
	endUnix, _ := time.ParseInLocation(timeLayout,  end,  loc)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	// 求相差天数
	date :=	(endTime - startTime) / 86400
	return date
}