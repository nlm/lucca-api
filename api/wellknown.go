package api

type ItemsList[T any] struct {
	Items []T     `json:"items"`
	Next  *string `json:"next"`
}
