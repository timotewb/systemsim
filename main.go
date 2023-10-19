package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	a "systemsim/app"
	m "systemsim/models"
)

func main(){
	file, _ := os.ReadFile("db.json")
	data := m.DataType{}

	_ = json.Unmarshal([]byte(file), &data)

	fmt.Println("--- Calculate links ---")
	a.CalculateLinks(data)

	fmt.Println("--- Pretty Print ---")
    //Pretty print josn with MarshalIndent
    dataJSON, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("%s\n", string(dataJSON))
}

// calculate dependencies


// how do we auto calculate things e.g. link node which calculates raw materials processed based on resource
// available - how do we auto calculate the loss n revenue if the supply of raw materials is greater than the available resource.
