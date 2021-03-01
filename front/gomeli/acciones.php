<?php
    include( "../gomeli/bd.php" );

    if (!isset($_SESSION['emp_id'])){
        header("Location: ../gomeli/index.php");
    }

    if (isset($_GET['export'])){
       
        
        $ch = curl_init("localhost:8080/gomeli/export?id=1"); 
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_HEADER, 0);
        $data = curl_exec($ch);
        $data = json_decode($data);
        curl_close($ch);

        $_SESSION['mensaje'] = $data->message;
        if ($data->status != 200){
            $_SESSION['tipo_mensaje'] = "danger";
        } else {
            $_SESSION['tipo_mensaje'] = "success";
        }
        header("Location: home.php");
        
    }

    if (isset($_GET['add'])){

        $data = array('title' => $_POST['title'],
                    'price' => $_POST['price'],
                    'available_quantity' => $_POST['available_quantity'],
                    'condition' => $_POST['condition']);

        $data = http_build_query($data);
       
        
        $crl = curl_init('http://localhost:8080/gomeli/additem?id=1');
        curl_setopt($crl, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($crl, CURLOPT_HEADER, 0);
        curl_setopt($crl, CURLOPT_POST, true);
        curl_setopt($crl, CURLOPT_POSTFIELDS, $data);

        $result = curl_exec($crl);
        curl_close($crl);
        $data = json_decode($result);
        $_SESSION['mensaje'] = $data->message;
        if ($data->status != 200){
            $_SESSION['tipo_mensaje'] = "danger";
        } else {
            $_SESSION['tipo_mensaje'] = "success";
        }
        header("Location: home.php");
        
    }

    if (isset($_GET['resp'])){
       
        $id = $_GET['resp'];
        $url = "http://localhost:8080/gomeli/answer?id=1&idq=$id";
        $data =array('answer' => $_POST['respuesta']);
        $data = http_build_query($data);
        $crl = curl_init($url);
        curl_setopt($crl, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($crl, CURLOPT_HEADER, 0);
        curl_setopt($crl, CURLOPT_POST, true);
        curl_setopt($crl, CURLOPT_POSTFIELDS, $data);

        $result = curl_exec($crl);
        curl_close($crl);
        $data = json_decode($result);
        $_SESSION['mensaje'] = $data->message;
        if ($data->status != 200){
            $_SESSION['tipo_mensaje'] = "danger";
        } else {
            $_SESSION['tipo_mensaje'] = "success";
        }
        header("Location: home.php");
    }

    
?>