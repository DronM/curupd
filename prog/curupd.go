package main

import(
	"encoding/json"
	"fmt"	
	
	"curupd"
)

func main() {
	rates, err := curupd.GetCurrencyRates()
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	
	b, err := json.Marshal(rates)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	fmt.Println(string(b))
}

