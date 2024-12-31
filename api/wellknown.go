package api

// ItemsList is a well-known object for item lists.
type ItemsList[T any] struct {
	Items []T     `json:"items,omitempty"`
	Next  *string `json:"next,omitempty"`
	Prev  *string `json:"prev,omitempty"`
}

type HeaderData[H, D any] struct {
	Header H `json:"header,omitempty"`
	Data   D `json:"data,omitempty"`
}
