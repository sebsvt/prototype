package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	ordering_adapters "github.com/sebsvt/prototype/adapters/ordering"
	ordering_domain "github.com/sebsvt/prototype/domain/ordering"
	ordering_service "github.com/sebsvt/prototype/services/ordering"
)

func main() {
	memory_adapter := ordering_adapters.NewOrderRepositoryMemory()
	srv := ordering_service.NewOrderService(memory_adapter)
	order, err := srv.CreateNewOrder(ordering_domain.OrderCreated{
		CustomerID: uuid.New().String(),
		ProductID:  uuid.New().String(),
	})
	if err != nil {
		log.Println("Ah: ", err)
	}
	has_order, err := srv.GetOrderByOrderReference(order.Reference)
	if err != nil {
		log.Println("Ah: ", err)
	}
	fmt.Println(has_order)
}
