package impl_test

import "testing"

func TestUnzipMySQLFile(t *testing.T) {
	err := svc.UnzipMySQLFile(ctx)
	if err != nil {
		t.Log(err)
	}
}
