package models

import (
	"fmt"
	"main/src/scenes"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"golang.org/x/exp/rand"
)

type Car struct {
	id    int
	timeInParking time.Duration
	image       *canvas.Image
	parkLot       int
}

func NewCar(id int) *Car {
	rand.Seed(uint64(time.Now().UnixNano()))
	timeInParking := time.Duration(rand.Intn(3)+3) * time.Second

	images := []string{
		"/Users/rextro/Documents/7mo/Concurrente/C2/ParkingSimulator/assets/car_1.png",
        "/Users/rextro/Documents/7mo/Concurrente/C2/ParkingSimulator/assets/car_2.png",
        "/Users/rextro/Documents/7mo/Concurrente/C2/ParkingSimulator/assets/car_3.png",
        "/Users/rextro/Documents/7mo/Concurrente/C2/ParkingSimulator/assets/car_4.png",
	}
	selectedImage := ChooseRandomImage(images)

	return &Car{
		id:          id,
		timeInParking : timeInParking,
		image : canvas.NewImageFromURI(storage.NewFileURI(selectedImage)),
	}
}

func ChooseRandomImage(images []string) string {
    rand.Seed(uint64(time.Now().UnixNano()))

    randomIndex := rand.Intn(len(images))
    
    return images[randomIndex]
}

func (c *Car) GetId() int {
	return c.id
}

func (c *Car) GetCarImage() *canvas.Image {
	return c.image
}

func (c *Car) TryPark(p *Parking, wg *sync.WaitGroup, s *scenes.Scene)  {
	var parkingSpaces = map[int][2]float32{
		0: {20, 10},
		1: {60, 10},
		2: {100, 10},
		3: {140, 10},
		4: {180, 10},
		5: {220, 10},
		6: {260, 10},
		7: {300, 10},
		8: {340, 10},
		9: {380, 10},
		10: {420, 10},
		11: {460, 10},
		12: {500, 10},
		13: {540, 10},
		14: {580, 10},
		15: {620, 10},
		16: {660, 10},
		17: {700, 10},
		18: {740, 10},
		19: {780, 10},
	}

    c.JoinPark(p, parkingSpaces)

	time.Sleep(c.timeInParking)

	c.LeavePark(p, s)

	wg.Done()
}

func (c *Car) JoinPark(p *Parking, parkingSpaces map[int][2]float32) {
	p.GetSpaces() <- c.GetId()
	p.GetEntrance().Lock()

	fmt.Printf("Auto %d ha entrado. Espacios ocupados: %d\n", c.GetId(), len(p.GetSpaces()))

	p.GetEntrance().Unlock()

	c.parkLot = p.FindAvailableSpace()
	p.OccupySpace(c.parkLot)
	

	pos := parkingSpaces[c.parkLot]

	c.image.Move(fyne.NewPos(pos[0], pos[1]))
}

func (c *Car) LeavePark(p *Parking, s *scenes.Scene) {
	p.GetEntrance().Lock()
	<-p.GetSpaces()

	p.FreeSpace(c.parkLot)
	c.image.Move(fyne.NewPos(425, 450))
	s.DeleteImage(c.image)

	fmt.Printf("Auto %d ha salido. Espacios ocupados: %d\n", c.GetId(), len(p.GetSpaces()))
	p.GetEntrance().Unlock()

}