<?php 
  include( "../gomeli/bd.php" );

  if (!isset($_SESSION['emp_id'])){
    header("Location: ../gomeli/index.php");
  }
  
?>
<!DOCTYPE html>    
<html>
 <head>
  <link type="text/css" rel="stylesheet" href="style.css">
  <title>MELIAPI</title>
  <link rel="icon" href="https://tse3.mm.bing.net/th?id=OIP.CQXGnz_bUjVSM_btcyZfmgAAAA&pid=Api" type="image/png" sizes="16x16">
  <center>
  <img src="https://tse3.mm.bing.net/th?id=OIP.CQXGnz_bUjVSM_btcyZfmgAAAA&pid=Api" width="70" heigth="70" alt="" class="d-inline-block align-top">
  </center>
  <!-- La linea de arriba es para importar estilos CSS a nuestro formulario -->
  <title>Formulario de contacto</title>
  <script src="https://kit.fontawesome.com/5fd270c7b7.js" crossorigin="anonymous"></script>
  <script type="text/javascript">
    function redireccionar(){
      window.locationf="home.php";
    } 
  </script>

 </head>
 <body>
    <font face="arial" font size=2>
      <h1 align = "center" style="color:cadetblue;" > Elegí el tipo de publicación </h1>
    </font>
    <section class="formulario">
      <!-- Formulario -->
      <!-- <form name="formulario" method="post" action="http://localhost:8080/gomeli/additem?id=1"> -->
        <form name="formulario" method="post" action="acciones.php?add=1">
          <label for="nombre">Nombre del producto:</label>
          <input id="title" type="text" name="title" placeholder="Nombre del producto" required="" />
          
          <p>Precio $: <input type="number" name="price" placeholder="1234" required="" ></p>
          <p>Stock: <input type="number" name="available_quantity" placeholder="1234" required=""></p>

          <p>Condicion del producto: <select name="condition">
              <option value="new">Nuevo</option>
              <option value="used">Usado</option>
          </select></p>
          <input id="submit" type="submit" name="submit" value="Enviar" />
          
      </form>

    </section>
    <a href="home.php" class="btvolver" >Volver</a>
  <!-- Scripts de Bootstrap -->
        <!-- JQuery -->
        <script src="https://code.jquery.com/jquery-3.5.1.slim.min.js" 
            integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" 
            crossorigin="anonymous">
        </script>
        <!-- Popper y JavaScript-->
        <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/js/bootstrap.bundle.min.js" 
                integrity="sha384-ygbV9kiqUc6oa4msXn9868pTtWMgiQaeYH7/t7LECLbyPA2x65Kgf80OJFdroafW" 
                crossorigin="anonymous">
        </script>
  </body>
              
</html>