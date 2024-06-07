package common

import (
	"encoding/json"
	"fmt"
	"os"
)

// all addresses in test
type Address struct {
	Market string `json:"market"`
	Access string `json:"access"`
	Credit string `json:"credit"`

	Token    string `json:"token"`
	Registry string `json:"registry"`
	Swap     string `json:"swap"`
	Pledge   string `json:"pledge"`
}

// save address struct into pathName
func Save(a Address, pathName string) {
	b, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// save data into json file
	err = os.WriteFile(pathName, b, 0644)
	if err != nil {
		fmt.Println("Error writing JSON:", err)
		return
	}

}

// load person from pathName
func Load(pathName string) Address {
	// load file
	b, err := os.ReadFile(pathName)
	if err != nil {
		fmt.Println("Error reading JSON:", err)
		return Address{}
	}

	a := Address{}
	err = json.Unmarshal(b, &a)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return Address{}
	}

	return a
}
