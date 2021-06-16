package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Erro o esperado era %d e deu %d", 30, total)
	}
}
