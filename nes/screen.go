package gones

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/tiagolobocastro/gones/nes/common"
	"image/color"
	"os"
	"runtime"
	"time"
)

type screen struct {
	nes *nes

	// window where we draw the sprite
	window *pixelgl.Window

	// front and back buffers
	buffer0 *pixel.PictureData
	buffer1 *pixel.PictureData
	sprite  *pixel.Sprite

	framebuffer common.Framebuffer

	// FPS stats
	fpsChannel   <-chan time.Time
	fpsLastFrame int
}

func (s *screen) init(nes *nes) {
	s.nes = nes

	s.setSprite()
}

func (s *screen) run() {
	go func() {
		runtime.LockOSThread()
		pixelgl.Run(s.runThread)
		os.Exit(0)
	}()
}

func (s *screen) runThread() {
	cfg := pixelgl.WindowConfig{
		Title:  "GoNes",
		Bounds: pixel.R(0, 0, screenXWidth, screenYHeight),
		VSync:  true,
	}
	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	s.window = window
	s.fpsChannel = time.Tick(time.Second)
	s.fpsLastFrame = 0

	s.runner()
}

func (s *screen) runner() {
	lastLoopFrames := 0
	for !s.window.Closed() {

		<-s.framebuffer.FrameUpdated

		frameDiff := s.framebuffer.Frames - lastLoopFrames
		if frameDiff > 0 {
			if frameDiff > 1 {
				fmt.Printf("Oops, skipped %v frames!\n", frameDiff)
			}

			s.draw()
			s.window.Update()
			lastLoopFrames = s.framebuffer.Frames
		}

		s.updateFpsTitle()
		s.updateControllers()
	}
}

var buttons = [8]struct {
	id  uint8
	key pixelgl.Button
}{
	{bitA, pixelgl.KeyS},
	{bitB, pixelgl.KeyA},
	{bitSelect, pixelgl.KeyLeftShift},
	{bitStart, pixelgl.KeyEnter},
	{bitUp, pixelgl.KeyUp},
	{bitDown, pixelgl.KeyDown},
	{bitLeft, pixelgl.KeyLeft},
	{bitRight, pixelgl.KeyRight},
}

func (s *screen) updateControllers() {
	onePressed := false
	for _, button := range buttons {
		pressed := s.window.Pressed(button.key)
		s.nes.ctrl.poke(0, button.id, pressed)
		if pressed {
			onePressed = true
		}
	}

	if s.window.Pressed(pixelgl.KeyLeftControl) && s.window.Pressed(pixelgl.KeyR) {
		s.nes.Reset()
		onePressed = true
	}

	if onePressed {
		s.window.UpdateInput()
	}
}

func (s *screen) updateFpsTitle() {
	select {
	case <-s.fpsChannel:
		frames := s.framebuffer.Frames - s.fpsLastFrame
		s.fpsLastFrame = s.framebuffer.Frames

		s.window.SetTitle(fmt.Sprintf("GoNes | FPS: %d", frames))
	default:
	}
}

func (s *screen) draw() {
	// seems to be required, for reasons unknown
	s.updateSprite()

	s.sprite.Draw(s.window, pixel.IM.Moved(s.window.Bounds().Center()).ScaledXY(s.window.Bounds().Center(), pixel.V(3, 3)))
}

func (s *screen) updateSprite() {
	if s.framebuffer.FrameIndex == 1 {
		// ppu is drawing new pixels on buffer1, which means the stable data is in buffer0
		s.sprite = pixel.NewSprite(s.buffer0, pixel.R(0, 0, frameXWidth, frameYHeight))
	} else {
		s.sprite = pixel.NewSprite(s.buffer1, pixel.R(0, 0, frameXWidth, frameYHeight))
	}
}

func (s *screen) setSprite() {

	s.buffer0 = &pixel.PictureData{
		Pix:    make([]color.RGBA, frameXWidth*frameYHeight),
		Stride: frameXWidth,
		Rect:   pixel.R(0, 0, frameXWidth, frameYHeight),
	}

	s.buffer1 = &pixel.PictureData{
		Pix:    make([]color.RGBA, frameXWidth*frameYHeight),
		Stride: frameXWidth,
		Rect:   pixel.R(0, 0, frameXWidth, frameYHeight),
	}

	s.framebuffer = common.Framebuffer{
		Buffer0:      s.buffer0.Pix,
		Buffer1:      s.buffer1.Pix,
		FrameIndex:   0,
		FrameUpdated: make(chan bool),
		Frames:       0,
	}

	s.updateSprite()
}
