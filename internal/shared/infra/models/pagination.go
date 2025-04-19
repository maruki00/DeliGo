package shared_models

type Pagination struct {
	Page   int
	Offset int
	Limit  int
}

func (p *Pagination) GetPage() int {
	return p.Page
}
func (p *Pagination) GetOffset() int {
	return p.Offset
}
func (p *Pagination) GetLimit() int {
	return p.Limit
}
