package helper

import (
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatThousandSeparatorFloat(n float64) string {
	p := message.NewPrinter(language.English)

	if n == math.Floor(n) {
		return p.Sprintf("%.0f", n)
	}
	return p.Sprintf("%.2f", n)
}

func FormatThousandSeparatorInt(n int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", n)
}
