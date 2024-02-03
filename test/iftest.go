package main

import (
	"fmt"
	"strings"
)

// assetDB
func main() {

	var connect string = ""
	var name = "upload"

	// table명 받아옴
	// table명에 따라 data 받아오는 값 달라짐
	// data를 db에 저장
	fmt.Println("AssetData : " + name)
	trim_name := strings.Trim(name, " ")

	// < 다른 함수에서 호출하는 경우 오류 발생 >
	// trim_name = strings.Trim(name, " ") // 이렇게 하는 경우 오류
	// 	fmt.Println(trim_name + "1connect")  // 이거는 호출만 됨(1connect만 출력됨) --> 이후 if문도 출력 안되는 문제
	// 	fmt.Println("1connect " + trim_name) // 1connect upload 잘 호출 됨

	fmt.Println("trim_name : " + trim_name)

	// upload table에 받아온 데이터 구분해서 저장
	if trim_name == "upload" {

		connect = trim_name + "connect"
		fmt.Println(connect)

	} else if trim_name == "file" {

		// file table에 받아온 데이터 구분해서 저장
		connect = trim_name + "connect"
		fmt.Println(connect)

	}

	connect = trim_name + " connect"
	fmt.Println(connect)

	//test message
	// message := fmt.Sprintf("%s success", connect)
	// return message
}
