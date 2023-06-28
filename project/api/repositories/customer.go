package repositories

import "gorm.io/gorm"

type Customer struct {
	ID   uint
	Name string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) (CustomerRepository, error) {
	err := db.AutoMigrate(&Customer{})
	if err != nil {
		return nil, err
	}

	var count int64
	db.Model(&Customer{}).Count(&count)
	if count <= 0 {
		customners := []Customer{
			{Name: "Bond"},
			{Name: "Yod"},
			{Name: "Ham"},
			{Name: "Ble"},
		}
		err = db.Create(&customners).Error
		if err != nil {
			return nil, err
		}
	}

	return &customerRepository{db}, nil
}

func (r *customerRepository) FindAll() (customers []Customer, err error) {
	err = r.db.Find(&customers).Error
	return customers, err
}
