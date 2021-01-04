package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado foi %v"

func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperando := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)
	if valor != valorEsperando {
		t.Errorf(erroPadrao, valorEsperando, valor)
	}
}
