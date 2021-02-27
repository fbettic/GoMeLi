package api_back

import (
	_ "github.com/go-sql-driver/mysql"
)

//type UserData struct{
//	Id int					`json:"id"`
//	Nickname string			`json:"nickname"`
//	AccessToken string		`json:"access_token"`
//	RefreshToken string 	`json:"refresh_token"`
//}
//
//
//func obtenerBaseDeDatos() (db *sql.DB, e error) {
//	usuario := "root"
//	pass := ""
//	host := "tcp(127.0.0.1:3306)"
//	nombreBaseDeDatos := "agenda"
//	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
//	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
//	if err != nil {
//		return nil, err
//	}
//	return db, nil
//}
//
//func Insertar(user userData) (e error) {
//	db, err := obtenerBaseDeDatos()
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
//	// Preparamos para prevenir inyecciones SQL
//	sentenciaPreparada, err := db.Prepare("INSERT INTO agenda (nombre, direccion, correo_electronico) VALUES(?, ?, ?)")
//	if err != nil {
//		return err
//	}
//	defer sentenciaPreparada.Close()
//	// Ejecutar sentencia, un valor por cada '?'
//	_, err = sentenciaPreparada.Exec(c.Nombre, c.Direccion, c.CorreoElectronico)
//	if err != nil {
//		return err
//	}
//	return nil
//}