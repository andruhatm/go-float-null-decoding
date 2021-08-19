package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

var jsonStr = []byte(
	`[{"data":[[ 1628586000000, null ], [ 1628586900000, 0.0 ], [ 1628587800000, null ]] },
{"data":[[ 1728586000000, 0 ], [ 1728586900000, null ], [ 1728587800000, 0 ]] }]`)

func main() {

	var info [][]float64
	var data []map[string]interface{}

	if err := json.Unmarshal(jsonStr, &data); err != nil {
		panic(err)
	}

	for index, _ := range data {
		newData := data[index]["data"].([]interface{})
		fmt.Println(newData)
		for _, in := range newData {
			var arr []float64
			//fmt.Println(in)

			result, _ := json.Marshal(in)
			fmt.Printf("marshalled result %v\n", string(result))
			rs := strings.ReplaceAll(string(result), "null", "-1")

			json.Unmarshal([]byte(rs), &arr)
			info = append(info, arr)
		}
		fmt.Println(info)
		info = [][]float64{}
	}

}
