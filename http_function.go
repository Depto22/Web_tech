package main

import "net/http"
func userHandler(){
	
}
func main(){
	http.HandleFunc("/depto",userHandler)
	http.ListenAndServe(":8084",nil)
}