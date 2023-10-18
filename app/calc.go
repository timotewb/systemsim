package app

import (
	"fmt"
	m "systemsim/models"
)

func Calc(data m.DataType){

	var e string
	for n := 0; n < len(data.Nodes); n++{

		// if the node type is link
		if data.Nodes[n].Type == "link"{
			e = data.Nodes[n].Equation
			fmt.Println(e)

			// loop over equation parts and update equation
			for p := 0; p < len(data.Nodes[n].Parts); p++{
				fmt.Println(data.Nodes[n].Parts[p].Attribute_ID)
				fmt.Println(data.Nodes[n].Parts[p].Node_ID)
			}
		} else {
			// if there are no link types which use the 'e' variable
			_ = e
		}
	}
}