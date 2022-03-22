package models

func (u *User) Get() error {
	const q = `SELECT id, username, password, email, is_active FROM users`
	err := u.Db.QueryRow(q+" WHERE id=?", u.ID).Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.IsActive)

	if err != nil {
		u.Log.Print(err)
		return err
	}

	return nil
}
