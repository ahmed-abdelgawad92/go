package models

import (
	"database/sql"
	"rest_api/entities"
)

type UserModel struct {
	Db *sql.DB
}

func (userModel UserModel) FindAll() (user []entities.User, err error) {
	rows, err := userModel.Db.Query("SELECT id, vorname, abrechnung, status FROM user")
	if err != nil {
		return nil, err
	} else {
		var users []entities.User
		for rows.Next() {
			var user entities.User
			err := rows.Scan(&user.ID, &user.Vorname, &user.Abrechnung, &user.Status)

			if err != nil {
				return nil, err
			}

			users = append(users, user)
		}
		return users, nil
	}
}
