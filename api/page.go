package api

import "fmt"

type Page struct {
	Page     int
	PageSize int
}

func (p Page) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d,%d", p.Page, p.PageSize)), nil
}
