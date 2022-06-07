package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	// Os dados a serem usados no teste sao inicializados (entrada/saida)
	num1 := 10
	num2 := 0
	expectedResult := fmt.Errorf("o denominador nao pode ser 0")

	// o teste e executado
	_, err := divide(num1, num2)

	// os resultados sao validados
	assert.Equal(t, err, expectedResult)
}
