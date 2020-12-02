package main

import (
	"Week02/service"
	"fmt"
)

func main() {
	customerService := service.NewCustomerService()
	var id = "10001"
	_, err := customerService.GetCustomerById(id)
	fmt.Printf("%v\n", err)
}
