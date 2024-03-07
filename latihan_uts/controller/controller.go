package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	m "latihan_uts/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetDetailUserTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT t.ID, u.ID, u.Name, u.Age, u.Address , p.ID, p.Name, p.Price, t.Quantity
	FROM transactions t JOIN users u ON t.UserID = u.ID JOIN products p ON t.ProductID = p.ID`

	detailTransactionRow, err := db.Query(query)
	if err != nil {
		SendErrorResponse(w, 500, "StatusBadRequest")
		return
	}

	var detailTransactions []m.Transaction
	for detailTransactionRow.Next() {

		var user m.User
		var product m.Product
		var transaction m.Transaction

		if err := detailTransactionRow.Scan(&transaction.ID, &user.ID, &user.Name, &user.Age, &user.Address, &product.ID, &product.Name, &product.Price, &transaction.Quantity); err != nil {
			SendErrorResponse(w, 500, "StatusBadRequest")
			return
		} else {
			transaction.UserID = user
			transaction.ProductID = product
			detailTransactions = append(detailTransactions, transaction)
		}
	}

	if len(detailTransactions) == 0 {
		SendErrorResponse(w, 500, "StatusBadRequest")
		return
	}

	SendSuccessResponse(w, 200, "Success")
	w.Header().Set("Content-Type", "application/json")
	var response m.TransactionsResponse
	json.NewEncoder(w).Encode(response)
}

func GetDetailUserTransByID(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	userID := mux.Vars(r)["id"]

	query := `SELECT t.ID, u.ID, u.Name, u.Age, u.Address , p.ID, p.Name, p.Price, t.Quantity
	FROM transactions t JOIN users u ON t.UserID = u.ID JOIN products p ON t.ProductID = p.ID WHERE u.ID = ? `

	detailTransactionRow, err := db.Query(query, userID)
	if err != nil {
		SendErrorResponse(w, 500, "StatusBadRequest")
		return
	}

	var detailTransactions []m.Transaction
	for detailTransactionRow.Next() {

		var user m.User
		var product m.Product
		var transaction m.Transaction

		if err := detailTransactionRow.Scan(&transaction.ID, &user.ID, &user.Name, &user.Age, &user.Address, &product.ID, &product.Name, &product.Price, &transaction.Quantity); err != nil {
			SendErrorResponse(w, 500, "StatusBadRequest")
			return
		} else {
			transaction.UserID = user
			transaction.ProductID = product
			detailTransactions = append(detailTransactions, transaction)
		}
	}
	if len(detailTransactions) == 0 {
		SendErrorResponse(w, 400, "Data Tidak Ditemukan!")
		return
	}

	SendSuccessResponse(w, 200, "Success")

}

func DeleteSingleProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	productID := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(productID)
	if err != nil {
		SendErrorResponse(w, 500, "StatusBadRequest")
		return
	}

	_, err = db.Exec("DELETE FROM transactions WHERE ProductID = ?", idInt)
	if err != nil {
		SendErrorResponse(w, 500, "StatusInternalServerError")
		return
	}

	stmt, err := db.Prepare("DELETE FROM products WHERE ID = ?")
	if err != nil {
		SendErrorResponse(w, 500, "StatusInternalServerError")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(idInt)
	if err != nil {
		SendErrorResponse(w, 500, "StatusInternalServerError")
		return
	}

	fmt.Fprintf(w, "Produk dengan ID %s berhasil dihapus bersama dengan semua transaksi terkait", productID)
}

func InsertNewProducts(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(w, 500, "StatusInternalServerError")
		return
	}

	id := r.Form.Get("id")
	userid := r.Form.Get("userid")
	productId := r.Form.Get("productid")
	quantity := r.Form.Get("quantity")

	// Validasi jumlah field
	if len(r.Form) >= 4 {
		SendErrorResponse(w, 500, "Too many fields provided")
		return
	}

	// cek klo ada field yang kosong
	if userid == "" || productId == "" || quantity == "" {
		SendErrorResponse(w, 500, "Incomplete data provided")
		return
	}

	// cek di DB dengan ID produk
	var count int
	tx, err := db.Begin()
	err = tx.QueryRow("SELECT COUNT(*) FROM products WHERE id = ?", productId).Scan(&count)
	if err != nil {
		SendErrorResponse(w, 500, "Error checking product existence")
		return
	}

	// Jika produk belum ada, tambahkan produk baru
	if count == 0 {
		_, err = tx.Exec("INSERT INTO products (id, name, price) VALUES (?, '', 0)", productId)
		if err != nil {
			SendErrorResponse(w, 500, "Error inserting new product")
			return
		}
	}

	// Insert transaksi baru
	_, err = tx.Exec("INSERT INTO transactions (id, userID, productID, quantity) VALUES (?, ?, ?, ?)", id, userid, productId, quantity)
	if err != nil {
		SendErrorResponse(w, 500, "Error inserting transaction")
		return
	}

	fmt.Fprintf(w, "Product inserted/updated successfully and transaction created!")
}

// Nomer 1 latihan /////////////////////////////////////////////////////////////////////////////////////////
func Login(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	email := r.FormValue("Email")
	password := r.FormValue("Password")

	if email == "" || password == "" {
		SendErrorResponse(w, http.StatusBadRequest, "Email dan Password harus diisi")
		return
	}

	var dbPassword string
	query := "SELECT userPassword FROM users WHERE userEmail = ?"
	err := db.QueryRow(query, email).Scan(&dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			SendErrorResponse(w, http.StatusUnauthorized, "Email tidak ditemukan!")
			return
		}
		SendErrorResponse(w, http.StatusInternalServerError, "Kesalahan data tidak ditemukan")
		return
	}

	if password != dbPassword {
		SendErrorResponse(w, http.StatusUnauthorized, "Password salah")
		return
	}

	SendSuccessResponse(w, http.StatusOK, "Success")
}

// NO 2 latihan /////////////////////////////////////////////////////////////////////////////////////////
func GetDetailUserTransactions2(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query :=
		`SELECT s.songId, s.songTittle, s.songDuration, s.songSinger, dp.timePlayed 
		FROM songs s JOIN detailplaylistsong dp  ON s.songID = dp.songID 
		ORDER BY dp.timeplayed DESC`

	detailSongsRow, err := db.Query(query)
	if err != nil {
		SendErrorResponse(w, 400, "Kesalahan saat query")
		return
	}

	var detailSong m.DetailPlaylistSong
	var detailSongs []m.DetailPlaylistSong
	for detailSongsRow.Next() {

		if err := detailSongsRow.Scan(&detailSong.Song.SongID, &detailSong.Song.SongTittle, &detailSong.Song.SongDuration, &detailSong.Song.SongSinger, &detailSong.TimePlayed); err != nil {
			SendErrorResponse(w, 400, "Gagal mendapatkan data")
			return
		} else {
			detailSongs = append(detailSongs, detailSong)
		}
	}

	if len(detailSongs) == 0 {
		SendErrorResponse(w, 400, "Data masih kosong")
		return
	}

	SendSuccessDetailSongResponse(w, detailSongs)
	// SendSuccessResponse(w, 200, "Success")
	// w.Header().Set("Content-Type", "application/json")
	// var response m.DetailPlaylistResponse
	// response.Data = detailSongs
	// json.NewEncoder(w).Encode(response)
}
