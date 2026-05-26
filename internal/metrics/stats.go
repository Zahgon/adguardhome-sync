package metrics

import (
	"github.com/bakito/adguardhome-sync/internal/client/model"
)

const labelTotal = "Total"

var (
	blue             = []int{78, 141, 245}
	blueAlternatives = [][]int{
		{44, 95, 163},
		{122, 166, 247},
		{30, 61, 92},
		{93, 158, 255},
		{58, 123, 213},
	}

	red             = []int{255, 94, 94}
	redAlternatives = [][]int{
		{204, 59, 59},
		{255, 127, 127},
		{140, 36, 36},
		{255, 153, 153},
		{255, 66, 66},
	}

	yellow             = []int{232, 198, 78}
	yellowAlternatives = [][]int{
		{196, 163, 60},
		{255, 220, 110},
		{140, 114, 36},
		{250, 233, 156},
		{212, 180, 84},
	}

	green             = []int{110, 224, 122}
	greenAlternatives = [][]int{
		{68, 160, 80},
		{142, 255, 158},
		{44, 140, 63},
		{163, 255, 192},
		{85, 198, 102},
	}
)

func StatsGraph() (t *model.Stats, dns, blocked, malware, adult []Line) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

func safeStats(stats *[]int) []int { _ = "STUB: not implemented"; return nil }

func graphLines(
	t *model.Stats,
	s OverallStats,
	baseColor []int,
	altColors [][]int,
	dataCB func(s *model.Stats) []int,
) []Line {
	_ = "STUB: not implemented"
	return nil
}

type graph struct {
	total    Line
	replicas []Line
}

type Line struct {
	Data  []int  `json:"data"`
	R     int    `json:"r"`
	G     int    `json:"g"`
	B     int    `json:"b"`
	Title string `json:"title"`
	Fill  bool   `json:"fill"`
}
