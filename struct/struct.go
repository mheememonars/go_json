package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Skill   []Skill `json:"skill"`
}

type Address struct {
	Province     string `json:"province"`
	ProvinceCode int    `json:"province_code"`
}

type Skill struct {
	Language string `json:"language"`
	Level    int    `json:"level"`
}

func main() {
	jsonData := `{  
		"name":"Mhee",
		"address":{  
		   "province":"กรุงเทพมหานคร",
		   "province_code":101
		},
		"skill":[  
		   {  
			  "language":"golang",
			  "level":80
		   },
		   {  
			  "language":"java",
			  "level":40
		   }
		]
	 }`
	var response Profile
	json.Unmarshal([]byte(jsonData), &response)
	fmt.Println(response)
}
