package webmotors

import (
	"fmt"
)
type Car struct {
	id uint32
	owner string
	buyer string
	model string
	color string
	year uint16
	km uint32
	price float32
}

type Webmotors struct {
	Cars []*Car
	UserInterface *Ui
}

func NewWebMotos() *Webmotors {
	w := Webmotors{make([]*Car, 0), &Ui{ClientStep: map[string]uint8{}}}
	return &w
}

func StringfyCar(c *Car) string{
	return " ### Id: " + fmt.Sprint(c.id) + " ### \nModel: " + c.model + "\nColor: " + c.color + "\nYear: " + fmt.Sprint(c.year) + "\nKm: " + fmt.Sprint(c.km) + "\nPrice: " + fmt.Sprintf("%f", c.price) + "\n"
}

func (w *Webmotors) ListOwnerCarsToSell(ownerName string) []*Car {
	response := make([]*Car, 0, len(w.Cars))
	for i := 0; i < len(w.Cars); i++ {
		if w.Cars[i].buyer == "" && w.Cars[i].owner == ownerName{
			response = append(response, w.Cars[i])
		}
	}
	return response
}

func (w *Webmotors) ListOwnerCarsSold(ownerName string) []*Car {
	response := make([]*Car, 0, len(w.Cars))
	for i := 0; i < len(w.Cars); i++ {
		if w.Cars[i].buyer != "" && w.Cars[i].owner == ownerName{
			response = append(response, w.Cars[i])
		}
	}
	return response
}

func (w *Webmotors) ListBuyerCars(buyerName string) []*Car {
	response := make([]*Car, 0, len(w.Cars))
	for i := 0; i < len(w.Cars); i++ {
		if w.Cars[i].buyer == buyerName{
			response = append(response, w.Cars[i])
		}
	}
	return response
}

func (w *Webmotors) AdvertiseCar(c *Car) {
	var higherId uint32 = 0
	for _, v := range w.Cars {
		if v.id > uint32(higherId){
			higherId = v.id
		}
	}
	c.id = higherId + 1
	w.Cars = append(w.Cars, c)
}

func(w *Webmotors) RemoveCarAd(cardId uint32) bool {
	removed := false
	for i, v := range w.Cars {
		if v.id == cardId {
			w.Cars = append(w.Cars[:i], w.Cars[i+1:]...)
			removed = true
			break
		}
	}
	return removed
}

func (w *Webmotors) BuyCar(carId uint32, buyerName string) bool {
	buyed := false
	for _, v := range w.Cars {
		if v.id == carId {
			if v.owner != buyerName {
				v.buyer = buyerName
				buyed = true
			}
		}
	}
	return buyed
}