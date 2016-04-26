package romannumerals

type Magnitude struct {
	i string
	v string
	x string
}

var mags = [...]Magnitude{
	{i: "I", v: "V", x: "X"},
	{i: "X", v: "L", x: "C"},
	{i: "C", v: "D", x: "M"},
	{i: "M", v: "_", x: "_"},
}

func placeString(mag Magnitude, digit int) string {
	switch digit {
	case 1:
		return mag.i
	case 2:
		return mag.i + mag.i
	case 3:
		return mag.i + mag.i + mag.i
	case 4:
		return mag.i + mag.v
	case 5:
		return mag.v
	case 6:
		return mag.v + mag.i
	case 7:
		return mag.v + mag.i + mag.i
	case 8:
		return mag.v + mag.i + mag.i + mag.i
	case 9:
		return mag.i + mag.x
	}

	return ""
}
