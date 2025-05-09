package pagination

type Pagination struct {
	Page  int
	Limit int
}

func (_this *Pagination) GetPage() int {
	return _this.Page
}
func (_this *Pagination) GetOffset() int {

	if _this.Page <= 0 {
		_this.Page = 1
	}
	return (_this.Page - 1) * _this.Limit
}
func (_this *Pagination) GetLimit() int {
	return _this.Limit
}
