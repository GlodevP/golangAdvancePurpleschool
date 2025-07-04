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

type UpdateRequest struct{
	Product
}

type UpdateResponse struct{
	Success bool `json:"success"`
	Product
}

type DeletResponse struct{
	Success bool `json:"success"`
}
