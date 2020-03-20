package main

import (
	// "fmt"
	"fmt"
	"io/ioutil"
	"net/http"
	// "address/csv"
)

var addressList = map[string]string{
	"BIKI": "0x6eff3372fa352b239bb24ff91b423a572347000d",
	"ZZEX": "0x86e793e413f519b450315fcc4b618eb25a3a54a4",
}

func getPath(exchange string) string {
	//路径和执行命令时的窗口有关
	return "../Date/" + exchange + "/" + addressList[exchange] + ".csv"
}

func get() {
	//get请求
	//http.Get的参数必须是带http://协议头的完整url,不然请求结果为空
	resp, _ := http.Get("http://api-cn.etherscan.com/api?module=account&action=tokentx&page=1&offset=10000&sort=desc&contractaddress=0xdac17f958d2ee523a2206206994597c13d831ec7&address=0x86e793e413f519b450315fcc4b618eb25a3a54a4&apikey=HVV7H6G99H6BSCVT5EHIJVDYKHR5J28H59")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	// fmt.Printf("Get request result: %s\n", string(body))
}

func main() {
	get()
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(path.resolve("~/sample.sh"))
	// os.Create("./A.txt")

	// getPath("ZZEX/0x3d475e9edef129acaacfb1cf282b842b723772f0")
	// dex := "ZZEX"
	// tr := csv.PresCsv(getPath(dex))

	// fmt.Println(tr)
	// var adds map[string]*CountTradeAll
	// adds := csv.CountAddress("0x86e793e413f519b450315fcc4b618eb25a3a54a4", tr)

	// fmt.Println(adds["0xa964b1bf3e496ad214ff0fb903f89ccf3b9428a0"])

	// var m = new(csv.MAddress)
	// m.MeasureAddress("0x86e793e413f519b450315fcc4b618eb25a3a54a4", tr) //addressList[dex]
	// fmt.Println("intoValue", m.IntoValue)
	// fmt.Println("intoNum", m.IntoNum)
	// fmt.Println("intoAddressNum", m.IntoAddressNum)
	// fmt.Println("outoValue", m.OutoValue)
	// fmt.Println("outoAddressNum", m.OutoAddressNum)
	// fmt.Println("outoNum", m.OutoNum)
	// fmt.Println("balance", m.Balance)

	// fmt.Println("intoValue", csv.GetSpecialAddress(m.OutoList))
}
