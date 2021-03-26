package webmotors

import (
	"fmt"
	"strconv"
	"strings"
)

/*
*1Menu:
1_ Anunciar carro
2_ Listar ofertas disponiveis
3_ Listar anuncios proprias


4_ Retirar anuncio próprio
*2
5_ Listar carros vendidos
6_ Comprar carro

Anunciar carro:
*3Digite o modelo,cor,ano,quilometragem e preco, separados por ponto e virgula (ex: 'Ford Ka;Branco;1999;200000;35000,00')

Listar Ofertas disponiveis

Listar anuncios proprios

Listar Ofertas proprias

Listar Carros vendidos

*4Comprar Carro
Selecione o id do carro para comprar
*/

type Ui struct {
	ClientStep map[string]uint8
	WebmotorsApp *Webmotors
}

func (u *Ui) Receive(message string, name string) string {
	if _, exist := u.ClientStep[name]; !exist {
		u.ClientStep[name] = 0;
	}
	fmt.Printf("[UI Receive] Message received %s in step %d\n", message, u.ClientStep[name])
	menuMessage := "\n### Menu ###\n1_Anunciar carro\n2_Listar ofertas disponiveis\n3_Listar anuncios próprios\n4_Retirar anuncio próprio\n5_Comprar Carro\n6_Listar meus carros\nDigite uma opção: "
	// Menu
	if u.ClientStep[name] == 0{
		// Get message
		code, err := strconv.ParseInt(message, 10, 8)
		// User sent an option
		if err == nil && code > 0 && code < 7 {
			u.ClientStep[name] = uint8(code)
		} else {
			return menuMessage
		}
	}
	// Anunciar Carro (perguntar informacao)
	if(u.ClientStep[name] == 1){
		u.ClientStep[name] = 7
		return "Digite o modelo,cor,ano,quilometragem e preco, separados por ponto e virgula (ex: 'Ford Ka;Branco;1999;200000;35000,00') \n> "
	}
	// Listar ofertas disponíveis
	if(u.ClientStep[name] == 2){
		cars := u.WebmotorsApp.Cars
		response := ""
		for _, v := range cars {
			if v.buyer == "" {
				response = response + "\n" + StringfyCar(v)
			}
		}
		u.ClientStep[name] = 0 // Return to menu step
		if len(response) == 0 {
			response = "\nNenhum carro a venda\n"
		}
		return response + menuMessage
		
	}
	// Listar anuncios próprios
	if(u.ClientStep[name] == 3){
		cars := u.WebmotorsApp.ListOwnerCarsToSell(name)
		response := ""
		for _, v := range cars {
			response = response + "\n" + StringfyCar(v)
		}
		u.ClientStep[name] = 0 // Return to menu step
		if len(response) == 0 {
			response = "\nNenhum carro a venda\n"
		}
		return response + menuMessage

	}
	// Retirar anuncio próprio (perguntar informacao)
	if(u.ClientStep[name] == 4){
		u.ClientStep[name] = 8
		return "Qual o id do veiculo para retirar o anuncio? "
	}
	// Comprar carro (perguntar informacao)
	if(u.ClientStep[name] == 5){
		u.ClientStep[name] = 9
		return "Qal o id do veiculo para comprar? "
	}
	// Listar carros proprios
	if(u.ClientStep[name] == 6) {
		cars := u.WebmotorsApp.ListOwnerCarsToSell(name)
		response := ""
		for _, v := range cars {
			response = response + "\n" + StringfyCar(v)
		}
		cars = u.WebmotorsApp.ListBuyerCars(name)
		for _, v := range cars {
			response = response + "\n" + StringfyCar(v)
		}
		u.ClientStep[name] = 0 // Return to menu step
		if len(response) == 0 {
			response = "\nNenhum carro a venda\n"
		}
		return response + menuMessage
	}
	// Anunciar carro
	if(u.ClientStep[name] == 7){
		if numberOfSeparators := strings.Count(message, ";"); numberOfSeparators != 4 {
			return "### ERRO Dados para carro inválidos ### \n" + menuMessage
		}
		data := strings.Split(message,  ";")
		year, err := strconv.ParseInt(data[2], 10, 16)
		if err != nil {
			return "### ERRO Campo ano ###" + menuMessage
		}
		km, err := strconv.ParseInt(data[3], 10, 32)
		if err != nil {
			return "### ERRO Campo quilometragem ###" + menuMessage
		}
		price, err := strconv.ParseFloat(strings.ReplaceAll(data[4], ",", "."), 32)
		if err != nil {
			return "### ERRO Campo preço ###" + menuMessage
		}

		u.WebmotorsApp.AdvertiseCar(&Car{
			owner: name,
			buyer: "",
			model: data[0],
			color: data[1],
			year: uint16(year),
			km: uint32(km),
			price: float32(price),
		})
		u.ClientStep[name] = 0
		return "### Carro anunciado com sucesso ### \n\n" + menuMessage 
	}
	// Retirar anuncio proprio
	if(u.ClientStep[name] == 8){
		// Get message
		code, err := strconv.ParseInt(message, 10, 8)
		u.ClientStep[name] = 0
		// User sent an option
		if err != nil {
			return fmt.Sprintf("ERRO com id %s\n", err)
		}
		fmt.Printf("Removendo carro do id %d\n", code)
		removed := u.WebmotorsApp.RemoveCarAd(uint32(code))
		if removed {
			return "### Removido com sucesso ###\n\n"+ menuMessage
		}
		return "### ERRO Id inexistente ### \n\n" + menuMessage
	}
	// Comprar carro
	if(u.ClientStep[name] == 9){
		// Get message
		code, err := strconv.ParseInt(message, 10, 8)
		u.ClientStep[name] = 0
		// User sent an option
		if err != nil {
			return "### ERRO com input ### \n" + menuMessage
		}
		removed := u.WebmotorsApp.BuyCar(uint32(code), name)
		if removed {
			return "### Comprado com sucesso ###\n\n"+ menuMessage
		}
		return "### ERRO Id ou tentativa de compra do proprio carro ### \n\n" + menuMessage
	}
	return ""
}