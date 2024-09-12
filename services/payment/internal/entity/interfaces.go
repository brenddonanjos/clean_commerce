package entity

type CardRepositoryInterface interface {
	Save(card *Card) (*Card, error)
}
