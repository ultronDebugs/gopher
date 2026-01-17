package handler

import (
	"fmt"
	"net/http"
)

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating an order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {

	fmt.Println("here is the orders list")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("get by Order ID")
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update order By Id")
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Delete order By Id")

}
