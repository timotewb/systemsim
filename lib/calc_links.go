package lib

import (
	"fmt"
	"go/token"
	"go/types"
	"strconv"
	"strings"
	m "systemsim/models"
)

func CalcLinks(s m.SimType){

	// for each link
	var e string
	var j int
	for i := 0; i < len(s.Links); i++ {

		// create executable equation
		e = s.Links[i].Equation

		// for each part (part of the equation) check its type
		for j = 0; j <len(s.Links[i].Parts); j++{
			fmt.Println(s.Links[i].Parts[j].LinkType)

			// if link type is attribute (part from from an input's attribute value)
			if s.Links[i].Parts[j].LinkType == "attribute" {
				e = strings.Replace(e, "|"+strconv.Itoa(s.Links[i].Parts[j].ID)+"|", strconv.FormatFloat(getInputAttributeValue(s, s.Links[i].Parts[j].InputID, s.Links[i].Parts[j].AttributeID), 'g', 5, 64), 1)
			} else if s.Links[i].Parts[j].LinkType == "link" {
				e = strings.Replace(e, "|"+strconv.Itoa(s.Links[i].Parts[j].ID)+"|", strconv.FormatFloat(getLinkOutputValue(s, s.Links[i].Parts[j].LinkID), 'g', 5, 64), 1)
			}
		}
		s.Links[i].EquationEvaluated = e

		// evaluate equation
		fs := token.NewFileSet()
		tv, err := types.Eval(fs, nil, token.NoPos, e)
		if err != nil {
			panic(err)
		}

		// udpate output
		s.Links[i].Output = tv.Value.String()

		fmt.Println(e)
		fmt.Println(s.Links[i].Output)
	}
}

func getInputAttributeValue(s m.SimType ,inputID int, attributeID int,) (float64){

	var r float64

	// loop over each of the inputs
	for i := 0; i < len(s.Inputs); i++ {
		if s.Inputs[i].ID == inputID{

			// loop over each of the attributes
			for j := 0; j < len(s.Inputs); j++ {
				if s.Inputs[i].Attributes[j].ID == attributeID{
					fmt.Println(s.Inputs[i].Attributes[j].Value)
					r = float64(s.Inputs[i].Attributes[j].Value)
					fmt.Println()
				}
			}
		}
	}
	return r
}

func getLinkOutputValue(s m.SimType ,linkID int,) (float64){

	var r float64

	// loop over each of the links
	for i := 0; i < len(s.Links); i++ {
		if s.Links[i].ID == linkID{
			r, _ = strconv.ParseFloat(s.Links[i].Output, 64)
		}
	}
	return r
}