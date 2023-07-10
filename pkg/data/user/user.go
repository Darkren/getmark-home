// Package user contains User model and its repository.
package user

// User is a model of a user of the service.
type User struct {
	ID       int64  `gorm:"column:id;primaryKey" json:"id"`
	Login    string `gorm:"column:login" json:"login"`
	Password string `gorm:"column:password" json:"password"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
}
