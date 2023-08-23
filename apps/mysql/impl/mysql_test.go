package impl_test

import "testing"

func TestUnzipMySQLFile(t *testing.T) {
	err := svc.UnzipMySQLFile(ctx)
	if err != nil {
		t.Log(err)
	}
}

func TestCreateMySQLDir(t *testing.T) {
	err := svc.CreateMySQLDir(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateMySQLUser(t *testing.T) {
	err := svc.CreateMySQLUser(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChangeMySQLDirPerm(t *testing.T) {
	err := svc.ChangeMySQLDirPerm(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInitialMySQL(t *testing.T) {
	err := svc.InitialMySQL(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
