package repo

func createMakulQuery() string {
	return "INSERT INTO matakuliah (nama_matakuliah,jumlah_sks,uuid) VALUES (?,?,?)"
}

func getAllMakulQuery() string {
	return "SELECT nama_matakuliah,jumlah_sks,uuid FROM matakuliah"
}

func updateMakulQuery() string {
	return "UPDATE matakuliah SET nama_matakuliah = ?, jumlah_sks = ? WHERE uuid = ?"
}

func deleteMakulQuery() string {
	return "DELETE FROM matakuliah WHERE uuid = ?"
}
