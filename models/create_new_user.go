package models

// Create new user
func (u *User) Create() error {
	const query = `
			INSERT INTO users (username, password, email, is_active, created, updated)
			VALUES (?, ?, ?, 0, NOW(), NOW())
	`
	stmt, err := u.Db.Prepare(query)
	if err != nil {
		u.Log.Println(err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(u.Username, u.Password, u.Email)
	if err != nil {
		u.Log.Println(err)
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		u.Log.Println(err)
		return err
	}

	u.ID = uint(id)

	return nil
}
