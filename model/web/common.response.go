package web

type CommonResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}
