package model

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Transaction struct {
	ID        int `json:"id"`
	UserID    User
	ProductID Product
	Quantity  string `json:"quantity"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type TransactionsResponse struct {
	Data []Transaction `json:"data"`
}
type DetailResponse struct {
	Data []Songs `json:"data"`
}
type ResultUser struct {
	ID      int
	Name    string
	Age     int
	Address string
}

type User2 struct {
	UserID       int    `json:"id"`
	UserName     string `json:"name"`
	UserEmail    string `json:"email"`
	UserPassword string `json:"password"`
	UserCountry  string `json:"country"`
	UserType     int    `json:"type"`
}

type Songs struct {
	SongID       int    `json:"id"`
	SongTittle   string `json:"tittle"`
	SongDuration string `json:"duration"`
	SongSinger   string `json:"singer"`
}

type DetailPlaylistSong struct {
	Song       Songs `json:"Song"`
	TimePlayed int   `json:"time played"`
}

type DetailPlaylistResponse struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    []DetailPlaylistSong `json:"data"`
}
