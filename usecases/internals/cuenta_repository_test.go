package internals

import (
	"testing"

	"github.com/iyaki/gastos/domain"
	"github.com/iyaki/gastos/helpers"
)

func TestAgregar(t *testing.T) {
	t.Run("Camino feliz", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}
		cuenta := domain.Cuenta{}

		cuentaReturn, err := repository.Agregar(cuenta)

		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, 1, len(repository.cuentas))
		helpers.AssertDeepEquals(t, cuentaReturn, repository.cuentas[0])
		cuenta.ID = repository.cuentas[0].ID
		helpers.AssertDeepEquals(t, cuenta, repository.cuentas[0])
	})

	t.Run("Cuenta con ID", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		_, err := repository.Agregar(domain.Cuenta{ID: 5})

		helpers.AssertError(t, domain.ErrCuentaNuevaConID, err)
		helpers.AssertEquals(t, 0, len(repository.cuentas))
	})
}

func TestEliminar(t *testing.T) {
	t.Run("Camino feliz", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		cuenta, err1 := repository.Agregar(domain.Cuenta{})

		helpers.AssertNoError(t, err1)
		helpers.AssertEquals(t, 1, len(repository.cuentas))

		cuentaEliminada, err2 := repository.Eliminar(cuenta.ID)

		helpers.AssertNoError(t, err2)
		helpers.AssertEquals(t, 0, len(repository.cuentas))
		helpers.AssertDeepEquals(t, cuenta, cuentaEliminada)
	})

	t.Run("Cuenta inexistente", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		cuenta, err := repository.Eliminar(2)

		helpers.AssertError(t, domain.ErrCuentaNoExiste, err)
		helpers.AssertEquals(t, 0, len(repository.cuentas))
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuenta)
	})

	t.Run("Cuenta con ID cero", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		cuenta, err := repository.Eliminar(0)

		helpers.AssertError(t, domain.ErrCuentaIDInvalido, err)
		helpers.AssertEquals(t, 0, len(repository.cuentas))
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuenta)
	})

	t.Run("Cuenta con ID negativo", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		cuenta, err := repository.Eliminar(0)

		helpers.AssertError(t, domain.ErrCuentaIDInvalido, err)
		helpers.AssertEquals(t, 0, len(repository.cuentas))
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuenta)
	})
}

func TestActualizar(t *testing.T) {
	t.Run("Camino feliz", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		cuenta, err1 := repository.Agregar(domain.Cuenta{})

		helpers.AssertNoError(t, err1)
		helpers.AssertEquals(t, 1, len(repository.cuentas))

		cuenta.Nombre = "Prueba"
		cuentaActualizada, err2 := repository.Actualizar(cuenta)

		helpers.AssertNoError(t, err2)
		helpers.AssertEquals(t, 1, len(repository.cuentas))
		helpers.AssertDeepEquals(t, cuenta, cuentaActualizada)
	})

	t.Run("Cuenta inexistente", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}
		cuenta := domain.Cuenta{}
		cuenta.ID = 3

		cuentaActualizada, err := repository.Actualizar(cuenta)

		helpers.AssertError(t, domain.ErrCuentaNoExiste, err)
		helpers.AssertEquals(t, 0, len(repository.cuentas))
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuentaActualizada)
	})

	t.Run("Cuenta con ID cero", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		cuenta, err := repository.Actualizar(domain.Cuenta{})

		helpers.AssertError(t, domain.ErrCuentaIDInvalido, err)
		helpers.AssertEquals(t, 0, len(repository.cuentas))
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuenta)
	})

	t.Run("Cuenta con ID negativo", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}
		cuenta := domain.Cuenta{}
		cuenta.ID = -3

		cuentaActualizada, err := repository.Actualizar(cuenta)

		helpers.AssertError(t, domain.ErrCuentaIDInvalido, err)
		helpers.AssertEquals(t, 0, len(repository.cuentas))
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuentaActualizada)
	})
}

func TestObtener(t *testing.T) {
	t.Run("Camino feliz", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		repository.Agregar(domain.Cuenta{Nombre: "Primera Cuenta"})
		cuenta2, _ := repository.Agregar(domain.Cuenta{Nombre: "Segunda Cuenta"})
		repository.Agregar(domain.Cuenta{Nombre: "Tercera Cuenta"})

		cuentaObtenida, err := repository.Obtener(cuenta2.ID)

		helpers.AssertNoError(t, err)
		helpers.AssertDeepEquals(t, cuenta2, cuentaObtenida)
	})

	t.Run("Cuenta inexistente", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		cuentaObtenida, err := repository.Obtener(1)

		helpers.AssertError(t, domain.ErrCuentaNoExiste, err)
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuentaObtenida)
	})

	t.Run("ID cero", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		cuentaObtenida, err := repository.Obtener(0)

		helpers.AssertError(t, domain.ErrCuentaIDInvalido, err)
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuentaObtenida)
	})

	t.Run("ID negativo", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		cuentaObtenida, err := repository.Obtener(-7)

		helpers.AssertError(t, domain.ErrCuentaIDInvalido, err)
		helpers.AssertDeepEquals(t, domain.Cuenta{}, cuentaObtenida)
	})
}

func TestObtenerTodas(t *testing.T) {
	t.Run("Camino feliz", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		cuenta1, _ := repository.Agregar(domain.Cuenta{Nombre: "Primera Cuenta"})
		cuenta2, _ := repository.Agregar(domain.Cuenta{Nombre: "Segunda Cuenta"})
		cuenta3, _ := repository.Agregar(domain.Cuenta{Nombre: "Tercera Cuenta"})

		cuentasObtenidas := repository.ObtenerTodas()

		helpers.AssertEquals(t, 3, len(cuentasObtenidas))
		helpers.AssertDeepEquals(t, cuenta1, cuentasObtenidas[0])
		helpers.AssertDeepEquals(t, cuenta2, cuentasObtenidas[1])
		helpers.AssertDeepEquals(t, cuenta3, cuentasObtenidas[2])
	})

	t.Run("Sin cuentas almacenadas", func(t *testing.T) {
		repository := CuentaRepositoryInMemory{}

		cuentasObtenidas := repository.ObtenerTodas()

		helpers.AssertEquals(t, 0, len(cuentasObtenidas))
	})
}
