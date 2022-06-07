package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscriber(t *testing.T) {
	// Os dados a serem usados no teste sao inicializados (entrada/saida)
	num1 := 5
	num2 := 3
	expectedResult := 2

	// o teste e executado
	result := calcSub(num1, num2)

	// os resultados sao validados
	assert.Equal(t, result, expectedResult, "devem ser iguais")
}
