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

type OrderImplement interface {
	CalculateOrder(req OrderRequest) float64
}

type orderImplement struct {
	Menu                  map[string]float64
	MenuBundleDiscount    map[string]bool
	MemberDiscountPercent float64
	BundleDiscountPercent float64
}

func NewOrder() OrderImplement {
	return &orderImplement{
		Menu: map[string]float64{
			"Red Set":    redSet,
			"Green Set":  greenSet,
			"Blue Set":   blueSet,
			"Yellow Set": yellowSet,
			"Pink Set":   pinkSet,
			"Purple Set": purpleSet,
			"Orange Set": orangeSet,
		},
		MenuBundleDiscount: map[string]bool{
			"Green Set":  true,
			"Pink Set":   true,
			"Orange Set": true,
		},
		MemberDiscountPercent: memberDiscountPercent,
		BundleDiscountPercent: bundleDiscountPercent,
	}
}

type OrderRequest struct {
	Items    []Item
	IsMember bool
}

type Item struct {
	Menu  string
	Count int
}

func (s *orderImplement) CalculateOrder(req OrderRequest) float64 {
	var total float64
	for _, order := range req.Items {
		if s.MenuBundleDiscount[order.Menu] {
			total += bundleDiscount(s.Menu[order.Menu], order.Count)
			continue
		}
		total += s.Menu[order.Menu] * float64(order.Count)
	}

	if req.IsMember {
		return total * (float64(1) - (s.MemberDiscountPercent / 100))
	}

	return total
}

func bundleDiscount(price float64, orderNumber int) float64 {
	discountPrice := (price * 2) * (bundleDiscountPercent / 100)
	bundle := float64(orderNumber / 2)
	return (price * float64(orderNumber)) - (discountPrice * bundle)
}
