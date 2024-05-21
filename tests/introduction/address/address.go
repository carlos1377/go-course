package address

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

// TipoDeEndereco verifica se o endereço tem um tipo valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {

		return cases.Title(language.BrazilianPortuguese, cases.NoLower).String(primeiraPalavraEndereco)
	}

	return "Tipo inválido"
}
