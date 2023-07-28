// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package model

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Gender string

const (
	GenderM Gender = "M"
	GenderF Gender = "F"
	GenderN Gender = "N"
)

func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

type NullGender struct {
	Gender Gender `json:"gender"`
	Valid  bool   `json:"valid"` // Valid is true if Gender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGender) Scan(value interface{}) error {
	if value == nil {
		ns.Gender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Gender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Gender), nil
}

type Role struct {
	RoleID   int32              `json:"role_id"`
	RoleName string             `json:"role_name"`
	CreateAt pgtype.Timestamptz `json:"create_at"`
}

// type User struct {
// 	UserID    pgtype.UUID        `json:"user_id"`
// 	Email     string             `json:"email"`
// 	Password  string             `json:"password"`
// 	Phone     pgtype.Text        `json:"phone"`
// 	Region    pgtype.Text        `json:"region"`
// 	Gender    utils.NullGender   `json:"gender"`
// 	Title     string             `json:"title"`
// 	FirstName string             `json:"first_name"`
// 	LastName  string             `json:"last_name"`
// 	CreateAt  pgtype.Timestamptz `json:"create_at"`
// 	UpdateAt  pgtype.Timestamptz `json:"update_at"`
// 	Age       pgtype.Int2        `json:"age"`
// }

type UserAddress struct {
	AddressID int32              `json:"address_id"`
	UserID    pgtype.UUID        `json:"user_id"`
	Address   pgtype.Text        `json:"address"`
	CreateAt  pgtype.Timestamptz `json:"create_at"`
	UpdateAt  pgtype.Timestamptz `json:"update_at"`
}

type UserRole struct {
	UserID          pgtype.UUID        `json:"user_id"`
	Role            int32              `json:"role"`
	RoleDescription pgtype.Text        `json:"role_description"`
	CreateAt        pgtype.Timestamptz `json:"create_at"`
	UpdateAt        pgtype.Timestamptz `json:"update_at"`
}
