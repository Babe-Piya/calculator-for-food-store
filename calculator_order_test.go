package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBundleDiscountOneOrder(t *testing.T) {
	actual := bundleDiscount(50, 1)

	assert.Equal(t, float64(50), actual)
}

func TestBundleDiscountTwoOrder(t *testing.T) {
	actual := bundleDiscount(50, 2)

	assert.Equal(t, float64(95), actual)
}

func TestBundleDiscountThreeOrder(t *testing.T) {
	actual := bundleDiscount(50, 3)

	assert.Equal(t, float64(145), actual)
}

func TestCalculateOrderNotMember(t *testing.T) {
	order := NewOrder()
	actual := order.CalculateOrder(OrderRequest{
		Items: []Item{{
			Menu:  "Green Set",
			Count: 3,
		}},
		IsMember: false})

	assert.Equal(t, float64(116), actual)
}

func TestCalculateOrderHaveMember(t *testing.T) {
	order := NewOrder()
	actual := order.CalculateOrder(OrderRequest{
		Items: []Item{{
			Menu:  "Green Set",
			Count: 3,
		}},
		IsMember: true})

	assert.Equal(t, 104.4, actual)
}

func TestCalculateOrderAllSetWithThreeOrderAndHaveMember(t *testing.T) {
	order := NewOrder()
	actual := order.CalculateOrder(OrderRequest{
		Items: []Item{
			{
				Menu:  "Red Set",
				Count: 3,
			}, {
				Menu:  "Green Set",
				Count: 3,
			}, {
				Menu:  "Blue Set",
				Count: 3,
			},
			{
				Menu:  "Yellow Set",
				Count: 3,
			},
			{
				Menu:  "Pink Set",
				Count: 3,
			},
			{
				Menu:  "Purple Set",
				Count: 3,
			},
			{
				Menu:  "Orange Set",
				Count: 3,
			},
		},
		IsMember: true})

	assert.Equal(t, 1220.4, actual)
}
