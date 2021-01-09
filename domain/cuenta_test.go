package cuenta

import (
	"testing"
	"time"
)

func TestDepositar(t *testing.T) {
	t.Run("Camino feliz", func(t *testing.T) {
		cuenta := Cuenta{}

		cantidadADepositar := 10.0
		hoy := time.Now()

		err := cuenta.Depositar(cantidadADepositar, hoy)

		assertNoError(t, err)

		balance := cuenta.Balance()
		assertEquals(t, cantidadADepositar, balance)

		movimientos := cuenta.movimientos
		assertEquals(t, 1, len(movimientos))
		assertEquals(t, cantidadADepositar, movimientos[0].cantidad)
		assertEquals(t, hoy, movimientos[0].fecha)
	})

	t.Run("Valores negativos", func(t *testing.T) {
		cuenta := Cuenta{}

		err := cuenta.Depositar(-10.0, time.Now())

		assertError(t, ErrValorNoPositivo, err)

		balance := cuenta.Balance()
		assertEquals(t, 0.0, balance)

		movimientos := cuenta.movimientos
		assertEquals(t, 0, len(movimientos))
	})

	t.Run("Valor cero", func(t *testing.T) {
		cuenta := Cuenta{}

		err := cuenta.Depositar(0, time.Now())

		assertError(t, ErrValorNoPositivo, err)

		balance := cuenta.Balance()
		assertEquals(t, 0.0, balance)

		movimientos := cuenta.movimientos
		assertEquals(t, 0, len(movimientos))
	})
}

func TestRetirar(t *testing.T) {
	cuenta := Cuenta{}

	hoy := time.Now()

	cuenta.Retirar(5, hoy)

	balance := cuenta.Balance()
	expected := -5.0
	assertEquals(t, expected, balance)

	movimientos := cuenta.movimientos
	assertEquals(t, 1, len(movimientos))
	assertEquals(t, expected, movimientos[0].cantidad)
	assertEquals(t, hoy, movimientos[0].fecha)

	t.Run("Camino feliz", func(t *testing.T) {
		cuenta := Cuenta{}

		hoy := time.Now()

		err := cuenta.Retirar(5, hoy)

		assertNoError(t, err)

		balance := cuenta.Balance()
		expected := -5.0
		assertEquals(t, expected, balance)

		movimientos := cuenta.movimientos
		assertEquals(t, 1, len(movimientos))
		assertEquals(t, expected, movimientos[0].cantidad)
		assertEquals(t, hoy, movimientos[0].fecha)
	})

	t.Run("Valores negativos", func(t *testing.T) {
		cuenta := Cuenta{}

		err := cuenta.Retirar(-10.0, time.Now())

		assertError(t, ErrValorNoPositivo, err)

		balance := cuenta.Balance()
		assertEquals(t, 0.0, balance)

		movimientos := cuenta.movimientos
		assertEquals(t, 0, len(movimientos))
	})

	t.Run("Valor cero", func(t *testing.T) {
		cuenta := Cuenta{}

		err := cuenta.Retirar(0, time.Now())

		assertError(t, ErrValorNoPositivo, err)

		balance := cuenta.Balance()
		assertEquals(t, 0.0, balance)

		movimientos := cuenta.movimientos
		assertEquals(t, 0, len(movimientos))
	})
}

func TestBalance(t *testing.T) {
	cuenta := Cuenta{}

	cuenta.Depositar(5.0, time.Now())
	cuenta.Depositar(10.25, time.Now())
	cuenta.Depositar(100.0, time.Now())
	cuenta.Retirar(75.5, time.Now())
	cuenta.Depositar(5.99, time.Now())
	cuenta.Retirar(2.0, time.Now())

	assertEquals(t, 43.74, cuenta.Balance())
}

func assertEquals(t *testing.T, expected, got interface{}) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("error not expected but %v", got)
	}
}

func assertError(t *testing.T, expected, got error) {
	t.Helper()

	if got == nil {
		t.Errorf("an error was expected but none given")
	}

	if expected != got {
		t.Errorf("expected %v got %v", expected, got)
	}
}
