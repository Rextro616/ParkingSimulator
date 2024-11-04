package scenes

import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Scene struct {
	scene fyne.Window
	container *fyne.Container
}

func NewScene(scene fyne.Window) *Scene {
	return &Scene{scene: scene, container: nil}
}

func (s *Scene) Init () {
	rect := canvas.NewRectangle(color.RGBA{R: 50, G: 50, B: 50, A: 255})
	rect.Resize(fyne.NewSize(900, 600))
	rect.Move(fyne.NewPos(0, 0))

	s.container = container.NewWithoutLayout(rect)
	s.scene.SetContent(s.container)
	
	border := canvas.NewRectangle(color.White)
	border.Resize(fyne.NewSize(900, 10))
	border.Move(fyne.NewPos(0, 0))

	s.container.Add(border)

	borderDownRignt := canvas.NewRectangle(color.White)
	borderDownRignt.Resize(fyne.NewSize(400, 10))
	borderDownRignt.Move(fyne.NewPos(500, 500))

	s.container.Add(borderDownRignt)

	borderDownLeft := canvas.NewRectangle(color.White)
	borderDownLeft.Resize(fyne.NewSize(400, 10))
	borderDownLeft.Move(fyne.NewPos(0, 500))

	s.container.Add(borderDownLeft)

	borderInit := 20
	
	for i := 1; i < 22; i++ {
		lineParkingSpace := canvas.NewRectangle(color.RGBA{R: 243, G: 233, B: 24, A: 255})
		lineParkingSpace.Resize(fyne.NewSize(3, 60))
		lineParkingSpace.Move(fyne.NewPos(float32(borderInit), 10))
		s.container.Add(lineParkingSpace) 
		borderInit = borderInit + 42
	}
}

func (s *Scene) AddImage(image *canvas.Image, posX, posY float32) {
	image.FillMode = canvas.ImageFill(canvas.ImageScaleFastest)
	image.Resize(fyne.NewSize(60, 60)) 
	image.Move(fyne.NewPos(posX, posY)) 

	s.container.Add(image)
	s.container.Refresh()
}

func (s *Scene) DeleteImage(image *canvas.Image) {
	s.container.Remove(image)
	s.container.Refresh()
}

func (s *Scene) Refresh() {
	s.container.Refresh()
}

func (s *Scene) GetContainer() *fyne.Container {
	return s.container
}