package cuenta

import (
	"testing"
	"time"
)

func TestDepositar(t *testing.T) {
	cuenta := Cuenta{}

	cantidadADepositar := 10.0
	hoy := time.Now()

	cuenta.Depositar(cantidadADepositar, hoy)

	balance := cuenta.Balance()
	asserEquals(t, cantidadADepositar, balance)
	
	movimientos := cuenta.movimientos
	asserEquals(t, 1, len(movimientos))
	asserEquals(t, cantidadADepositar, movimientos[0].cantidad)
	asserEquals(t, hoy, movimientos[0].fecha)
}

func TestRetirar(t *testing.T) {
	cuenta := Cuenta{}

	hoy := time.Now()

	cuenta.Retirar(5, hoy)

	balance := cuenta.Balance()
	expected := -5.0
	asserEquals(t, expected, balance)
	
	movimientos := cuenta.movimientos
	asserEquals(t, 1, len(movimientos))
	asserEquals(t, expected, movimientos[0].cantidad)
	asserEquals(t, hoy, movimientos[0].fecha)
}

func asserEquals(t *testing.T, expected, got interface{}) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %v got %v", expected, got)
	}
}
