package entity

type CategoryEntity struct{
	CategoryID   int
	CategoryName string
}

type IusecaseCategory interface {
	GetCategory() (data []CategoryEntity, err error)
}

type IrepoCategory interface {
	GetCategory() (data []CategoryEntity, err error)
}