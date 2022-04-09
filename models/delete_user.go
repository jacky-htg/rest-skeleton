package models

import "context"

// Delete user by id
func (u *User) Delete(ctx context.Context) error {
	const q string = `DELETE FROM users WHERE id = ?`
	stmt, err := u.Db.PrepareContext(ctx, q)
	if err != nil {
		u.Log.Print(err)
		return err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, u.ID); err != nil {
		u.Log.Print(err)
		return err
	}

	return nil
}
