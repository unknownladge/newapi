package hander

type selection1 interface {
	selectone(string) (string, error)
	selectall2() (string, error)
}

// func (d Dbaseapi) Selectsinglearticle(id int) (string, error) {
// 	name, err := d.Dbaserespo.Selectsinglearticle(id)
// 	if err != nil {
// 		return "empty", err
// 	}
// 	return name, nil
// }

// func (m *DbRepoMock) Selectsinglearticle(id int) (string, error) {
// 	arguments := m.Called(id)
// 	return arguments.String(0), arguments.Error(1)
// }
