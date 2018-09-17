package main

type Satisfaccion struct {
	Comentario_NLU     float64 `bson:"Comentario_NLU" json:"Comentario_NLU"`
	Valoracion     float64   `bson:"Valoracion" json:"Valoracion"`
	Cantidad     int   `bson:"Cantidad" json:"Cantidad"`
}
type Satisfaccions []Satisfaccion
