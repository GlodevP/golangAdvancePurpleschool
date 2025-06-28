package order

import "4-order-api/internal/store"


type GetOrderResponce struct{
	Success bool
	store.Order
}

type AddOrderRequest struct{
	store.Order
}

type AddOrderResponce struct{
	Success bool
}
