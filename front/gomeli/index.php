<?php include( "../gomeli/bd.php" ) ?>
<?php
    $page = "../gomeli/home.php";
    if (isset($_GET['page'])){
        $page = $_GET['page'];
    }
    if (isset($_SESSION['emp_id'])){
        header("Location: $page");
    }
?>

<!DOCTYPE html>
<html lang='es'>
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>MELIAPI</title>
        <link rel="icon" href="https://tse3.mm.bing.net/th?id=OIP.CQXGnz_bUjVSM_btcyZfmgAAAA&pid=Api" type="image/png" sizes="16x16">
        <link rel="stylesheet" href="login.css">
        <!-- Bootstrap. Para obtener dirigirse a getbootstrap.com -->
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-giJF6kkoqNQ00vy+HMDP7azOuL0xtbfIcaT9wjKHr8RbDVddVHyTfAAsrekwKmP1" crossorigin="anonymous">

        <script src="https://kit.fontawesome.com/5fd270c7b7.js" crossorigin="anonymous"></script>
    </head>
    <body>
        <div class="container py-4">
            <div class="row">
                <div class="col-xl">
                    <?php if(isset($_SESSION['mensaje'])){?>

                        <div class="alert alert-<?=$_SESSION['tipo_mensaje']?> alert-dismissible fade show" role="alert">
                            <?= $_SESSION['mensaje'] ?>
                            <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
                        </div>

                    <?php 
                        unset($_SESSION['tipo_mensaje']);
                        unset($_SESSION['mensaje']);
                     } ?>
                </div>
            </div>
        </div>

        <div class="container h-80">
            <div class="row align-items-center h-100">
                <div class="col-sm-auto mx-auto">
                    <div class="text-center">
                        <img id="profile-img" class="rounded-circle profile-img-card" src="https://tse3.mm.bing.net/th?id=OIP.CQXGnz_bUjVSM_btcyZfmgAAAA&pid=Api" />
                        <h1 id="profile-name" class="profile-name-card">MELIAPI</h1>

                        <div class="accordion" id="acordeon">

                            <div class="accordion-item">
                                <h3 class="accordion-header" id="headLogin">
                                    <button class="accordion-button btn-light btn-sm" type="button" data-bs-toggle="collapse" data-bs-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
                                        Login
                                    </button>
                                </h3>
                                <div id="collapseOne" class="accordion-collapse collapse show" aria-labelledby="headLogin" data-bs-parent="#acordeon">
                                    <div class="accordion-body">

                                        <form  class="form-signin mb-3" action="../gomeli/validar.php?accion=validar&page=<?=$page?>" method="POST">

                                            <div class="mb-3">
                                                <input type="text" name="username" id="inputUsername" class="form-control form-group" placeholder="usuario" required autofocus autocomplete="off">
                                            </div>
                                            
                                            <div class="mb-3">
                                                <input type="password" name="password" id="inputPassword" class="form-control form-group" placeholder="contraseña" required autofocus>
                                            </div>
                                            
                                            <button class="btn btn-lg btn-block btn-signin" style="background-color: #fdff00; color: black;" type="submit">Ingresar</button>

                                        </form>

                                    </div>
                                </div>
                            </div>

                            <div class="accordion-item">
                                <h3 class="accordion-header" id="headRegister">
                                    <button class="accordion-button collapsed  btn-light btn-sm" type="button" data-bs-toggle="collapse" data-bs-target="#collapseTwo" aria-expanded="false" aria-controls="collapseTwo">
                                        Registrar
                                    </button>
                                </h3>
                                <div id="collapseTwo" class="accordion-collapse collapse" aria-labelledby="headRegister" data-bs-parent="#acordeon">
                                    <div class="accordion-body">

                                        <form  class="form-signin mb-3" action="../gomeli/validar.php?accion=nuevo&tabla=empleados&pag=../gomeli/index.php" method="POST">

                                            <div class="mb-3">
                                                <input type="text" name="nombre" id="inputName" class="form-control form-group" placeholder="nombre" required autofocus autocomplete="off">
                                            </div>
                                            <div class="mb-3">
                                                <input type="text" name="usuario" id="inputUsername" class="form-control form-group" placeholder="usuario" required autofocus autocomplete="off">
                                            </div>
                                            <div class="mb-3">
                                                <input type="password" name="contraseña" id="inputPassword" class="form-control form-group" placeholder="contraseña" required autofocus>
                                            </div>
                                            <button class="btn btn-lg btn-block btn-signin" style="background-color: #fdff00; color: black;" type="submit">Registrarse</button>
                                            
                                        </form>

                                    </div>
                                </div>
                            </div>

                        </div>

                    </div>
                </div>
            </div>
        </div>

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
</html>