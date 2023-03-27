package abstractfactory

import "fmt"

type ISportsFactory interface {
	MakeShoes() IShoe
	MakeShirt() IShirt
}

type Brand int64

const (
	AdidasType Brand = 0
)

func (b Brand) GetFactory() (ISportsFactory, error) {
	switch b {
	case AdidasType:
		return &Adidas{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
