package domain

import (
	"errors"
	"time"
)

// ErrValorNoPositivo es el error devuelto al intentar operar con valores no positivos.
var ErrValorNoPositivo = errors.New("No es posible operar con valores no positivos")

// Cuenta representa una cuenta en la que depositar y retirar dinero
type Cuenta struct {
	ID          int
	Nombre      string
	balance     float64
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
	fecha    time.Time
}

// CuentaRepository es la interfaz de los repositorios de cuenta.
type CuentaRepository interface {
	Agregar(cuenta Cuenta) (Cuenta, error)
	Actualizar(cuenta Cuenta) (Cuenta, error)
	Eliminar(cuentaID int) (Cuenta, error)
	Obtener(cuentaID int) (Cuenta, error)
	ObtenerTodas() []Cuenta
}

// ErrCuentaNuevaConID es el error devuelto por el repositorio al intentar agregar
// una cuenta con un ID repetido
var ErrCuentaNuevaConID = errors.New("No es posible agregar al repositorio cuentas que poseen un ID distinto de cero")

// ErrCuentaNoExiste es el error devuelto por el repositorio al intentar operar sobre
// una cuenta con un ID inexistente
var ErrCuentaNoExiste = errors.New("El ID de la cuenta que intenta eliminar o actualizar del repositorio no existe")

// ErrCuentaIDInvalido es el error devuelto por el repositorio al intentar realizar
// operaciones con un ID de cuenta menor o igual 0
var ErrCuentaIDInvalido = errors.New("El ID de la cuenta debe ser mayor a cero")
