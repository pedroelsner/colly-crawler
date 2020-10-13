package app

import "github.com/pedroelsner/colly-crawler/internal/service/zukerman"

func Main() {
	config()
	migration()

	// Start
	zukerman.List("https://www.zukerman.com.br/leilao-de-casa")
	// zukerman.List("https://www.zukerman.com.br/leilao-de-apartamento")

	// Desativado
	//zukerman.Detail("https://www.zukerman.com.br/galpao-aclimacao-sao-paulo-sp-20157-148016")

	// Dormitórios e Garagem
	// https://www.zukerman.com.br/casa-jardim-tatiana-votorantim-sp-20472-149307
}
