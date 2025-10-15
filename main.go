package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
 func main (){
	// fmt.Println("hello go")
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
		http.HandleFunc("/health", handleHealth)
	log.Println("..listening on localhost:4242")
	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil{
		log.Fatal(err)
	}

}

	func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request){
		// createResponse := [] byte ("you want to create payment?")
		// _, err := writer.Write(createResponse)
		// if err !=nil{
		// 	fmt.Println(err)
		// }
		// fmt.Println("endpoint called")

		if request.Method != "POST" {
			http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		// fmt.Println("request method correct")


	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName string `json:"Last_name"`
		Address string `json:"address"`
		City string `json:"city"`
		Country string `json:"country"`
	}

err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil{
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(req.FirstName)
	fmt.Println(req.LastName)
		fmt.Println(req.Address)
			fmt.Println(req.City)
				fmt.Println(req.Country)
					fmt.Println(req.ProductId)

}


func handleHealth(writer http.ResponseWriter, request *http.Request){
	response := []byte ("server is up and running")
	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}

