package dashboard

type Badge struct {
	Label string `json:"label"`
	Tag   string `json:"tag"`
	Color string `json:"color"`
}

func (badge Badge) labelWidth() float64 {
	return float64(len(badge.Label))*6.5 + 3
}

func (badge Badge) tagWidth() float64 {
	return float64(len(badge.Tag)) * 6.5
}

func (badge Badge) endWidth() float64 {
	return badge.labelWidth() + badge.tagWidth() + 13
}
