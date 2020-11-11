package api_back

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// En esta variable se guardara la lista de usuarios que se
// registraron en nuestra api con sus respectivos datos
var userList typeUserList

// user actual desde el cual se enviaran las consultas
// esto se quitara luego ya que luego la identificacion del cliente
// sera enviada en cada peticion que este haga
var User userData

// Estructura que guardara la lista de los datos de clientes
type typeUserList struct {
	UserList []userData 	`json:"user_list"`
}

type userData struct{
	Id int					`json:"id"`
	Nickname string			`json:"nickname"`
	AccessToken string		`json:"access_token"`
	RefreshToken string 	`json:"refresh_token"`
}

// Se carga la lista de usuarios (por ahora sera desde un json)
func ReadUserList() {

	userFile := []byte(openJson("./api_back/json/users.json"))

	err := json.Unmarshal(userFile, &userList)

	if err != nil{
		fmt.Println(err)
		fmt.Errorf("Error: ",err.Error())
		return
	}

	fmt.Println("Lista de usuarios cargada")
}

// se busca el usuario en la lista (luego sera una peticion a la base de datos)
func findUser(id int) (int,error) {

	for i:=0; i<len(userList.UserList); i++ {

		if userList.UserList[i].Id == id{
			return i, nil

		}
	}
	return -1, errors.New("Usuario no encontrado")
}

// se guarda el nuevo token y el nuevo refresh token
func SaveToken(id int, token string, refreshToken string) error{

	userPos,err := findUser(id)

	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Error: ",err.Error())
		return err
	}

	userList.UserList[userPos].RefreshToken = refreshToken
	userList.UserList[userPos].AccessToken = token

	b, err := json.Marshal(userList)

	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Error: ",err.Error())
		return err
	}

	// se actualiza el json de la lista de usuarios
	err = ioutil.WriteFile("./api_back/json/users.json",b,0644)

	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Error: ",err.Error())
		return err
	}

	// se cargan los nuevos datos en el usuario que se esta utilizando
	loadUserDataAt(userPos)

	return nil
}

// se busca el usuario que queremos cargar en la lista de usuarios
func LoadUserData(id int) error{

	userPos, err := findUser(id)

	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Error: ",err.Error())
		return err
	}

	loadUserDataAt(userPos)

	return nil
}

// se cargan los datos del usuario especificado
func loadUserDataAt(pos int) error{

	if pos > len(userList.UserList){
		return errors.New("Out of range")
	}
	User = userList.UserList[pos]

	AccessToken = User.AccessToken
	code = User.RefreshToken
	return nil
}