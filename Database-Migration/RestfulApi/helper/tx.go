package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		ErrorRollback := tx.Rollback()
		if ErrorRollback != nil {
			panic(ErrorRollback)
		}
		panic(err)
	} else {
		ErrorCommit := tx.Commit()
		PanicIfError(ErrorCommit)
	}
}