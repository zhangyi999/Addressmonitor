package main

import (
	"fmt"

	"./csv"
)

func getPath(url string) string {
	return "../Date/" + url + ".csv"
}

func main() {
	tr := csv.PresCsv(getPath("ZZEX/0x3d475e9edef129acaacfb1cf282b842b723772f0"))

	fmt.Println(tr)
	// fmt.Println(tr[0].from)
}
