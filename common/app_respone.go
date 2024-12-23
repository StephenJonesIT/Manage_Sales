package common

type SuccessRes struct{
	Data 	interface{} 	`json:"data"`
	Paging 	interface{} 	`json:"paging,omitempty"`
	Filter 	interface{} 	`json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *SuccessRes{
	return &SuccessRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *SuccessRes{
	return NewSuccessResponse(data, nil, nil)
}

type SuccessSupplier struct{
	Data 	interface{} 	`json:"data"`
	Paging 	interface{} 	`json:"paging,omitempty"`
}

func NewSuccessResponseSupplier(data, paging interface{}) *SuccessSupplier{
	return &SuccessSupplier{Data: data, Paging: paging}
}

func SimpleSuccessResponseSupplier(data interface{}) *SuccessSupplier{
	return NewSuccessResponseSupplier(data, nil)
}