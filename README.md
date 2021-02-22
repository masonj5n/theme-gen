# theme-gen
theme-gen takes a 7 color palette (with the first color being the background) from [coolors.co](https://coolors.io) and generates
a json file suitable for use with [pywal](https://github.com/dylanaraps/pywal), then applies that theme using `wal`.

## Example Usage
Head over to [coolors.co](https://coolors.io) and pick a 7 color palette. The first color will be the background, the subsequent 6 will represent color1
color2, color3, etc.

When you have your palette picked, just copy the end of the url that has the 7 colors in hex format separated by hyphens. For example, if I wanted
the color palette https://coolors.co/312f2f-84dccf-a6d9f7-bccce0-bf98a0-b48291-fa9f42, I would copy `312f2f-84dccf-a6d9f7-bccce0-bf98a0-b48291-fa9f42`
and use it as the value for the `-p` flag to `theme-gen`. Then use the -f flag to specify the name of the resulting theme file. If the -f flag isn't 
specified the name `default.json` will be used.

`theme-gen -p 312f2f-84dccf-a6d9f7-bccce0-bf98a0-b48291-fa9f42 -f myCoolTheme.json`

This command will generate your theme file and place it in the directory `~/.config/walthemes/`, and then apply the theme.

If you want to use a previously generated theme, just apply it with `wal`: `wal --theme ~/.config/walthemes/myCoolTheme.json`

If you use a light background color, use the flag -l to have a black foreground instead of white:

`theme-gen -p 312f2f-84dccf-a6d9f7-bccce0-bf98a0-b48291-fa9f42 -f myLightTheme.json -l`
