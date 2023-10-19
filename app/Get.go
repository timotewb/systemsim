package app

import (
	"errors"
	"fmt"
	m "systemsim/models"
)

func GetNodeAttributeValue(data m.DataType, node_id int, attribute_id int)(float64, error){

	// loop over each of the nodes
	for n := 0; n < len(data.Nodes); n++{

		// check node_id
		if (data.Nodes[n].ID == node_id){

			// loop over the attributes
			for a := 0; a < len(data.Nodes[n].Attributes); a++{

				// check attribute_id
				if (data.Nodes[n].Attributes[a].ID == attribute_id){
					return float64(data.Nodes[n].Attributes[a].Value), nil
				} else{
					return 0, errors.New(fmt.Sprintf("Could not find attribute with id %v in node with id %v.", attribute_id, node_id))
				}
			}
		} else{
			return 0, errors.New(fmt.Sprintf("Could not find node with id %v.", node_id))
		}
	}
	return 0, errors.New("No nodes in data.")
}