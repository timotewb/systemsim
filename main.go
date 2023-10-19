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
	a.Calc(data)

	fmt.Println("--- Pretty Print ---")
    //Pretty print josn with MarshalIndent
    dataJSON, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("%s\n", string(dataJSON))
}

// updating the equation with attribute values
// execute equation
// update model