package users

import "time"

type UserDTO struct {
	UserID              int32     `json:"user_id"`
	UserFirstname       string    `json:"user_firstname"`
	UserLastname        string    `json:"user_lastname"`
	UserEmail           string    `json:"user_email"`
	UserPhone           string    `json:"user_phone"`
	UserIsActive        bool      `json:"user_is_active"`
	UserDateCreated     time.Time `json:"user_date_created"`
	UserRecruitmentDate time.Time `json:"user_recruitment_date"`
}
