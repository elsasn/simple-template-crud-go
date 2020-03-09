package usecase

import (
	"belajar_api/app"
	"belajar_api/modules/v1/makul/model"
	"belajar_api/modules/v1/makul/repo"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// HandleCreateMakul mengelola routing POST /students
// menerima request body mahasiswa yang akan ditambahkan
// validasi sederhana berupa pengecekan nim dan nama
// proses penyimpanan dilakukan oleh fungsi 'repo.CreateStudent'
func HandleCreateMakul(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		makulRequest := &model.Makul{}
		if err := c.Bind(makulRequest); err != nil {
			fmt.Printf("[ERROR], %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		if makulRequest.Nama == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "nim tidak boleh kosong"})
		}

		if makulRequest.JumlahSKS == 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "jumlah sks tidak boleh kosong"})
		}

		err := repo.CreateMakul(a, makulRequest)
		if err != nil {
			fmt.Printf("[ERROR], %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Berhasil tambah mata kuliah"})
	}
	return echo.HandlerFunc(h)
}

// HandleGetMakul mengelola routing GET /students
// menampilkan data seluruh mahasiswa
// proses pengambilan data dari database dilakukan oleh fungsi 'repo.GetStudents'
func HandleGetMakul(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		makuls, err := repo.GetAllMakul(a)
		if err != nil {
			fmt.Printf("[ERROR], %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.MakulResponse{
			Count: len(makuls),
			Data:  makuls,
		})
	}
	return echo.HandlerFunc(h)
}

// HandleUpdateMakul mengelola routing PUT /students
// menerima query param uuid mahasiswa yang akan diubah, dan
// menerima request body mahasiswa yang akan diubah
// validasi sederhana berupa pengecekan uuid dan nama
// proses pengubahan data dilakukan oleh fungsi 'repo.UpdateStudent'
func HandleUpdateMakul(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		if uuid == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "uuid tidak boleh kosong"})
		}

		makulRequest := &model.Makul{}
		if err := c.Bind(makulRequest); err != nil {
			fmt.Printf("[ERROR], %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		if makulRequest.Nama == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "nama makul tidak boleh kosong"})
		}

		makulRequest.UUID = uuid

		err := repo.UpdateMakul(a, makulRequest)
		if err != nil {
			fmt.Printf("[ERROR], %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Berhasil ubah mata kuliah"})
	}
	return echo.HandlerFunc(h)
}

// HandleDeleteMakul mengelola routing DELETE /students
// menerima query param uuid mahasiswa yang akan dihapus
// validasi sederhana berupa pengecekan uuid
// proses penghapusan data dilakukan oleh fungsi 'repo.DeleteStudent'
func HandleDeleteMakul(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		if uuid == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "uuid tidak boleh kosong"})
		}

		err := repo.DeleteMakul(a, uuid)
		if err != nil {
			fmt.Printf("[ERROR], %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Berhasil hapus mata kuliah"})
	}
	return echo.HandlerFunc(h)
}
