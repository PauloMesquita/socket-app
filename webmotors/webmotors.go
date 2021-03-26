package webmotors

import "fmt"

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
	CarsToSell []*Car
	CarsSold []*Car
}

func NewWebMotos() *Webmotors {
	w := Webmotors{[]*Car{},[]*Car{}}
	return &w
}

func StringfyCar(c *Car) string{
	return "Model: " + c.model + "\nColor: " + c.color + "\nYear: " + fmt.Sprint(c.year) + "\nKm: " + fmt.Sprint(c.km) + "\nPrice: " + fmt.Sprintf("%f", c.price) + "\n"
}

func (w *Webmotors) ListOwnerCarsToSell(ownerName string) []*Car {
	response := make([]*Car, 0, len(w.CarsToSell))
	for i := 0; i < len(w.CarsToSell); i++ {
		if w.CarsToSell[i].owner == ownerName{
			response = append(response, w.CarsToSell[i])
		}
	}
	return response
}

func (w *Webmotors) ListOwnerCarsSold(ownerName string) []*Car {
	response := make([]*Car, 0, len(w.CarsSold))
	for i := 0; i < len(w.CarsSold); i++ {
		if w.CarsSold[i].owner == ownerName{
			response = append(response, w.CarsSold[i])
		}
	}
	return response
}

func (w *Webmotors) ListBuyerCars(buyerName string) []*Car {
	response := make([]*Car, 0, len(w.CarsSold))
	for i := 0; i < len(w.CarsSold); i++ {
		if w.CarsSold[i].buyer == buyerName{
			response = append(response, w.CarsSold[i])
		}
	}
	return response
}