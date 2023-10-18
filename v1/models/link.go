package models

//----------------------------------------------------------------------------------------
// Link
//----------------------------------------------------------------------------------------
type LinkType struct{
	ID int	`json:"id"`
	Name string `json:"name"`
	Parts []PartType `json:"parts"`
	Equation string `json:"equation"`
	EquationEvaluated string `json:"equation_evaluated"`
	Output string `json:"output"`
}

type PartType struct{
	ID int	`json:"id"`
	LinkType string `json:"link_type"`
	LinkID int `json:"link_id"`
	InputID int `json:"input_id"`
	AttributeID int `json:"attribute_id"`
}