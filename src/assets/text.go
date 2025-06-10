package assets

import (
	"bytes"
	"os"

	"github.com/go-text/typesetting/font"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var textFont font.Face

var Fonts = make(map[string]*text.GoTextFaceSource)

func AddFont(fontName string, fontPath string) error {
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		return err
	}

	ft, err := text.NewGoTextFaceSource(bytes.NewReader(fontData))
	if err != nil {
		return err
	}

	Fonts[fontName] = ft
	return nil
}
