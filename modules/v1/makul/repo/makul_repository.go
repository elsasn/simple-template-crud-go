package repo

import (
	"belajar_api/app"
	"belajar_api/modules/v1/makul/model"
	"fmt"
)

// CreateMakul mengelola proses penambahan data mahasiswa ke database
// query diletakkan di file yang berbeda
// Penambahan diawali dengan membuat transaksi
// query akan dievaluasi dengan 'tx.Prepare'
// apabila eksekusi query berhasil, transaksi akan dicommit
// jika gagal, tranksaksi dirollback
func CreateMakul(a app.App, makulRequest *model.Makul) error {
	sqlQuery := createMakulQuery()

	tx, err := a.DB.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction, %s", err.Error())
	}

	stmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return fmt.Errorf("error preparing query, %s", err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(&makulRequest.Nama, &makulRequest.JumlahSKS, &makulRequest.UUID)
	if err != nil {
		return fmt.Errorf("error executing query, %s", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error commiting query, %s", err.Error())
	}

	return nil
}

// GetAllMakul mengelola proses pengambilan data mahasiswa dari database
// Jumlah kolom yang diambil dalam query harus sama dengan jumlah variabel yang discan
func GetAllMakul(a app.App) ([]model.Makul, error) {
	sqlQuery := getAllMakulQuery()

	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get all makul, %s", err.Error())
	}
	defer rows.Close()

	makuls := []model.Makul{}
	for rows.Next() {
		var makul model.Makul
		err := rows.Scan(&makul.Nama, &makul.JumlahSKS, &makul.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan makul row, %s", err.Error())
		}
		makuls = append(makuls, makul)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error student rows, %s", err.Error())
	}

	return makuls, nil
}

// UpdateMakul mengelola proses pengubahan data mahasiswa di database
// Pengubahan diawali dengan membuat transaksi
// query akan dievaluasi dengan 'tx.Prepare'
// apabila eksekusi query berhasil, transaksi akan dicommit
// jika gagal, tranksaksi dirollback
func UpdateMakul(a app.App, makulRequest *model.Makul) error {
	sqlQuery := updateMakulQuery()

	tx, err := a.DB.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction, %s", err.Error())
	}

	stmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return fmt.Errorf("error preparing query, %s", err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(&makulRequest.Nama, &makulRequest.JumlahSKS, &makulRequest.UUID)
	if err != nil {
		return fmt.Errorf("error executing update query, %s", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error commiting query, %s", err.Error())
	}

	return nil
}

// DeleteMakul mengelola proses penghapusan data mahasiswa di database
// Membutuhkan uuid mahasiswa yang akan dihapus
// Penambahan diawali dengan membuat transaksi
// query akan dievaluasi dengan 'tx.Prepare'
// apabila eksekusi query berhasil, transaksi akan dicommit
// jika gagal, tranksaksi dirollback
func DeleteMakul(a app.App, uuid string) error {
	sqlQuery := deleteMakulQuery()

	tx, err := a.DB.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction, %s", err.Error())
	}

	stmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return fmt.Errorf("error preparing query, %s", err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(&uuid)
	if err != nil {
		return fmt.Errorf("error executing delete query, %s", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error commiting query, %s", err.Error())
	}

	return nil
}
