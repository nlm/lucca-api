package api

// type Header struct {
// 	Generated  time.Time `json:"generated,omitempty"`
// 	ServerTime int       `json:"serverTime,omitempty"`
// 	QueryTime  int       `json:"queryTime,omitempty"`
// 	QueryCount int       `json:"queryCount,omitempty"`
// 	Principal  string    `json:"principal,omitempty"`
// 	ProcessId  int       `json:"processId,omitempty"`
// }

type DataList[T any] struct {
	// Header Header   `json:"header,omitempty"`
	Data ItemsList[T] `json:"data,omitempty"`
}

type DataItem[T any] struct {
	// Header Header   `json:"header,omitempty"`
	Data T `json:"data,omitempty"`
}

type ItemsList[T any] struct {
	Items []T    `json:"items"`
	Next  string `json:"next"`
}
