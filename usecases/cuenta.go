package usecases

import (
	"time"

	"github.com/iyaki/gastos/domain"
)

// CrearCuenta crea una nueva cuenta sobre la cual realizar operaciones
func CrearCuenta(
	cuentaRepository domain.CuentaRepository,
	nombre string,
	cantidadInicial float64,
) (domain.Cuenta, error) {
	cuenta := domain.Cuenta{Nombre: nombre}

	var err error

	if cantidadInicial != 0 {
		err = cuenta.Depositar(cantidadInicial, time.Now())
		if err != nil {
			return domain.Cuenta{}, err
		}
	}

	return cuentaRepository.Agregar(cuenta)
}

// DepositarEnCuenta deposita una cierta cantidad de dinero en una cuenta
func DepositarEnCuenta(
	cuentaRepository domain.CuentaRepository,
	cuentaID domain.CuentaID,
	cantidad float64,
) (domain.Cuenta, error) {
	cuenta, err := cuentaRepository.Obtener(cuentaID)
	if err != nil {
		return cuenta, err
	}

	err = cuenta.Depositar(cantidad, time.Now())
	if err != nil {
		return cuenta, err
	}

	return cuentaRepository.Actualizar(cuenta)
}

// RetirarDeCuenta retira una cierta cantidad de dinero de una cuenta pudiendo
// dejar la cuenta con saldo negativo
func RetirarDeCuenta(
	cuentaRepository domain.CuentaRepository,
	cuentaID domain.CuentaID,
	cantidad float64,
) (domain.Cuenta, error) {
	cuenta, err := cuentaRepository.Obtener(cuentaID)
	if err != nil {
		return cuenta, err
	}

	err = cuenta.Retirar(cantidad, time.Now())
	if err != nil {
		return cuenta, err
	}

	return cuentaRepository.Actualizar(cuenta)
}

// EliminarCuenta elimina una cuenta de manera irreversible
func EliminarCuenta(
	cuentaRepository domain.CuentaRepository,
	cuentaID domain.CuentaID,
) (domain.Cuenta, error) {
	return cuentaRepository.Eliminar(cuentaID)
}
