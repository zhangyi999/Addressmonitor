package csv

import (
	"io/ioutil"
	"os"
	"strconv"
	//    "reflect"
)

type Transaction struct {
	from  string
	to    string
	value string
	txid  string
	time  string
}

func PresCsv(path string) []Transaction {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	var teade []Transaction
	var val [7]string
	col := 0
	for _, s := range content {
		v := string(s)
		if v == "\"" {
			col++
			continue
		}
		if col%2 != 0 {
			cols := (col - 1) / 2
			val[cols] += v
		} else {
			cols := col / 2
			if cols == 7 {
				col = 0
				cols = 0
				t := Transaction{
					from:  val[4],
					to:    val[5],
					value: val[6],
					txid:  val[0],
					time:  val[2],
				}
				teade = append(teade, t)
			}
			val[cols] = ""
		}
	}
	return teade
}

func GetAddress(addrss string, tradeList []Transaction) ([]float64, []float64) {
	var ds []float64
	var count []float64
	for k, d := range tradeList {
		v, _ := strconv.ParseFloat(d.value, 64)
		var i float64
		if d.to == addrss {
			i = 1
		}
		if d.from == addrss {
			i = -1
		}
		ds = append(ds, i*v)
		if k > 0 {
			count = append(count, count[k-1]+i*v)
		} else {
			count = append(count, i*v)
		}
	}
	return ds, count
}
