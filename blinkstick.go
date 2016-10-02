package main

import (
	"flag"
	"fmt"
	"github.com/boombuler/led"
	"image/color"
	"time"
)

var (
	col       string
	lighttype string
	COLOR_OFF = color.RGBA{0x00, 0x00, 0x00, 0xff}
)

func main() {
	flag.Usage = func() {
		fmt.Printf("The tool 'blinkstick' provides a simple functions to play with colors for the first found BlinkStick device\n")
		fmt.Printf("Usage:\n")
		fmt.Printf("blinkstick -color <colorname>\n")
		fmt.Printf("blinkstick -color <colorname> -lighttype blink\n")
		fmt.Printf("Example:\n")
		fmt.Printf("blinkstick -color blue -lighttype blink -duration 100 -times 10\n\n")
		flag.PrintDefaults()
	}

	flag.StringVar(&col, "color", "black", "color (for example, red, lime, white, etc. or off")
	flag.StringVar(&lighttype, "lighttype", "static", "lighttype (static, blink or pulse)")
	duration := flag.Int64("duration", 300, "time between blinks")
	times := flag.Int("times", 5, "defines how many times it should blink")
	steps := flag.Int("steps", 15, "steps between pulse color and black")
	flag.Parse()

	color, err := LookupColorName(col)
	if err != nil {
		fmt.Println(err)
		return
	}

	for devInfo := range led.Devices() {
		// keep sure to find BlinkStick devices only
		if devInfo.GetType() != led.BlinkStick {
			continue
		}

		dev, err := devInfo.Open()

		// stop on error
		if err != nil {
			fmt.Println(err)
			break
		}
		defer dev.Close()

		dev.SetKeepActive(true)
		Static(dev, COLOR_OFF)
		switch lighttype {
		case "static":
			Static(dev, color)
		case "blink":
			Blink(dev, color, *duration, *times)
		case "pulse":
			Pulse(dev, color, *times, *steps)
		}
		time.Sleep(20 * time.Millisecond)

		return
	}
}

// Static color
func Static(d led.Device, col color.Color) {
	time.Sleep(20 * time.Millisecond)
	d.SetColor(col)
}

// Blink color with given duration and times
func Blink(d led.Device, col color.Color, dur int64, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(time.Duration(dur) * time.Millisecond)
		Static(d, col)
		time.Sleep(time.Duration(dur) * time.Millisecond)
		Static(d, COLOR_OFF)
	}
}

// Blink color with given steps and times
func Pulse(d led.Device, col color.Color, times int, steps int) {
	r, g, b, a := col.RGBA()
	for i := 0; i < times; i++ {
		for j := 0; j < steps; j++ {
			r2 := r * uint32(j) / uint32(steps)
			g2 := g * uint32(j) / uint32(steps)
			b2 := b * uint32(j) / uint32(steps)
			Static(d, color.RGBA{byte(r2 >> 8), byte(g2 >> 8), byte(b2 >> 8), uint8(a)})
		}
		for j := steps; j > 0; j-- {
			r2 := r * uint32(j) / uint32(steps)
			g2 := g * uint32(j) / uint32(steps)
			b2 := b * uint32(j) / uint32(steps)
			Static(d, color.RGBA{byte(r2 >> 8), byte(g2 >> 8), byte(b2 >> 8), uint8(a)})
		}
		Static(d, COLOR_OFF)
	}
}

func LookupColorName(s string) (color.Color, error) {
	colors := map[string]color.RGBA{
		"aliceblue":            color.RGBA{0xf0, 0xf8, 0xff, 0xff},
		"antiquewhite":         color.RGBA{0xfa, 0xeb, 0xd7, 0xff},
		"aqua":                 color.RGBA{0x00, 0xff, 0xff, 0xff},
		"aquamarine":           color.RGBA{0x7f, 0xff, 0xd4, 0xff},
		"azure":                color.RGBA{0xf0, 0xff, 0xff, 0xff},
		"beige":                color.RGBA{0xf5, 0xf5, 0xdc, 0xff},
		"bisque":               color.RGBA{0xff, 0xe4, 0xc4, 0xff},
		"black":                color.RGBA{0x00, 0x00, 0x00, 0xff},
		"blanchedalmond":       color.RGBA{0xff, 0xeb, 0xcd, 0xff},
		"blue":                 color.RGBA{0x00, 0x00, 0xff, 0xff},
		"blueviolet":           color.RGBA{0x8a, 0x2b, 0xe2, 0xff},
		"brown":                color.RGBA{0xa5, 0x2a, 0x2a, 0xff},
		"burlywood":            color.RGBA{0xde, 0xb8, 0x87, 0xff},
		"cadetblue":            color.RGBA{0x5f, 0x9e, 0xa0, 0xff},
		"chartreuse":           color.RGBA{0x7f, 0xff, 0x00, 0xff},
		"chocolate":            color.RGBA{0xd2, 0x69, 0x1e, 0xff},
		"coral":                color.RGBA{0xff, 0x7f, 0x50, 0xff},
		"cornflowerblue":       color.RGBA{0x64, 0x95, 0xed, 0xff},
		"cornsilk":             color.RGBA{0xff, 0xf8, 0xdc, 0xff},
		"crimson":              color.RGBA{0xdc, 0x14, 0x3c, 0xff},
		"cyan":                 color.RGBA{0x00, 0xff, 0xff, 0xff},
		"darkblue":             color.RGBA{0x00, 0x00, 0x8b, 0xff},
		"darkcyan":             color.RGBA{0x00, 0x8b, 0x8b, 0xff},
		"darkgoldenrod":        color.RGBA{0xb8, 0x86, 0x0b, 0xff},
		"darkgray":             color.RGBA{0xa9, 0xa9, 0xa9, 0xff},
		"darkgrey":             color.RGBA{0xa9, 0xa9, 0xa9, 0xff},
		"darkgreen":            color.RGBA{0x00, 0x64, 0x00, 0xff},
		"darkkhaki":            color.RGBA{0xbd, 0xb7, 0x6b, 0xff},
		"darkmagenta":          color.RGBA{0x8b, 0x00, 0x8b, 0xff},
		"darkolivegreen":       color.RGBA{0x55, 0x6b, 0x2f, 0xff},
		"darkorange":           color.RGBA{0xff, 0x8c, 0x00, 0xff},
		"darkorchid":           color.RGBA{0x99, 0x32, 0xcc, 0xff},
		"darkred":              color.RGBA{0x8b, 0x00, 0x00, 0xff},
		"darksalmon":           color.RGBA{0xe9, 0x96, 0x7a, 0xff},
		"darkseagreen":         color.RGBA{0x8f, 0xbc, 0x8f, 0xff},
		"darkslateblue":        color.RGBA{0x48, 0x3d, 0x8b, 0xff},
		"darkslategray":        color.RGBA{0x2f, 0x4f, 0x4f, 0xff},
		"darkslategrey":        color.RGBA{0x2f, 0x4f, 0x4f, 0xff},
		"darkturquoise":        color.RGBA{0x00, 0xce, 0xd1, 0xff},
		"darkviolet":           color.RGBA{0x94, 0x00, 0xd3, 0xff},
		"deeppink":             color.RGBA{0xff, 0x14, 0x93, 0xff},
		"deepskyblue":          color.RGBA{0x00, 0xbf, 0xff, 0xff},
		"dimgray":              color.RGBA{0x69, 0x69, 0x69, 0xff},
		"dimgrey":              color.RGBA{0x69, 0x69, 0x69, 0xff},
		"dodgerblue":           color.RGBA{0x1e, 0x90, 0xff, 0xff},
		"firebrick":            color.RGBA{0xb2, 0x22, 0x22, 0xff},
		"floralwhite":          color.RGBA{0xff, 0xfa, 0xf0, 0xff},
		"forestgreen":          color.RGBA{0x22, 0x8b, 0x22, 0xff},
		"fuchsia":              color.RGBA{0xff, 0x00, 0xff, 0xff},
		"gainsboro":            color.RGBA{0xdc, 0xdc, 0xdc, 0xff},
		"ghostwhite":           color.RGBA{0xf8, 0xf8, 0xff, 0xff},
		"gold":                 color.RGBA{0xff, 0xd7, 0x00, 0xff},
		"goldenrod":            color.RGBA{0xda, 0xa5, 0x20, 0xff},
		"gray":                 color.RGBA{0x80, 0x80, 0x80, 0xff},
		"grey":                 color.RGBA{0x80, 0x80, 0x80, 0xff},
		"green":                color.RGBA{0x00, 0x80, 0x00, 0xff},
		"greenyellow":          color.RGBA{0xad, 0xff, 0x2f, 0xff},
		"honeydew":             color.RGBA{0xf0, 0xff, 0xf0, 0xff},
		"hotpink":              color.RGBA{0xff, 0x69, 0xb4, 0xff},
		"indianred":            color.RGBA{0xcd, 0x5c, 0x5c, 0xff},
		"indigo":               color.RGBA{0x4b, 0x00, 0x82, 0xff},
		"ivory":                color.RGBA{0xff, 0xff, 0xf0, 0xff},
		"khaki":                color.RGBA{0xf0, 0xe6, 0x8c, 0xff},
		"lavender":             color.RGBA{0xe6, 0xe6, 0xfa, 0xff},
		"lavenderblush":        color.RGBA{0xff, 0xf0, 0xf5, 0xff},
		"lawngreen":            color.RGBA{0x7c, 0xfc, 0x00, 0xff},
		"lemonchiffon":         color.RGBA{0xff, 0xfa, 0xcd, 0xff},
		"lightblue":            color.RGBA{0xad, 0xd8, 0xe6, 0xff},
		"lightcoral":           color.RGBA{0xf0, 0x80, 0x80, 0xff},
		"lightcyan":            color.RGBA{0xe0, 0xff, 0xff, 0xff},
		"lightgoldenrodyellow": color.RGBA{0xfa, 0xfa, 0xd2, 0xff},
		"lightgray":            color.RGBA{0xd3, 0xd3, 0xd3, 0xff},
		"lightgrey":            color.RGBA{0xd3, 0xd3, 0xd3, 0xff},
		"lightgreen":           color.RGBA{0x90, 0xee, 0x90, 0xff},
		"lightpink":            color.RGBA{0xff, 0xb6, 0xc1, 0xff},
		"lightsalmon":          color.RGBA{0xff, 0xa0, 0x7a, 0xff},
		"lightseagreen":        color.RGBA{0x20, 0xb2, 0xaa, 0xff},
		"lightskyblue":         color.RGBA{0x87, 0xce, 0xfa, 0xff},
		"lightslategray":       color.RGBA{0x77, 0x88, 0x99, 0xff},
		"lightslategrey":       color.RGBA{0x77, 0x88, 0x99, 0xff},
		"lightsteelblue":       color.RGBA{0xb0, 0xc4, 0xde, 0xff},
		"lightyellow":          color.RGBA{0xff, 0xff, 0xe0, 0xff},
		"lime":                 color.RGBA{0x00, 0xff, 0x00, 0xff},
		"limegreen":            color.RGBA{0x32, 0xcd, 0x32, 0xff},
		"linen":                color.RGBA{0xfa, 0xf0, 0xe6, 0xff},
		"magenta":              color.RGBA{0xff, 0x00, 0xff, 0xff},
		"maroon":               color.RGBA{0x80, 0x00, 0x00, 0xff},
		"mediumaquamarine":     color.RGBA{0x66, 0xcd, 0xaa, 0xff},
		"mediumblue":           color.RGBA{0x00, 0x00, 0xcd, 0xff},
		"mediumorchid":         color.RGBA{0xba, 0x55, 0xd3, 0xff},
		"mediumpurple":         color.RGBA{0x93, 0x70, 0xd8, 0xff},
		"mediumseagreen":       color.RGBA{0x3c, 0xb3, 0x71, 0xff},
		"mediumslateblue":      color.RGBA{0x7b, 0x68, 0xee, 0xff},
		"mediumspringgreen":    color.RGBA{0x00, 0xfa, 0x9a, 0xff},
		"mediumturquoise":      color.RGBA{0x48, 0xd1, 0xcc, 0xff},
		"mediumvioletred":      color.RGBA{0xc7, 0x15, 0x85, 0xff},
		"midnightblue":         color.RGBA{0x19, 0x19, 0x70, 0xff},
		"mintcream":            color.RGBA{0xf5, 0xff, 0xfa, 0xff},
		"mistyrose":            color.RGBA{0xff, 0xe4, 0xe1, 0xff},
		"moccasin":             color.RGBA{0xff, 0xe4, 0xb5, 0xff},
		"navajowhite":          color.RGBA{0xff, 0xde, 0xad, 0xff},
		"navy":                 color.RGBA{0x00, 0x00, 0x80, 0xff},
		"oldlace":              color.RGBA{0xfd, 0xf5, 0xe6, 0xff},
		"olive":                color.RGBA{0x80, 0x80, 0x00, 0xff},
		"olivedrab":            color.RGBA{0x6b, 0x8e, 0x23, 0xff},
		"orange":               color.RGBA{0xff, 0xa5, 0x00, 0xff},
		"orangered":            color.RGBA{0xff, 0x45, 0x00, 0xff},
		"orchid":               color.RGBA{0xda, 0x70, 0xd6, 0xff},
		"palegoldenrod":        color.RGBA{0xee, 0xe8, 0xaa, 0xff},
		"palegreen":            color.RGBA{0x98, 0xfb, 0x98, 0xff},
		"paleturquoise":        color.RGBA{0xaf, 0xee, 0xee, 0xff},
		"palevioletred":        color.RGBA{0xd8, 0x70, 0x93, 0xff},
		"papayawhip":           color.RGBA{0xff, 0xef, 0xd5, 0xff},
		"peachpuff":            color.RGBA{0xff, 0xda, 0xb9, 0xff},
		"peru":                 color.RGBA{0xcd, 0x85, 0x3f, 0xff},
		"pink":                 color.RGBA{0xff, 0xc0, 0xcb, 0xff},
		"plum":                 color.RGBA{0xdd, 0xa0, 0xdd, 0xff},
		"powderblue":           color.RGBA{0xb0, 0xe0, 0xe6, 0xff},
		"purple":               color.RGBA{0x80, 0x00, 0x80, 0xff},
		"red":                  color.RGBA{0xff, 0x00, 0x00, 0xff},
		"rosybrown":            color.RGBA{0xbc, 0x8f, 0x8f, 0xff},
		"royalblue":            color.RGBA{0x41, 0x69, 0xe1, 0xff},
		"saddlebrown":          color.RGBA{0x8b, 0x45, 0x13, 0xff},
		"salmon":               color.RGBA{0xfa, 0x80, 0x72, 0xff},
		"sandybrown":           color.RGBA{0xf4, 0xa4, 0x60, 0xff},
		"seagreen":             color.RGBA{0x2e, 0x8b, 0x57, 0xff},
		"seashell":             color.RGBA{0xff, 0xf5, 0xee, 0xff},
		"sienna":               color.RGBA{0xa0, 0x52, 0x2d, 0xff},
		"silver":               color.RGBA{0xc0, 0xc0, 0xc0, 0xff},
		"skyblue":              color.RGBA{0x87, 0xce, 0xeb, 0xff},
		"slateblue":            color.RGBA{0x6a, 0x5a, 0xcd, 0xff},
		"slategray":            color.RGBA{0x70, 0x80, 0x90, 0xff},
		"slategrey":            color.RGBA{0x70, 0x80, 0x90, 0xff},
		"snow":                 color.RGBA{0xff, 0xfa, 0xfa, 0xff},
		"springgreen":          color.RGBA{0x00, 0xff, 0x7f, 0xff},
		"steelblue":            color.RGBA{0x46, 0x82, 0xb4, 0xff},
		"tan":                  color.RGBA{0xd2, 0xb4, 0x8c, 0xff},
		"teal":                 color.RGBA{0x00, 0x80, 0x80, 0xff},
		"thistle":              color.RGBA{0xd8, 0xbf, 0xd8, 0xff},
		"tomato":               color.RGBA{0xff, 0x63, 0x47, 0xff},
		"turquoise":            color.RGBA{0x40, 0xe0, 0xd0, 0xff},
		"violet":               color.RGBA{0xee, 0x82, 0xee, 0xff},
		"wheat":                color.RGBA{0xf5, 0xde, 0xb3, 0xff},
		"white":                color.RGBA{0xff, 0xff, 0xff, 0xff},
		"whitesmoke":           color.RGBA{0xf5, 0xf5, 0xf5, 0xff},
		"yellow":               color.RGBA{0xff, 0xff, 0x00, 0xff},
		"yellowgreen":          color.RGBA{0x9a, 0xcd, 0x32, 0xff},
		"off":                  color.RGBA{0x00, 0x00, 0x00, 0xff},
	}
	if color, ok := colors[s]; ok {
		return color, nil
	} else {
		return nil, fmt.Errorf("Invalid color %q", s)
	}
}
