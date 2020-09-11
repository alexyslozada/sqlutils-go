package sqlutils_go

import (
	"testing"
)

func TestModel_NewConnection(t *testing.T) {
	m := Model{
		"postgres",
		"edcursos",
		"edcursos",
		"localhost",
		"edcursos",
		5432,
	}

	db, err := m.NewConnection()
	if err != nil {
		t.Error(err)
	}

	_ = db
}
