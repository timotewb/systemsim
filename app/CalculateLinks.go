package app

import (
	"fmt"
	"strconv"
	"strings"
	m "systemsim/models"
)

func CalculateLinks(data m.DataType)(error){

	var e string
	var v float64
	for n := 0; n < len(data.Nodes); n++{

		// if the node type is link
		if data.Nodes[n].Type == "link"{
			e = data.Nodes[n].Equation
			fmt.Println(e)

			// loop over equation parts and update equation
			for p := 0; p < len(data.Nodes[n].Parts); p++{
				fmt.Println(data.Nodes[n].Parts[p].Attribute_ID)
				fmt.Println(data.Nodes[n].Parts[p].Node_ID)

				v, err := GetNodeAttributeValue(data, data.Nodes[n].Parts[p].Node_ID, data.Nodes[n].Parts[p].Attribute_ID)
				if (err == nil){
					e = strings.Replace(e, "|"+strconv.Itoa(data.Nodes[n].Parts[p].ID)+"|", strconv.FormatFloat(v, 'g', 5, 64), 1)
				} else{
					return err
				}
			}
		} else {
			// if there are no link types which use the 'e' and 'v' variables
			_ = e
			_ = v
		}
	}
	return nil
}