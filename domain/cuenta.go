package cuenta

import (
	"errors"
	"time"
)

// ErrValorNoPositivo es el error devuelto al intentar operar con valores no positivos.
var ErrValorNoPositivo = errors.New("No es posible operar con valores no positivos")

// Cuenta representa una cuenta en la que depositar y retirar dinero
type Cuenta struct {
	balance float64
	movimientos []movimiento
}

// Depositar permite depositar una cierta cantidad de dinero en una cuenta.
// Si se intenta invocar esta funcion con un valor no positivo como primer argumento
// recibirá el error ErrValorNoPositivo
func (c *Cuenta) Depositar(cantidad float64, fecha time.Time) error {
	if cantidad <= 0 {
		return ErrValorNoPositivo
	}

	c.agregarMovimiento(cantidad, fecha)
	return nil
}

// Retirar permite retirar una cierta cantidad de dinero de una cuenta.
// Si se intenta invocar esta funcion con un valor no positivo como primer argumento
// recibirá el error ErrValorNoPositivo
func (c *Cuenta) Retirar(cantidad float64, fecha time.Time) error {
	if cantidad <= 0 {
		return ErrValorNoPositivo
	}

	c.agregarMovimiento(-cantidad, fecha)
	return nil
}

func (c *Cuenta) agregarMovimiento(cantidad float64, fecha time.Time) {
	c.movimientos = append(c.movimientos, movimiento{cantidad: cantidad, fecha: fecha})
}

// Balance devuelve el balance de la cuenta basado en sus movimientos.
func (c *Cuenta) Balance() float64 {
	var balance float64
	for _, movimiento := range c.movimientos {
		balance += movimiento.cantidad
	}
	return balance
}

type movimiento struct {
	cantidad float64
	fecha time.Time
}
