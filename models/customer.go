package models

type Customer struct {
	Id   int64  `gorm:"primaryKey;column:customer_id;autoIncrement"`
	Name string `gorm:"column:customer_name"`
}

func (Customer) TableName() string {
	return "t_customer"
}