package usecases

import (
	"testing"

	"github.com/iyaki/gastos/domain"
	"github.com/iyaki/gastos/helpers"
	"github.com/iyaki/gastos/usecases/internals"
)

func TestCrearCuenta(t *testing.T) {
	t.Run("Camino feliz", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		nombre := "Cuenta de prueba"
		cantidadInicial := 0.0
		cuenta, err := CrearCuenta(
			cuentaRepository,
			nombre,
			cantidadInicial,
		)

		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, nombre, cuenta.Nombre)
		helpers.AssertEquals(t, cantidadInicial, cuenta.Balance())
		helpers.AssertEquals(t, 1, len(cuentaRepository.ObtenerTodas()))
		cuentaAlmacenada, err := cuentaRepository.Obtener(1)
		helpers.AssertNoError(t, err)
		helpers.AssertDeepEquals(t, cuenta, cuentaAlmacenada)
	})

	t.Run("Camino feliz con cantidad inicial", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		nombre := "Cuenta de prueba"
		cantidadInicial := 5.0
		cuenta, err := CrearCuenta(
			cuentaRepository,
			nombre,
			cantidadInicial,
		)

		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, nombre, cuenta.Nombre)
		helpers.AssertEquals(t, cantidadInicial, cuenta.Balance())
		helpers.AssertEquals(t, 1, len(cuentaRepository.ObtenerTodas()))
		cuentaAlmacenada, err := cuentaRepository.Obtener(1)
		helpers.AssertNoError(t, err)
		helpers.AssertDeepEquals(t, cuenta, cuentaAlmacenada)
	})

	t.Run("Cantidad inicial negativa", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		nombre := "Cuenta de prueba"
		cuenta, err := CrearCuenta(
			cuentaRepository,
			nombre,
			-5.0,
		)

		helpers.AssertError(t, err, domain.ErrValorNoPositivo)
		helpers.AssertEquals(t, "", cuenta.Nombre)
		helpers.AssertEquals(t, 0.0, cuenta.Balance())
		helpers.AssertEquals(t, 0, len(cuentaRepository.ObtenerTodas()))
	})
}

func TestDepositarEnCuenta(t *testing.T) {
	t.Run("Camino feliz", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		cantidadInicial := 5.0
		nuevaCuenta, _ := CrearCuenta(
			cuentaRepository,
			"",
			cantidadInicial,
		)

		cantidad := 10.5
		cuenta, err := DepositarEnCuenta(
			cuentaRepository,
			nuevaCuenta.ID,
			cantidad,
		)

		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, cantidadInicial, nuevaCuenta.Balance())
		helpers.AssertEquals(t, cantidadInicial+cantidad, cuenta.Balance())
		helpers.AssertEquals(t, 1, len(cuentaRepository.ObtenerTodas()))
		cuentaAlmacenada, err := cuentaRepository.Obtener(1)
		helpers.AssertDeepEquals(t, cuenta, cuentaAlmacenada)
	})

	t.Run("Cuenta con id cero", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		cuenta, err := DepositarEnCuenta(
			cuentaRepository,
			0,
			10.0,
		)

		helpers.AssertError(t, err, domain.ErrCuentaIDInvalido)
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuenta)
		helpers.AssertDeepEquals(t, 0, len(cuentaRepository.ObtenerTodas()))
	})

	t.Run("Cuenta inexistente", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		cuenta, err := DepositarEnCuenta(
			cuentaRepository,
			5,
			10.0,
		)

		helpers.AssertError(t, err, domain.ErrCuentaNoExiste)
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuenta)
		helpers.AssertDeepEquals(t, 0, len(cuentaRepository.ObtenerTodas()))
	})

	t.Run("Depositar cantidad cero", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		nuevaCuenta, _ := CrearCuenta(cuentaRepository, "", 5.0)

		cuenta, err := DepositarEnCuenta(
			cuentaRepository,
			nuevaCuenta.ID,
			0.0,
		)

		helpers.AssertError(t, domain.ErrValorNoPositivo, err)
		helpers.AssertEquals(t, nuevaCuenta.Balance(), cuenta.Balance())
		cuentaAlmacenada, _ := cuentaRepository.Obtener(1)
		helpers.AssertEquals(t, cuentaAlmacenada.Balance(), cuenta.Balance())
		helpers.AssertDeepEquals(t, 1, len(cuentaRepository.ObtenerTodas()))
	})

	t.Run("Depositar cantidad negativa", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		nuevaCuenta, _ := CrearCuenta(cuentaRepository, "", 5.0)

		cuenta, err := DepositarEnCuenta(
			cuentaRepository,
			nuevaCuenta.ID,
			-6.5,
		)

		helpers.AssertError(t, domain.ErrValorNoPositivo, err)
		helpers.AssertEquals(t, nuevaCuenta.Balance(), cuenta.Balance())
		cuentaAlmacenada, _ := cuentaRepository.Obtener(1)
		helpers.AssertEquals(t, cuentaAlmacenada.Balance(), cuenta.Balance())
		helpers.AssertDeepEquals(t, 1, len(cuentaRepository.ObtenerTodas()))
	})
}

func TestRetirarDeCuenta(t *testing.T) {
	t.Run("Camino feliz", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		cantidadInicial := 10.0
		nuevaCuenta, _ := CrearCuenta(
			cuentaRepository,
			"Cuenta de prueba",
			cantidadInicial,
		)

		cantidad := 6.5
		cuenta, err := RetirarDeCuenta(
			cuentaRepository,
			nuevaCuenta.ID,
			cantidad,
		)

		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, cantidadInicial, nuevaCuenta.Balance())
		helpers.AssertEquals(t, cantidadInicial-cantidad, cuenta.Balance())
		helpers.AssertEquals(t, 1, len(cuentaRepository.ObtenerTodas()))
		cuentaAlmacenada, err := cuentaRepository.Obtener(1)
		helpers.AssertDeepEquals(t, cuenta, cuentaAlmacenada)
	})

	t.Run("Cuenta con id cero", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		cuenta, err := RetirarDeCuenta(
			cuentaRepository,
			0,
			10.0,
		)

		helpers.AssertError(t, err, domain.ErrCuentaIDInvalido)
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuenta)
		helpers.AssertDeepEquals(t, 0, len(cuentaRepository.ObtenerTodas()))
	})

	t.Run("Cuenta inexistente", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		cuenta, err := RetirarDeCuenta(
			cuentaRepository,
			5,
			10.0,
		)

		helpers.AssertError(t, err, domain.ErrCuentaNoExiste)
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuenta)
		helpers.AssertDeepEquals(t, 0, len(cuentaRepository.ObtenerTodas()))
	})

	t.Run("Retirar cantidad cero", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		nuevaCuenta, _ := CrearCuenta(cuentaRepository, "", 5.0)

		cuenta, err := RetirarDeCuenta(
			cuentaRepository,
			nuevaCuenta.ID,
			0.0,
		)

		helpers.AssertError(t, domain.ErrValorNoPositivo, err)
		helpers.AssertEquals(t, nuevaCuenta.Balance(), cuenta.Balance())
		cuentaAlmacenada, _ := cuentaRepository.Obtener(1)
		helpers.AssertEquals(t, cuentaAlmacenada.Balance(), cuenta.Balance())
		helpers.AssertDeepEquals(t, 1, len(cuentaRepository.ObtenerTodas()))
	})

	t.Run("Retirar cantidad negativa", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()
		nuevaCuenta, _ := CrearCuenta(cuentaRepository, "", 5.0)

		cuenta, err := RetirarDeCuenta(
			cuentaRepository,
			nuevaCuenta.ID,
			-6.5,
		)

		helpers.AssertError(t, domain.ErrValorNoPositivo, err)
		helpers.AssertEquals(t, nuevaCuenta.Balance(), cuenta.Balance())
		cuentaAlmacenada, _ := cuentaRepository.Obtener(1)
		helpers.AssertEquals(t, cuentaAlmacenada.Balance(), cuenta.Balance())
		helpers.AssertDeepEquals(t, 1, len(cuentaRepository.ObtenerTodas()))
	})
}

func TestEliminarCuenta(t *testing.T) {
	t.Run("Camino feliz", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()

		nuevaCuenta, _ := CrearCuenta(cuentaRepository, "", 0.0)

		cuenta, err := EliminarCuenta(cuentaRepository, nuevaCuenta.ID)

		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, 0, len(cuentaRepository.ObtenerTodas()))
		helpers.AssertDeepEquals(t, nuevaCuenta, cuenta)
	})

	t.Run("Cuenta con id cero", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()

		cuenta, err := EliminarCuenta(cuentaRepository, 0)

		helpers.AssertError(t, err, domain.ErrCuentaIDInvalido)
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuenta)
		helpers.AssertDeepEquals(t, 0, len(cuentaRepository.ObtenerTodas()))
	})

	t.Run("Cuenta inexistente", func(t *testing.T) {
		cuentaRepository := createCuentaRepository()

		cuenta, err := EliminarCuenta(cuentaRepository, 5)

		helpers.AssertError(t, err, domain.ErrCuentaNoExiste)
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuenta)
		helpers.AssertDeepEquals(t, 0, len(cuentaRepository.ObtenerTodas()))
	})
}

func createCuentaRepository() domain.CuentaRepository {
	return &internals.CuentaRepositoryInMemory{}
}
