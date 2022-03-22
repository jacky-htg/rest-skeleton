package models

// ListUsers : http handler for returning list of users
func (u *User) List() ([]User, error) {
	var list []User

	const q = `SELECT id, username, password, email, is_active FROM users`

	rows, err := u.Db.Query(q)
	if err != nil {
		u.Log.Print(err)
		return list, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.IsActive); err != nil {
			u.Log.Print(err)
			return list, err
		}
		list = append(list, user)
	}

	if err := rows.Err(); err != nil {
		u.Log.Print(err)
		return list, err
	}

	return list, nil
}
