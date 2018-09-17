package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "github.com/gorilla/mux"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

var collection =  getSession().DB("interacciones").C("satisfaccions")

func SatisList(w http.ResponseWriter, r *http.Request) {
	var results []Satisfaccion 
	err := collection.Find(nil).All(&results)

	var aum float64= 0
	var prom float64 = 0
	var sati Satisfaccion
	var aum2 float64= 0
	var prom2 float64= 0
	var numero  = len(results)
	 f := float64(numero)
	

	if err != nil {
		log.Fatal(err)
	}else {

		for i:= 0; i < len(results); i++ {
			var arr = results[i]
			var arr2 = arr.Valoracion
			a := float64(arr2)
			aum = aum + a
			prom = aum / f
			// fmt.Println("prueba",arr2,prom)
		}
		
		for i:= 0; i < len(results); i++ {
			var arr = results[i]
			var arr2 = arr.Comentario_NLU
			a := float64(arr2)
			aum2 = aum2 + a
			prom2 = aum2 / f
		}
		fmt.Println("prom2",prom2)
		sati.Cantidad = numero
		sati.Comentario_NLU = prom2
		sati.Valoracion = prom
		fmt.Println("resultados",results,prom,len(results))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	// json.NewEncoder(w).Encode(results)
	json.NewEncoder(w).Encode(sati)

}

