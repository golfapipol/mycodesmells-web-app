package db

type DataSource interface {
	Countries() ([]Country, error)
}
