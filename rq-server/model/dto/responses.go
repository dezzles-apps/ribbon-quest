package dto

type Page[T any] struct {
	Items []T     `json:"items"`
	Next  *string `json:"next"`
}
