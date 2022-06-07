package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrd(t *testing.T) {
	// Os dados a serem usados no teste sao inicializados (entrada/saida)
	num1 := []int{50, 5, 10, 20}
	expectedResult := []int{5, 10, 20, 50}

	// o teste e executado
	result := sortedNum(num1[:])

	// os resultados sao validados
	assert.Equal(t, result, expectedResult)
}
