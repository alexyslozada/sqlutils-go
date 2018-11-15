package sqlutil_go_test

import (
	"testing"

	"github.com/alexyslozada/sqlutil-go"
)

func TestModel_NewConnection(t *testing.T) {
	m := sqlutil_go.Model{
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

func TestSingleton(t *testing.T) {
	_, err := sqlutil_go.GetConnection()
	if err != nil {
		t.Error(err)
	}

	m2 := sqlutil_go.Model{}
	_, err = m2.NewConnection()
	if err == nil {
		t.Error("no validó que el motor no viniera vacío")
	}
}

//func TestModel_NewConnection_error(t *testing.T) {
//	err := sqlutil_go.CloseConnection()
//	if err != nil {
//		t.Error(err)
//	}
//
//	m2 := sqlutil_go.Model{}
//	db, err := m2.NewConnection()
//	if db.Ping() == nil {
//		t.Error("se esperaba error de conexión")
//	}
//}
