package model

type Transaction struct {
	ID        string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Tanggal   string `json:"tanggal" gorm:"column:tanggal;type:date;not null"`
	Jenis     string `json:"jenis" gorm:"column:jenis;type:varchar(10);not null"`
	Kategori  string `json:"kategori" gorm:"column:kategori;type:varchar(50);not null"`
	Deskripsi string `json:"deskripsi" gorm:"column:deskripsi;type:text"`
	Jumlah    int    `json:"jumlah" gorm:"column:jumlah;type:int;not null"`
}

func (Transaction) TableName() string {
	return "transactions"
}

type User struct {
	ID       string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Username string `json:"username" gorm:"column:username;unique;not null"`
	Password string `json:"-" gorm:"column:password;not null"`
	Role     string `json:"role" gorm:"column:role;type:varchar(30);default:user"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (User) TableName() string { return "users" }
