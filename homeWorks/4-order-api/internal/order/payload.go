package order

import "4-order-api/internal/store"


type GetProductResponce struct{
	Success bool
	store.Product
}

type AddProductRequest struct{
	store.Product
}

type AddProductResponce struct{
	Success bool
}
