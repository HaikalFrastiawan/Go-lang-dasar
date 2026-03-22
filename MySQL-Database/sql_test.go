package mysqldatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('Haikal','HAIKAL')"
	_, err := db.ExecContext(ctx,script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}