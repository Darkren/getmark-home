package product

import "github.com/Darkren/getmark-home/pkg/data/user"

// Product is a model of a product.
type Product struct {
	Barcode string `gorm:"column:barcode;primaryKey" json:"barcode"`
	Name    string `gorm:"column:name" json:"name"`
	Desc    string `gorm:"column:desc" json:"desc"`
	Cost    int64  `gorm:"column:cost" json:"cost"`
	UserID  int64  `gorm:"user_id" json:"user_id"`

	User user.User `gorm:"foreignKey:UserID" json:"-"`
}
