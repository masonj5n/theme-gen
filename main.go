package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Palette represents the json formatted color palette
// expected by a pywal theme.
type Palette struct {
	Special map[string]string `json:"special"`
	Colors  map[string]string `json:"colors"`
}

var defaultPalette Palette = Palette{
	Special: map[string]string{"background": "#1F1F1F", "foreground": "#FFFFFF", "cursor": "#FFFFFF"},
	Colors: map[string]string{
		"color0":  "#1F1F1F",
		"color1":  "#C87004",
		"color10": "#FA900F",
		"color11": "#FA900F",
		"color12": "#796F67",
		"color13": "#FFFFFF",
		"color14": "#796F67",
		"color15": "#FFFFFF",
		"color2":  "#FA900F",
		"color3":  "#FA900F",
		"color4":  "#796F67",
		"color5":  "#FFFFFF",
		"color6":  "#796F67",
		"color7":  "#FFFFFF",
		"color8":  "#AAAAAA",
		"color9":  "#C87004"},
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

	filename := flag.String("f", "default.json", "Specify the name of the theme file (including .json at the end)")

	light := flag.Bool("l", false, "Specify that the palette has a light background, omit if using a dark background")

	flag.Parse()

	if *light {
		defaultPalette.Special["cursor"] = "#000000"
		defaultPalette.Special["foreground"] = "#000000"
		defaultPalette.Colors["color7"] = "#000000"
		defaultPalette.Colors["color15"] = "#000000"
	}

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
		return
	}

	err = writeThemeFile(jsonOutput, *filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating theme file:", err)
	}

	err = applyTheme(*filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error applying theme:", err)
	}

}

func writeThemeFile(jsonBytes []byte, filename string) error {

	homedir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error looking up home dir: %v", err)
	}
	fullThemePath := filepath.Join(homedir, ".config/walthemes")
	err = os.MkdirAll(fullThemePath, 0777)
	if err != nil {
		return fmt.Errorf("Error creating ~/.config/walthemes/ directory: %v", err)
	}

	err = os.WriteFile(filepath.Join(fullThemePath, filename), jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("Error creating/writing theme file: %v", err)
	}

	return nil
}

// applyTheme takes the filename of a theme file residing
// in .config/walthemes/ and runs the wal command to apply it.
func applyTheme(filename string) error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error getting home directory: %v", err)
	}

	fullFilePath := filepath.Join(homedir, ".config/walthemes", filename)

	walcmd := exec.Command("wal", "--theme", fullFilePath)

	walcmd.Stdout = os.Stdout
	walcmd.Stderr = os.Stderr

	err = walcmd.Run()
	if err != nil {
		return fmt.Errorf("Error running wal cmd: %v", err)
	}

	return nil
}
