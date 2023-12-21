package data

import (
	"gorm.io/gorm"
)

type Employee struct {
	ID                    int
	Full_name             string
	Department            string
	Post                  string
	Internal_phone_number string
	City_phone_number     string
	Mobile_phone_number   string
}

type employeesRepo struct {
	db *gorm.DB
}

func newEmployeesRepo(conn *gorm.DB) *employeesRepo {
	return &employeesRepo{conn}
}

func (er *employeesRepo) Employee() ([]Employee, error) {
	employees := make([]Employee, 0)

	err := er.db.Find(&employees).Error
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (er *employeesRepo) AddEmployee(emp *Employee) (int, error) {
	//TODO: check data integrity

	err := er.db.Create(&emp).Error
	if err != nil {
		return 0, err
	}
	id := emp.ID

	return id, nil
}

func (er *employeesRepo) DeleteEmployee(id int) error {
	err := er.db.First(&Employee{}, id).Error
	if err != nil {
		return err
	}

	err = er.db.Delete(&Employee{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

//func (er *employeesRepo) SearchEmployee(emp *Employee) ([]Employee, error) {
//	employees := make([]Employee, 0)
//
//	query := er.db.Where(&emp)
//	err := query.Find(&emp).Error
//	if err != nil {
//		return nil, err
//	}
//
//	return employees, nil
//}
