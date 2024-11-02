package schemas

import (
	"time"

	"gorm.io/gorm"
)

// Definimos a estrutura da tabela no banco de dados

type Productions struct {
	gorm.Model
	Name        string
	Producer    string
	Movie       bool
	Series      bool
	Protagonist string
	ReleaseDate time.Time
	Notice      string
	Assessment  string
}

// Criamos a amarração entre os campos com relação aos campos do json

type ProductionsResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deteledAt,omitempty"`
	Name        string    `json:"name"`
	Producer    string    `json:"producer"`
	Movie       bool      `json:"movie"`
	Series      bool      `json:"series"`
	Protagonist string    `json:"protagonist"`
	ReleaseDate time.Time `json:"releaseDate"`
	Notice      string    `json:"notice"`
	Assessment  string    `json:"assessment"`
}
