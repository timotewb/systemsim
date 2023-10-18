package models

//----------------------------------------------------------------------------------------
// Node
//----------------------------------------------------------------------------------------
type NodeType struct{
	ID int	`json:"id"`
	Type string `json:"type"`
	Description string `json:"description"`
	Attributes []AttributeType `json:"attributes"`
	Equation string `json:"equation"`
	Parts []PartType `json:"parts"`
	Output float64 `json:"output"`
}