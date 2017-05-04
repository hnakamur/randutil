package randutil

import "github.com/pkg/errors"

type Choice struct {
	Weight int
	Item   interface{}
}

type Chooser struct {
	intner      Intner
	choices     []Choice
	totalWeight int
}

func calcTotalWeight(choices []Choice) int {
	total := 0
	for _, c := range choices {
		total += c.Weight
	}
	return total
}

func NewChooser(intner Intner, choices []Choice) (*Chooser, error) {
	totalWeight := calcTotalWeight(choices)
	if totalWeight <= 0 {
		return nil, errors.New("total weight must be positive")
	}
	return &Chooser{
		intner:      intner,
		choices:     choices,
		totalWeight: totalWeight,
	}, nil
}

func (c *Chooser) Choose() (interface{}, error) {
	v, err := c.intner.Intn(c.totalWeight)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get random int for chooser")
	}
	for _, choice := range c.choices {
		if v < choice.Weight {
			return choice.Item, nil
		}
		v -= choice.Weight
	}
	panic("this should not happen")
}
