package vips

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestI_GIF_Animated_Pages(t *testing.T) {
	Startup(nil)
	image, err := NewImageFromFile(resources + "gif-animated.gif")
	require.NoError(t, err)

	pages := image.Pages()
	assert.Equal(t, 8, pages)
}

func TestImage_GIF_Animated(t *testing.T) {
	goldenAnimatedTest(t, resources+"gif-animated.gif",
		-1,
		nil,
		nil,
		exportGif(NewGifExportParams()))
}

func TestImage_GIF_Animated_ExportNative(t *testing.T) {
	goldenAnimatedTest(t, resources+"gif-animated.gif",
		3,
		nil,
		nil,
		nil)
}

func TestImage_GIF_Animated_to_WebP(t *testing.T) {
	goldenAnimatedTest(t, resources+"gif-animated.gif",
		3,
		nil,
		nil,
		exportWebp(NewWebpExportParams()))
}

func TestImage_GIF_Animated_Resize(t *testing.T) {
	goldenAnimatedTest(t, resources+"gif-animated.gif",
		3,
		func(img *ImageRef) error {
			return img.Resize(2, KernelCubic)
		},
		nil,
		nil)
}

func TestImage_GIF_Animated_Rotate90(t *testing.T) {
	goldenAnimatedTest(t, resources+"gif-animated.gif",
		-1,
		func(img *ImageRef) error {
			return img.Rotate(Angle90)
		},
		nil,
		nil)
}

func TestImage_GIF_Animated_Rotate270(t *testing.T) {
	goldenAnimatedTest(t, resources+"gif-animated.gif",
		-1,
		func(img *ImageRef) error {
			return img.Rotate(Angle270)
		},
		nil,
		nil)
}
