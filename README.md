# theme-gen
theme-gen takes a 7 color palette (with the first color being the background) from [coolors.co](https://coolors.io) and generates
a json file suitable for use with [pywal](https://github.com/dylanaraps/pywal).

## Example Usage
Head over to [coolors.co](https://coolors.io) and pick a 7 color palette. The first color will be the background, the subsequent 6 will represent color1
color2, color3, etc.

When you have your palette picked, just copy the end of the url that has the 7 colors in hex format separated by hyphens. For example, if I wanted
the color palette https://coolors.co/312f2f-84dccf-a6d9f7-bccce0-bf98a0-b48291-fa9f42, I would copy `312f2f-84dccf-a6d9f7-bccce0-bf98a0-b48291-fa9f42`
and use it as the value for the `-p` flag to `gen-theme`.

`theme-gen -p 312f2f-84dccf-a6d9f7-bccce0-bf98a0-b48291-fa9f42`

The output goes to stdout, so you may want to append it directly to a file like so:

`theme-gen -p 312f2f-84dccf-a6d9f7-bccce0-bf98a0-b48291-fa9f42 > mytheme.json`

Once you have the json file, just run:

`wal --theme mytheme.json`
