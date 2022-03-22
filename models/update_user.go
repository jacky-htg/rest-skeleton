package models

// Update user by id
func (u *User) Update() error {
	const q string = `UPDATE users SET is_active = ? WHERE id = ?`
	stmt, err := u.Db.Prepare(q)
	if err != nil {
		u.Log.Print(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.IsActive, u.ID)
	if err != nil {
		u.Log.Print(err)
		return err
	}

	return nil
}
