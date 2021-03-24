package main

type Car struct {
	owner string
	buyer string
	model string
	color string
	year uint16
	km uint32
	price float32
}

type Webmotors struct {
	carsToSell []*car
	carsSold []*car
}

func NewWebMotos(){
	return &webmotors{[],[]}
}

func StringfyCar(c *Car){
	
}

func (w *Webmotors) ListOwnerCarsToSell(ownerName string){

}

func (w *Webmotors) ListOwnerCarsSold(ownerName string){

}

func (w *Webmotors) ListBuyerCars(buyerName string){

}