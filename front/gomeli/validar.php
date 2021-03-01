<?php 
    include( "../gomeli/bd.php" );
    unset($_SESSION['emp_id']);
    
    $page = "index.php";

    $accion = $_GET['accion'];

    if (isset($_GET['tabla'])){
        $tabla = $_GET['tabla'];
    }

    if ($accion == 'validar'){

        $user = $_POST['username'];
        $pass = $_POST['password'];

        $query = "SELECT * FROM usuarios WHERE nickname = '$user'";
        $result = mysqli_query($conn,$query);
        while ($row = mysqli_fetch_array($result)){
            if ($pass == $row['password']){
                $_SESSION['emp_id'] = $row['id_user'];
                $_SESSION['mensaje'] = 'Bienvenido';
                $_SESSION['tipo_mensaje'] = 'success';
                if (isset($_GET['page'])){
                    $page = $_GET['page'];
                    $_SESSION['mensaje'] = 'Bienvenido nuevamente';
                    $_SESSION['tipo_mensaje'] = 'success';
                    header("Location: $page");
                } else {
                    $_SESSION['mensaje'] = 'Bienvenido nuevamente';
                    $_SESSION['tipo_mensaje'] = 'success';
                    header("Location: ../gomeli/home.php");
                }
                
            } else {
                $_SESSION['mensaje'] = 'Los datos ingresados no son válidos. Intente nuevamente';
                $_SESSION['tipo_mensaje'] = 'danger';
                header("Location: ../gomeli/index.php");
            }
        }
    } elseif($accion == 'nuevo'){

        $nombre = $_POST['nombre'];
        $usuario = $_POST['usuario'];
        $pass = $_POST['contraseña'];

        $hash = password_hash($pass, PASSWORD_DEFAULT);

        $query = "INSERT INTO $tabla (nombre,usuario,pass) VALUES ('$nombre','$usuario','$hash')";
        $result = mysqli_query($conn,$query);
        
        if (!$result){
                $_SESSION['mensaje'] = 'Error cargando los datos solicitados';
                $_SESSION['tipo_mensaje'] = 'danger';          
        } else {
                $_SESSION['mensaje'] = 'Empleado cargado correctamente. Ya puede iniciar sesión';
                $_SESSION['tipo_mensaje'] = 'success';
        }
        header("Location: $pag");
    }
?>