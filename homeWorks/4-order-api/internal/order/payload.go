package order

type GetProductResponce struct{
	Success bool `json:"success"`
	Product
}

type AddProductRequest struct{
	Product
}

type AddProductResponce struct{
	Success bool `json:"success"`
}
