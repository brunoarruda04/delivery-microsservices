package domain

type Restaurant struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Repository interface {
	Create(Restaurant) error
	Get(string) (Restaurant, error)
	Update(string, Restaurant) error
	Delete(string) error
}
