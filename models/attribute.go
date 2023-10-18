package models

//----------------------------------------------------------------------------------------
// Attribute
//----------------------------------------------------------------------------------------
type AttributeType struct{
	ID int	`json:"id"`
	Description string `json:"description"`
	Value float64 `json:"value"`
}