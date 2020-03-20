package csv

import (
	"io/ioutil"
	"os"
	"sort"
	"strconv"

	"github.com/shopspring/decimal"
	//    "reflect"
)

type Transaction struct {
	from  string
	to    string
	value string
	txid  string
	time  string
}

type CountTrade struct {
	value      float64
	proportion float64
}

type CountTradeAll struct {
	to   CountTrade
	from CountTrade
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

// func GetAddress(addrss string, tradeList []Transaction) ([]decimal.Decimal, []decimal.Decimal) {
// 	var ds []decimal.Decimal
// 	var count []decimal.Decimal
// 	for k, d := range tradeList {
// 		// v, _ := strconv.ParseFloat(d.value, 64)
// 		v, _ := decimal.NewFromString(d.value)

// 		ds = append(ds, v)
// 		if k > 0 {
// 			count = append(count, count[k-1].Add(v))
// 		} else {
// 			count = append(count, v)
// 		}
// 	}
// 	return ds, count
// }

func addValue(List map[string]*CountTradeAll, key string) {

	if List[key] == nil {
		List[key] = &CountTradeAll{
			to:   CountTrade{value: 0, proportion: 0},
			from: CountTrade{value: 0, proportion: 0},
		}
	}

}

func CountAddress(addrss string, tradeList []Transaction) map[string]*CountTradeAll {
	// var List map[string]*CountTradeAll
	List := make(map[string]*CountTradeAll)
	var toTotal float64
	var fromTotal float64
	toTotal = 0.0
	fromTotal = 0.0
	for _, d := range tradeList {
		v, _ := strconv.ParseFloat(d.value, 64)
		if d.from == addrss {
			addValue(List, d.to)
			List[d.to].to.value += v
			toTotal += v
		} else if d.to == addrss {
			addValue(List, d.from)
			List[d.from].from.value += v
			fromTotal += v
		}
	}

	/*使用键输出地图值 */
	for country := range List {
		List[country].from.proportion = List[country].from.value / fromTotal * 100
		List[country].to.proportion = List[country].to.value / toTotal * 100
	}
	return List
}

// 1. 转入数量: intoValue
// 2. 转入笔数: intoNum
// 3. 转出数量: outoValue
// 4. 转出笔数: outoNum
// 5. 转入地址数量: intoAddressNum
// 6. 转出地址数量: outoAddressNum
// 7. 持有 token 数量: balance
type MAddress struct {
	IntoValue      float64
	IntoNum        int
	OutoValue      float64
	OutoNum        int
	IntoAddressNum int
	OutoAddressNum int
	Balance        float64
	IntoList       map[string]decimal.Decimal
	OutoList       map[string]decimal.Decimal
}

//要对golang map按照value进行排序，思路是直接不用map，用struct存放key和value，实现sort接口，就可以调用sort.Sort进行排序了。
// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value decimal.Decimal
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return (p[i].Value.Sub(p[j].Value)).IsPositive() }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]decimal.Decimal) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func (m *MAddress) MeasureAddress(mainAddrss string, tradeList []Transaction) {

	var intoValue decimal.Decimal
	var outoValue decimal.Decimal
	var iNum int
	var oNum int
	iAddressNum := make(map[string]decimal.Decimal)
	oAddressNum := make(map[string]decimal.Decimal)
	// intoValue = 0.0
	// outoValue = 0.0
	for _, d := range tradeList {
		v, _ := decimal.NewFromString(d.value) //strconv.ParseFloat(d.value, 64)
		if d.from == mainAddrss {
			outoValue = outoValue.Add(v)
			oNum++
			oAddressNum[d.to] = oAddressNum[d.to].Add(v)
		} else if d.to == mainAddrss {
			intoValue = intoValue.Add(v)
			iNum++
			iAddressNum[d.from] = iAddressNum[d.from].Add(v)
		}
	}

	m.IntoValue, _ = intoValue.Float64()
	m.IntoNum = iNum
	m.IntoAddressNum = len(iAddressNum)
	m.OutoValue, _ = outoValue.Float64()
	m.OutoNum = oNum
	m.OutoAddressNum = len(oAddressNum)
	m.Balance, _ = (intoValue.Sub(outoValue)).Float64()
	m.IntoList = iAddressNum
	m.OutoList = oAddressNum
}

func GetSpecialAddress(m map[string]decimal.Decimal) PairList {
	i := decimal.NewFromInt(int64(len(m))).Div(decimal.NewFromInt(40)).IntPart()
	// fmt.Println()
	return sortMapByValue(m)[:i]
}
