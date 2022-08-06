package domain

import "time"

type newModelFromFunc func(createdAt time.Time, deletedAt *time.Time) Model

type newModelFunc func() Model

var (
	NewModelFrom newModelFromFunc = newModelFrom
	NewModel     newModelFunc     = newModel
)

type Model struct {
	createdAt time.Time
	deletedAt *time.Time
}

func newModel() Model {
	data := time.Now().In(time.UTC)

	return Model{
		createdAt: data,
		deletedAt: nil,
	}
}

func newModelFrom(
	createdAt time.Time,
	deletedAt *time.Time,
) Model {
	var deleted time.Time
	if deletedAt != nil {
		deleted = deletedAt.In(time.UTC)
	}

	return Model{
		createdAt: createdAt.In(time.UTC),
		deletedAt: &deleted,
	}
}

//func (m *Model) Update() {
//	m.updatedAt = time.Now().In(time.UTC)
//}

func (m *Model) Delete() {
	now := time.Now().In(time.UTC)
	m.deletedAt = &now
}

func (m Model) GetCreatedAt() time.Time {
	return m.createdAt
}

//func (m Model) GetUpdatedAt() time.Time {
//	return m.updatedAt
//}

func (m Model) GetDeletedAt() *time.Time {
	return m.deletedAt
}
