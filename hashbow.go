package hashbow

import colorful "github.com/lucasb-eyer/go-colorful"

func Hashbow(value interface{}) string {
	return Color(value, 50, 50).Hex()
}

func Color(value interface{}, saturation float64, lightness float64) colorful.Color {
	var sum float64

	switch v := value.(type) {
	case string:
		chars := []rune(v)
		for _, ch := range chars {
			sum += float64(ch)
		}
	case float64:
		sum = v
	case int:
		sum = float64(v)
	default:
		sum = 0
	}

	sum = sum * sum
	return colorful.Hsl(float64(int(sum)%360), saturation, lightness)
}
