package dtos

type CreateNewOrderDTO struct {
	CostomerId int
	Products   map[int]int
}
