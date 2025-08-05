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
	actual := CalculateOrder(Order{
		RedSet:    0,
		GreenSet:  3,
		BlueSet:   0,
		YellowSet: 0,
		PinkSet:   0,
		PurpleSet: 0,
		OrangeSet: 0,
		IsMember:  false})

	assert.Equal(t, float64(116), actual)
}

func TestCalculateOrderHaveMember(t *testing.T) {
	actual := CalculateOrder(Order{
		RedSet:    0,
		GreenSet:  3,
		BlueSet:   0,
		YellowSet: 0,
		PinkSet:   0,
		PurpleSet: 0,
		OrangeSet: 0,
		IsMember:  true})

	assert.Equal(t, 104.4, actual)
}

func TestCalculateOrderAllSetWithThreeOrderAndHaveMember(t *testing.T) {
	actual := CalculateOrder(Order{
		RedSet:    3,
		GreenSet:  3,
		BlueSet:   3,
		YellowSet: 3,
		PinkSet:   3,
		PurpleSet: 3,
		OrangeSet: 3,
		IsMember:  true})

	assert.Equal(t, 1220.4, actual)
}
