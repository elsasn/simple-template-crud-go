package main

import (
	"belajar_api/app"
	"belajar_api/app/database"
	lg "log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// Membuat koneksi ke database, koneksi hanya dibuat satu kali dan akan digunakan di seluruh proses service
	db, err := database.Connect()
	if err != nil {
		lg.Println("Can't connect to db:", err.Error())
	}

	timeLocation := app.GetFixedTimeZone()
	// Variabel a akan digunakan sepanjang proses service
	// berisi koneksi database
	// dan data lain yang memungkinkan untuk digunakan secara berulang
	a := app.App{
		DB:           db,
		HttpClient:   &http.Client{},
		Name:         "CRUD Template Service",
		TimeLocation: timeLocation,
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Memanggil fungsi yang mengelola routing
	initRoute(a, e)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"status":${status},"method":"${method}","uri":"${uri}","time":"${time_rfc3339_nano}","x-member":"${header:X-Member}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"error":"${error}","latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))

	e.Use(middleware.Recover())
	if e.Debug {
		e.Logger.SetLevel(log.DEBUG)
	}

	// Menjalankan service di port 80
	e.Logger.Fatal(e.Start(":80"))
}
