package app

import (
	"go/token"
	"go/types"
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
			// loop over equation parts and update equation
			for p := 0; p < len(data.Nodes[n].Parts); p++{

				// check if the part points to a nodes value
				if(data.Nodes[n].Parts[p].Is_Node_Value){
					v, err := getNodeValue(data, data.Nodes[n].Parts[p].Node_ID)
					if (err == nil){
						e = strings.Replace(e, "|"+strconv.Itoa(data.Nodes[n].Parts[p].ID)+"|", strconv.FormatFloat(v, 'g', 5, 64), 1)
					} else{
						return err
					}
				} else{
					v, err := getNodeAttributeValue(data, data.Nodes[n].Parts[p].Node_ID, data.Nodes[n].Parts[p].Attribute_ID)
					if (err == nil){
						e = strings.Replace(e, "|"+strconv.Itoa(data.Nodes[n].Parts[p].ID)+"|", strconv.FormatFloat(v, 'g', 5, 64), 1)
					} else{
						return err
					}
				}
			}
			// update data model equation
			data.Nodes[n].Equation = e

			// evaluate equation and update model
			fs := token.NewFileSet()
			tv, err := types.Eval(fs, nil, token.NoPos, e)
			if err != nil {
				// panic(err)
				return err
			}

			f, err := strconv.ParseFloat(tv.Value.String(), 64)
			if err != nil {
				return err
			} else {
				data.Nodes[n].Value = f
			}

		} else {

			// if there are no link types which use the 'e' and 'v' variables
			_ = e
			_ = v
		}
	}
	return nil
}