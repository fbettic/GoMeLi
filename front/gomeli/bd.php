<?php

    @session_start();
    if (!isset($_SESSION['tiempo'])) {
        $_SESSION['tiempo']=time();
    }
    else if (time() - $_SESSION['tiempo'] > 1800) {
        session_unset();
        /* Aquí redireccionas a la url especifica */
        header("Location: ../gomeli/index.php"); 
    }
    $_SESSION['tiempo']=time(); //Si hay actividad seteamos el valor al tiempo actual

    $conn = mysqli_connect(
        '127.0.0.1',
        'root',
        '',
        'vendedores'

    );



?>