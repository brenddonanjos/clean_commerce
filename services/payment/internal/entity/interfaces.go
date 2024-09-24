package entity

type CardRepositoryInterface interface {
	Save(card *Card) (*Card, error)
}

type BillingAddressRepositoryInterface interface {
	Save(address *BillingAddress) (*BillingAddress, error)
}
