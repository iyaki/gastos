package domain

import (
	"testing"
	"time"

	"github.com/iyaki/gastos/helpers"
)

func TestDepositar(t *testing.T) {
	t.Run("Camino feliz", func(t *testing.T) {
		cuenta := Cuenta{}

		cantidadADepositar := 10.0
		hoy := time.Now()

		err := cuenta.Depositar(cantidadADepositar, hoy)

		helpers.AssertNoError(t, err)

		balance := cuenta.Balance()
		helpers.AssertEquals(t, cantidadADepositar, balance)

		movimientos := cuenta.movimientos
		helpers.AssertEquals(t, 1, len(movimientos))
		helpers.AssertEquals(t, cantidadADepositar, movimientos[0].cantidad)
		helpers.AssertEquals(t, hoy, movimientos[0].fecha)
	})

	t.Run("Valores negativos", func(t *testing.T) {
		cuenta := Cuenta{}

		err := cuenta.Depositar(-10.0, time.Now())

		helpers.AssertError(t, ErrValorNoPositivo, err)

		balance := cuenta.Balance()
		helpers.AssertEquals(t, 0.0, balance)

		movimientos := cuenta.movimientos
		helpers.AssertEquals(t, 0, len(movimientos))
	})

	t.Run("Valor cero", func(t *testing.T) {
		cuenta := Cuenta{}

		err := cuenta.Depositar(0, time.Now())

		helpers.AssertError(t, ErrValorNoPositivo, err)

		balance := cuenta.Balance()
		helpers.AssertEquals(t, 0.0, balance)

		movimientos := cuenta.movimientos
		helpers.AssertEquals(t, 0, len(movimientos))
	})
}

func TestRetirar(t *testing.T) {
	cuenta := Cuenta{}

	hoy := time.Now()

	cuenta.Retirar(5, hoy)

	balance := cuenta.Balance()
	expected := -5.0
	helpers.AssertEquals(t, expected, balance)

	movimientos := cuenta.movimientos
	helpers.AssertEquals(t, 1, len(movimientos))
	helpers.AssertEquals(t, expected, movimientos[0].cantidad)
	helpers.AssertEquals(t, hoy, movimientos[0].fecha)

	t.Run("Camino feliz", func(t *testing.T) {
		cuenta := Cuenta{}

		hoy := time.Now()

		err := cuenta.Retirar(5, hoy)

		helpers.AssertNoError(t, err)

		balance := cuenta.Balance()
		expected := -5.0
		helpers.AssertEquals(t, expected, balance)

		movimientos := cuenta.movimientos
		helpers.AssertEquals(t, 1, len(movimientos))
		helpers.AssertEquals(t, expected, movimientos[0].cantidad)
		helpers.AssertEquals(t, hoy, movimientos[0].fecha)
	})

	t.Run("Valores negativos", func(t *testing.T) {
		cuenta := Cuenta{}

		err := cuenta.Retirar(-10.0, time.Now())

		helpers.AssertError(t, ErrValorNoPositivo, err)

		balance := cuenta.Balance()
		helpers.AssertEquals(t, 0.0, balance)

		movimientos := cuenta.movimientos
		helpers.AssertEquals(t, 0, len(movimientos))
	})

	t.Run("Valor cero", func(t *testing.T) {
		cuenta := Cuenta{}

		err := cuenta.Retirar(0, time.Now())

		helpers.AssertError(t, ErrValorNoPositivo, err)

		balance := cuenta.Balance()
		helpers.AssertEquals(t, 0.0, balance)

		movimientos := cuenta.movimientos
		helpers.AssertEquals(t, 0, len(movimientos))
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

	helpers.AssertEquals(t, 43.74, cuenta.Balance())
}
