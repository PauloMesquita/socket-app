package webmotors

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
	step uint8
}

func (u *Ui) receive(message string) string {
	if u.step == 0 {
		return "Menu\n1_Anunciar carro\n2_Listar ofertas disponiveis\n3_Listar anuncios próprios\n4_Retirar anuncio próprio\n5_Comprar Carro\nDigite uma opção: "
	}
	return ""
}