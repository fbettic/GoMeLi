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

1. Ejecutar la API

   *NOTA: no es necesario realizar los siguiente pasos (2 y 3) ya que el token y el refresh token ya han sido guardados con anterioridad en un json, y se actualizaran cuando intente hacer una consulta en caso de que el token este vencido, aun así puede ejecutar el siguiente paso si se desea probar la función de verificación de usuario, NO SE RECOMIANDA UTILIZAR UN USUARIO DISTINTO AL LOS AQUI PROPORCIONADOS YA QUE NO SE HA IMPLEMENTADO UNA FUNCION PARA GUARDAR LOS DATOS DEL MISMO LO QUE PUEDE OCACIONAR ERRORES)*

2. Entrar en la dirección http://localhost:8080/webtest y hacer click en Autorizar API

3. Iniciar sesión con el usuario Vendedor y dar autorización (basta con solo iniciar sesión ya que el usuario vendedor ya ha sido verificado con anterioridad)

4. luego de el paso anterior se lo redireccionará a home.html, si usted no realizo los pasos 2 y 3 puede entrar en http://localhost:8080/webtest/home.html aquí tendrá una lista de todas las funciones disponibles, solo elija una de ellas y obtendrá su resultado.

   *NOTA: en caso de que desee publicar un producto, necesitara rellenar la información del formulario primero, no se ha implementado ningún sistema para detectar si usted relleno o no los datos o no, por lo que puede ocasionar algún problema si no lo hace correctamente*

Lista de endpoints:

http://localhost:8080/webtest/oauth : Aquí se envía el Access Code de MeLi para luego ser canjeado por el Access Token, para probar este endpoint obligatoriamente debe acceder a la siguiente dirección http://localhost:8080/webtest y hacer click en el botón "Autorizar API"

http://localhost:8080/webtest/additem : Aqui se hace un POST request desde el front para publicar un nuevo producto, para probar este endpoint obligatoriamente debe acceder a la siguiente dirección http://localhost:8080/webtest/newproduct.html y rellenar el formulario con los datos que se indica *(no posee sistema para verificar si se han rellenado los campos por lo que si no se completan pueden presentarse errores)*

http://localhost:8080/webtest/itemlist : Al acceder a este endpoint se mostrara la lista de productos publicados por el vendedor, puede acceder directamente a este endpont y al acceder deberá mostrarse un json con la lista de productos

http://localhost:8080/webtest/soldlist : Al acceder a este endpoint se mostrara la lista de productos que han sido vendidos, puede acceder directamente a este endpont y al acceder deberá mostrarse un json con la lista de productos vendidos

http://localhost:8080/webtest/questlist : Al acceder a este endpoint se mostrara la lista de preguntas no respondidas por el vendedor, puede acceder directamente a este endpont y al acceder deberá mostrarse un json con la lista de preguntas

NOTA: Si realiza una consulta y no se muestra los datos deseados y en su lugar se muestra "Error hay que actualizar token, por favor intente nuevamente", solo recargue la ventana y se deberán mostrar los datos sin problemas, esto sucede cuando el token del usuario desde el cual hacemos la consulta se encuentra vencido, entonces, la API realiza la consulta con un token vencido, MeLi devuelve una respuesta informando de esto y la API al recibir este aviso envía un refresh token, luego de esto ya esta listo para realizan las peticiones nuevamente

