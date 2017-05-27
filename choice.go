package randutil

import (
	"fmt"

	"errors"
)

// Choice represents an item to be choosed. It has a weight which affects possibility to be choosed.
type Choice struct {
	Weight uint
	Item   interface{}
}

// Chooser represents a chooser that chooses an item randomly.
type Chooser struct {
	intner      Intner
	choices     []Choice
	totalWeight int
}

func calcTotalWeight(choices []Choice) int {
	var total int
	for _, c := range choices {
		total += int(c.Weight)
	}
	return total
}

// NewChooser returns a new chooser for choices.
// It returns an error if total weight of choices is not positive.
func NewChooser(intner Intner, choices []Choice) (*Chooser, error) {
	totalWeight := calcTotalWeight(choices)
	if totalWeight == 0 {
		return nil, errors.New("total weight must be positive")
	}
	return &Chooser{
		intner:      intner,
		choices:     choices,
		totalWeight: totalWeight,
	}, nil
}

// Choose returns an randomly selected item from choices.
// It returns an error if the underlying intner returned en error.
func (c *Chooser) Choose() (interface{}, error) {
	v, err := c.intner.Intn(c.totalWeight)
	if err != nil {
		return nil, fmt.Errorf("failed to get random int for chooser, err=%v", err)
	}
	for _, choice := range c.choices {
		if v < int(choice.Weight) {
			return choice.Item, nil
		}
		v -= int(choice.Weight)
	}
	panic("this should not happen")
}
