package main

const (
	redSet                float64 = 50
	greenSet              float64 = 40
	blueSet               float64 = 30
	yellowSet             float64 = 50
	pinkSet               float64 = 80
	purpleSet             float64 = 90
	orangeSet             float64 = 120
	bundleDiscountPercent float64 = 5
	memberDiscountPercent float64 = 10
)

type Order struct {
	RedSet    int
	GreenSet  int
	BlueSet   int
	YellowSet int
	PinkSet   int
	PurpleSet int
	OrangeSet int
	IsMember  bool
}

func CalculateOrder(req Order) float64 {
	redTotal := redSet * float64(req.RedSet)
	greenTotal := bundleDiscount(greenSet, req.GreenSet)
	blueTotal := blueSet * float64(req.BlueSet)
	yellowTotal := yellowSet * float64(req.YellowSet)
	pinkTotal := bundleDiscount(pinkSet, req.PinkSet)
	purpleTotal := purpleSet * float64(req.PurpleSet)
	orangeTotal := bundleDiscount(orangeSet, req.OrangeSet)
	total := redTotal + greenTotal + blueTotal + yellowTotal + pinkTotal + purpleTotal + orangeTotal
	if req.IsMember {
		return total * (float64(1) - (memberDiscountPercent / 100))
	}

	return total
}

func bundleDiscount(price float64, orderNumber int) float64 {
	discountPrice := (price * 2) * (bundleDiscountPercent / 100)
	bundle := float64(orderNumber / 2)
	return (price * float64(orderNumber)) - (discountPrice * bundle)
}
