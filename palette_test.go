package palette2048_tia_test

import (
	"github.com/reiver/go-palette2048"
	"github.com/reiver/go-palette2048_tia"

	"fmt"
	"image/color"

	"testing"
)

func TestPalette_palette2048(t *testing.T) {

	var palette palette2048.Slice = palette2048.Slice(palette2048_tia.Palette[:])

	if nil == palette {
		t.Error("This should never happen.")
		return
	}

	{
		var yellow color.Color = color.RGBA{255,199,6,55}

		actualColor := palette.Convert(yellow)
		actual := fmt.Sprint(actualColor)

		expected := "rgba(208,208,80,255)" // #d0d050

		if expected != actual {
			t.Errorf("The actual color was not what was expected.")
			t.Log("EXPECTED:", expected)
			t.Log("ACTUAL:  ", actual)
		}
	}
}
