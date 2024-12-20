package api

type Page struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"limit,omitempty"`
}

// func (p Page) MarshalJSON() ([]byte, error) {
// 	return []byte(fmt.Sprintf("%d,%d", p.Page, p.PageSize)), nil
// }
