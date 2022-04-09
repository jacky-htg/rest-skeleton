package models

import "context"

// Update user by id
func (u *User) Update(ctx context.Context) error {
	const q string = `UPDATE users SET is_active = ? WHERE id = ?`
	stmt, err := u.Db.PrepareContext(ctx, q)
	if err != nil {
		u.Log.Print(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, u.IsActive, u.ID)
	if err != nil {
		u.Log.Print(err)
		return err
	}

	return nil
}
