package models

//----------------------------------------------------------------------------------------
// Part
//----------------------------------------------------------------------------------------
type PartType struct{
	ID int	`json:"id"`
	Node_ID int	`json:"node_id"`
	Attribute_ID int `json:"attibute_id"`
	Is_Node_Value bool `json:"is_node_value"`
}