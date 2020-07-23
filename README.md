# go-palette2048_tia

Package **palette2048_tia** provides the TIA color palette in the form of 256 RGBA color with 8-bits per color channel,
that when used with the https://github.com/reiver/go-palette2048 package works with Go's built-in `"image"`, `"image/color"`, and `"image/draw"` packages.

## Documention

Online documentation, which includes examples, can be found at: https://github.com/reiver/go-palette2048_tia

[![GoDoc](https://godoc.org/github.com/reiver/go-palette2048_tia?status.svg)](https://godoc.org/github.com/reiver/go-palette2048_tia)

## Example

```go
import (
	"github.com/reiver/go-palette2048"
	"github.com/reiver/go-palette2048_tia"
)

// ...

palette := palette2048.Slice(palette2048_tia.Palette)

// ...

// Get the color at a specific index in the palette.
color := palette.Color(index)

// ...

// Get the color in the palette that is closet to a reference color.
closestColor := palette.Convert(referenceColor)
```

## See Also

Also see:

* https://github.com/reiver/go-palette2048
