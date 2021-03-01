# GoMeLi
Proyecto creado para la materia LAB3

Instalación del paquete:

```
go get -u github.com/Chino976/GoMeLi
```

Usuarios a utilizar :

- Vendedor:

  ```json
  {
      "id": 666272328,
      "nickname": "TETE5197999",
      "password": "qatest8481",
      "site_status": "active",
      "email": "test_user_82019645@testuser.com"
  }
  ```

- Comprador:

  ```json
  {
      "id": 671233552,
      "nickname": "TESTTPKDVOE8",
      "password": "qatest9127",
      "site_status": "active",
      "email": "test_user_17215619@testuser.com"
  }
  ```

Lista de endpoints:

http://localhost:8080/gomeli/oauth : Aquí se envía el Access Code de MeLi para luego ser canjeado por el Access Token, para probar este endpoint obligatoriamente debe acceder a la siguiente dirección http://localhost:8080/webtest y hacer click en el botón "Autorizar API"

http://localhost:8080/gomeli/additem : Aqui se hace un POST request desde el front para publicar un nuevo producto, para probar este endpoint obligatoriamente debe acceder a la siguiente dirección http://localhost:8080/webtest/newproduct.html y rellenar el formulario con los datos que se indica 

http://localhost:8080/gomeli/home : Al realizar un GET a este endpoint se recibirá una respuesta json con los datos de:

- Lista de Productos publicados
- Lista de productos vendidos
- Lista de preguntas sin responder

http://localhost:8080/gomeli/export : Al realizar un GET a este endpoint se realizara la carga de la lista de todos los productos publicados a una tabla en la base de datos que usted asocie

http://localhost:8080/gomeli/answer : Al realizar un POST con el ID de la pregunta que desea responder mas el mensaje a contestar este será enviado automáticamente en mercado libre

