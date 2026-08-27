package cart

// RepositoryMock — простейший ручной мок RepositoryInterface.
// Каждому методу интерфейса соответствует своё поле-функция,
// которую задаёт тест.
type RepositoryMock struct {
	addFunc    func(userID int64, item Item)
	removeFunc func(userID, sku int64)
	clearFunc  func(userID int64)
	getFunc    func(userID int64) map[int64]Item
}

func (m *RepositoryMock) add(userID int64, item Item) {
	m.addFunc(userID, item)
}

func (m *RepositoryMock) remove(userID, sku int64) {
	m.removeFunc(userID, sku)
}

func (m *RepositoryMock) clear(userID int64) {
	m.clearFunc(userID)
}

func (m *RepositoryMock) get(userID int64) map[int64]Item {
	return m.getFunc(userID)
}
