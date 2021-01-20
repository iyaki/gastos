package internals

import "github.com/iyaki/gastos/domain"

// CuentaRepositoryInMemory es una implementacion en memoria de
// la interfaz CuentaRepository
type CuentaRepositoryInMemory struct {
	cuentas []domain.Cuenta
}

// Agregar agrega una cuenta al repositorio y devuelve la cuenta con un ID asignado en caso de exito
func (c *CuentaRepositoryInMemory) Agregar(cuenta domain.Cuenta) (domain.Cuenta, error) {

	if cuenta.ID != 0 {
		cuenta.ID = 0
		return cuenta, domain.ErrCuentaNuevaConID
	}

	c.cuentas = append(c.cuentas, cuenta)
	c.cuentas[len(c.cuentas)-1].ID = domain.CuentaID(len(c.cuentas))
	return c.cuentas[len(c.cuentas)-1], nil
}

// Actualizar actualizar y devuelve una cuenta del repositorio o una cuenta vacia en caso de
// que no existiera una con el ID solicitado
func (c *CuentaRepositoryInMemory) Actualizar(cuenta domain.Cuenta) (domain.Cuenta, error) {
	if cuenta.ID == 0 {
		return domain.Cuenta{}, domain.ErrCuentaIDInvalido
	}

	for i, cuentaAlmacenada := range c.cuentas {
		if cuentaAlmacenada.ID == cuenta.ID {
			c.cuentas[i] = cuenta
			return c.cuentas[i], nil
		}
	}

	return domain.Cuenta{}, domain.ErrCuentaNoExiste
}

// Eliminar elimina y devuelve una cuenta del repositorio o una cuenta vacia en caso de
// que no existiera una con el ID solicitado
func (c *CuentaRepositoryInMemory) Eliminar(cuentaID domain.CuentaID) (domain.Cuenta, error) {
	if cuentaID == 0 {
		return domain.Cuenta{}, domain.ErrCuentaIDInvalido
	}

	for i, cuentaAlmacenada := range c.cuentas {
		if cuentaAlmacenada.ID == cuentaID {
			c.cuentas[i] = c.cuentas[len(c.cuentas)-1]
			c.cuentas = c.cuentas[:len(c.cuentas)-1]
			return cuentaAlmacenada, nil
		}
	}

	return domain.Cuenta{}, domain.ErrCuentaNoExiste
}

// Obtener devuelve una cuenta del repositorio o una cuenta vacia en caso de
// que no existiera una con el ID solicitado
func (c *CuentaRepositoryInMemory) Obtener(cuentaID domain.CuentaID) (domain.Cuenta, error) {
	if cuentaID == 0 {
		return domain.Cuenta{}, domain.ErrCuentaIDInvalido
	}

	for _, cuentaAlmacenada := range c.cuentas {
		if cuentaAlmacenada.ID == cuentaID {
			return cuentaAlmacenada, nil
		}
	}

	return domain.Cuenta{}, domain.ErrCuentaNoExiste
}

// ObtenerTodas devuelve todas las cuentas del repositorio
func (c *CuentaRepositoryInMemory) ObtenerTodas() []domain.Cuenta {
	return c.cuentas
}

var _ domain.CuentaRepository = &CuentaRepositoryInMemory{}
