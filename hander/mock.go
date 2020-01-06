package hander

import (
	"github.com/stretchr/testify/mock"
)

type Dbase interface {
	Selectsinglearticle(int) (string, error)
}
type Dbaseapi struct {
	Dbaserespo Dbase
}

func (d Dbaseapi) Selectsinglearticle(id int) (string, error) {
	name, err := d.Dbaserespo.Selectsinglearticle(id)
	if err != nil {
		return "empty", err
	}
	return name, nil
}

type DbRepoMock struct {
	mock.Mock
}

func (m *DbRepoMock) Selectsinglearticle(id int) (string, error) {
	arguments := m.Called(id)
	return arguments.String(0), arguments.Error(1)
}

// type OrderRepository interface {
// 	FindAllOrdersByUserID(userID int) (Article, error)
// }
// type orderRepository struct {
// 	db orm.DB
// }

// func (r *orderRepository) SetDB(db orm.DB) {
// 	r.db = db
// }

// func (r *orderRepository) getDB() orm.DB {
// 	if r.db != nil {
// 		return r.db
// 	}

// 	//r.db = application.ResolveDB()
// 	return r.db
// }

// func (r *orderRepository) FindAllOrdersByUserID(userID int) (Article, error) {
// 	orders := Article{}
// 	err := r.getDB().Model(&orders).Where("user_id = ?", userID).Order("placed_at DESC").Select()
// 	return orders, err
// }

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// type Customer interface {
// 	GetCustomerByID(int) (string, error)
// } //

// type CustomerAPI struct {
// 	CustomerRepo Customer
// } //

// func (c CustomerAPI) GetCustomerName(id int) (string, error) {
// 	name, err := c.CustomerRepo.GetCustomerByID(id)
// 	if err != nil {
// 		return "empty", err
// 	}
// 	return name, nil
// }

// type CustomerRepoMock struct {
// 	mock.Mock
// }

// func (m *CustomerRepoMock) GetCustomerByID(id int) (string, error) {
// 	arguments := m.Called(id)
// 	return arguments.String(0), arguments.Error(1)
// }
