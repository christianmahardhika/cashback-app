package entity

type Tenant struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	ApiKey    string `json:"api_key" db:"api_key"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	DeletedAt string `json:"deleted_at" db:"deleted_at"`
}
