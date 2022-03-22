package models

// Delete user by id
func (u *User) Delete() error {
	const q string = `DELETE FROM users WHERE id = ?`
	stmt, err := u.Db.Prepare(q)
	if err != nil {
		u.Log.Print(err)
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(u.ID); err != nil {
		u.Log.Print(err)
		return err
	}

	return nil
}
