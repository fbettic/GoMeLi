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


Como utilizar el paquete:

1. Ejecutar el paquete

2. Entrar en la dirección http://localhost:80/webtest y hacer click en Autorizar API

   ![image-20201112203723937](C:\Users\feder\AppData\Roaming\Typora\typora-user-images\image-20201112203723937.png)

3. Iniciar sesión con el usuario Vendedor y dar autorización (basta con solo iniciar sesión ya que el usuario vendedor ya ha sido verificado con anterioridad)

4. luego de el paso anterior se lo redireccionará a http://localhost:80/webtest/home.html aquí tendrá una lista de todas las funciones disponibles, solo elija una de ellas y obtendrá su resultado 

   ![image-20201112203632056](C:\Users\feder\AppData\Roaming\Typora\typora-user-images\image-20201112203632056.png)

   NOTA: en caso de que desee publicar un producto, necesitara rellenar la información del formulario primero 

Lista de endpoints:

http://localhost:80/webtest/oauth : Aqui se envia el Access Code de MeLi para luego ser canjeado por el Access Token

http://localhost:80/webtest/additem : Aqui se hace un POST request desde el front para publicar un nuevo producto

http://localhost:80/webtest/itemlist : Al acceder a este endpoint se mostrara la lista de productos publicados por el vendedor

http://localhost:80/webtest/soldlist : Al acceder a este endpoint se mostrara la lista de productos que han sido vendidos 

http://localhost:80/webtest/questlist : Al acceder a este endpoint se mostrara la lista de preguntas no respondidas por el vendedor