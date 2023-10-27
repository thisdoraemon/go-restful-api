package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		if errorRollback != nil {
			PanicIfError(errorRollback)
			panic(err)
		} else {
			errorCommit := tx.Commit()
			PanicIfError(errorCommit)
		}
	}
}
