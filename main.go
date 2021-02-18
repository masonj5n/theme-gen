package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Palette represents the json formatted color palette
// expected by a pywal theme.
type Palette struct {
	Special map[string]string `json:"special"`
	Colors  map[string]string `json:"colors"`
}

var defaultPalette Palette = Palette{
	Special: map[string]string{"background": "#283618", "foreground": "#FFFFFF", "cursor": "#FFFFFF"},
	Colors: map[string]string{
		"color0":  "#283618",
		"color1":  "#AA6122",
		"color2":  "#9AAB5F",
		"color3":  "#606C38",
		"color4":  "#DDA15E",
		"color5":  "#BC6C25",
		"color6":  "#C27939",
		"color7":  "#FFFFFF",
		"color8":  "#AAAAAA",
		"color9":  "#AA6122",
		"color10": "#9AAB5F",
		"color11": "#606C38",
		"color12": "#DDA15E",
		"color13": "#BC6C25",
		"color14": "#C27939",
		"color15": "#FFFFFF"},
}

func main() {

	var hexColors []string
	flag.Func("p", "Palette string from coolors.co", func(p string) error {
		hexColors = strings.Split(p, "-")
		if len(hexColors) != 7 {
			return fmt.Errorf("Did not provide 7 colors")
		}
		for i := range hexColors {
			hexColors[i] = strings.Join([]string{"#", strings.ToUpper(hexColors[i])}, "")
		}
		return nil
	})

	flag.Parse()
	if len(hexColors) > 0 {
		defaultPalette.Special["background"] = hexColors[0]
		defaultPalette.Colors["color0"] = hexColors[0]

		for i := 1; i < len(hexColors); i++ {
			defaultPalette.Colors[fmt.Sprintf("color%d", i)] = hexColors[i]
			defaultPalette.Colors[fmt.Sprintf("color%d", i+8)] = hexColors[i]
		}
	}

	jsonOutput, err := json.Marshal(defaultPalette)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling palette into json")
	}

	fmt.Println(string(jsonOutput))
}
