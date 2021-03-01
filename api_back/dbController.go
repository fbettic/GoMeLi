package api_back

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type InsertUserData struct{
	IdMeli int
	Nickname string
	Password string
	AccessToken string
	RefreshToken string
}

type ExportProducts struct {
	Title string
	Quantity int
	Price float64
	IdUser int
}

type ReqUserData struct{
	IdMeli int
	AccessToken string
	RefreshToken string
}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "vendedores"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

//func InsertarUsuario(user InsertUserData) (e error) {
//	db, err := obtenerBaseDeDatos()
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
//	// Preparamos para prevenir inyecciones SQL
//	sentenciaPreparada, err := db.Prepare("INSERT INTO usuarios (id_meli, nickname, password, access_token, refresh_token) VALUES( ?, ?, ?, ?, ?)")
//	if err != nil {
//		return err
//	}
//	defer sentenciaPreparada.Close()
//
//	// Ejecutar sentencia, un valor por cada '?'
//	_, err = sentenciaPreparada.Exec(user.IdMeli, user.Nickname, user.Password, user.AccessToken, user.RefreshToken)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func obtenerDatosUsuario (id string) (ReqUserData, error) {
	user:= ReqUserData{}

	db, err := obtenerBaseDeDatos()
	if err != nil {
		return ReqUserData{}, err
	}
	defer db.Close()

	resultado, err := db.Query("SELECT id_meli, access_token, refresh_token FROM usuarios WHERE id_user = " + id)

	if err != nil {
		fmt.Println("error en query select")
		return ReqUserData{}, err
	}
	// Si llegamos aquí, significa que no ocurrió ningún error
	defer resultado.Close()

	for resultado.Next() {
		err = resultado.Scan(&user.IdMeli, &user.AccessToken, &user.RefreshToken)
	}
	if err != nil {
		return ReqUserData{}, err
	}

	//fmt.Println(user)

	return user, nil
}


func actualizarToken(user ReqUserData) error {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("UPDATE usuarios SET access_token = ?, refresh_token = ? WHERE id_meli = ?")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer sentenciaPreparada.Close()
	// Pasar argumentos en el mismo orden que la consulta
	_, err = sentenciaPreparada.Exec(user.AccessToken, user.RefreshToken, user.IdMeli)
	return err // Ya sea nil o sea un error, lo manejaremos desde donde hacemos la llamada
}

func exportarProductos(itemList string) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	resp, err := db.Query("TRUNCATE TABLE productos;")
	if err != nil {
		return err
	}
	defer resp.Close()

	// Preparamos para prevenir inyecciones SQL
	_, err = db.Query("INSERT INTO productos (title, quantity, price, id_user) VALUES" + itemList)

	if err != nil {
		return err
	}

	return nil
}