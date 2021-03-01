<?php 
    include( "../gomeli/bd.php" );

    if (!isset($_SESSION['emp_id'])){
        header("Location: ../gomeli/index.php");
    }

    
    $ch = curl_init("localhost:8080/gomeli/home?id=1"); 
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_HEADER, 0);
    $data = curl_exec($ch);
    $data = json_decode($data);
    curl_close($ch);

    $itemsVendidos = $data->sold_item_list;
    $items = $data->item_list;
    $pregSinResp = $data->quest_list;

?>

<!DOCTYPE html>
<html lang='es' style="margin-bottom: 100px;">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>MELIAPI</title>

        <link rel="icon" href="https://tse3.mm.bing.net/th?id=OIP.CQXGnz_bUjVSM_btcyZfmgAAAA&pid=Api" type="image/png" sizes="16x16">

        <!-- Bootstrap. Para obtener dirigirse a getbootstrap.com -->
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-giJF6kkoqNQ00vy+HMDP7azOuL0xtbfIcaT9wjKHr8RbDVddVHyTfAAsrekwKmP1" crossorigin="anonymous">
        <script src="https://kit.fontawesome.com/5fd270c7b7.js" crossorigin="anonymous"></script>
    </head>

    <body>
        <nav class="navbar navbar-light" style="background-color:#fdff00;" >
            <div class="container">
                <a href="home.php" class="navbar-brand"><img src="https://tse3.mm.bing.net/th?id=OIP.CQXGnz_bUjVSM_btcyZfmgAAAA&pid=Api" width="50" heigth="50" alt="" class="d-inline-block align-top"><font size="5" color="#000000">MELIAPI</font></a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarToggleExternalContent" aria-controls="navbarToggleExternalContent" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
            </div>
        </nav>
        <div class="collapse dropdown-menu-light" id="navbarToggleExternalContent" style="position: absolute; z-index: 101; width: 100%;">
            <div class="p-4" style="background-color: #fdff00; ">
                <hr class="text-white dropdown-divider">
                <a class="dropdown-item" href="../gomeli/logout.php">Logout</a>
            </div>
        </div>
        <div class="container py-4">
            <div class="row">
                <div class="col-xl">
                    <?php if(isset($_SESSION['mensaje'])){?>

                        <div class="alert alert-<?php echo $_SESSION['tipo_mensaje']?> alert-dismissible fade show" role="alert">
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
        <div class="container border border-primary">
            <div style="height: 500px;">
                <div class="container pt-4">
                    <div class="row">
                        <div class="col-xl">
                            <!-- LISTA DE PRODUCTOS -->
                            <div class="table-responsive h-50">
                                <table class="table border table-hover">
                                    <thead class="sticky-top">
                                        <th>Imagen</th>
                                        <th>Artículo</th>
                                        <th>Cantidad</th>
                                        <th>Precio</th>
                                    </thead>
                                    <tbody>
                                        <?php for ($i=0;$i<count($items);$i++){?>
                                        <tr>
                                            <td style="width:30%; height:60px"><img src=<?php echo $items[$i]->picture ?> width="50" heigth="50" alt="" class="d-inline-block align-top"></td>
                                            <td><?php echo $items[$i]->title ?></td>
                                            <td><?php echo $items[$i]->quantity ?></td>
                                            <td>$<?php echo $items[$i]->price ?></td>
                                        </tr>
                                        <?php } ?>
                                    </tbody>
                                </table>
                            </div>
                            <a href="artcarga.php" class="btn btn-block" style="background-color: #fdff00; color: black;">Agregar Artículo</a>
                        </div>
                        <div class="col">
                            <!-- LISTA DE PREGUNTAS SIN RESPONDER -->
                            <div class="table-responsive h-50">
                                <table class="table border table-hover table-bordered">
                                    <thead>
                                        <th>Preguntas Sin Responder</th>
                                    </thead>
                                    <tbody>
                                        <?php for ($i=0;$i<count($pregSinResp);$i++){?>
                                            <tr>
                                                <td>
                                                    <div class="row-md">
                                                        <p>Producto: <?php echo $pregSinResp[$i]->item_title ?></p>
                                                    </div>
                                                    <div class="row-md">
                                                        <p>Pregunta: <?php echo $pregSinResp[$i]->text ?></p>
                                                    </div>
                                                    <?php 
                                                        if( isset($_GET['resp']) && $_GET['resp'] == $pregSinResp[$i]->item_title ){?>
                                                            <div class="row-md justify-content-end">
                                                                <div class="col align-self-end">
                                                                    <form action="acciones.php?resp=<?=$pregSinResp[$i]->id?>" method="post">
                                                                        <input type="text" name="respuesta">
                                                                        <input class="btn btn-primary" type="submit" Value="Enviar">
                                                                    </form>
                                                                </div>
                                                            </div>
                                                        <?php } else {?>
                                                            <div class="row-md justify-content-end">
                                                                <div class="col align-self-end">
                                                                    <a href="home.php?resp=<?=$pregSinResp[$i]->item_title?>" class="btn btn-primary">Responder</a>
                                                                </div>
                                                            </div>
                                                        <?php }
                                                    ?>
                                                </td>
                                            </tr>
                                        <?php } ?>
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="container pt-4 border-top border-primary">
                <div class="row">
                    <div class="col">
                    <!-- LISTA DE PRODUCTOS VENDIDOS -->
                        <div class="table-responsive">
                            <table class="table border table-bordered">
                                <thead><th>Artículos Vendidos</th><th></th><th></th></thead>
                                <thead>
                                    <th>Artículo</th>
                                    <th>Cantidad</th>
                                    <th>Ganancia del artículo</th>
                                </thead>
                                <tbody>
                                    <?php for ($i=0;$i<count($itemsVendidos);$i++){ ?>
                                        <tr>
                                            <td><?php echo $itemsVendidos[$i]->item ?></td>
                                            <td><?php echo $itemsVendidos[$i]->quantity ?></td>
                                            <td><?php echo $itemsVendidos[$i]->total_paid_amount?></td>
                                        </tr>
                                    <?php } ?>
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="container py-4">
            <div class="row">
                <div class="d-grid gap-2 d-md-flex justify-content-md-end">
                    <a href="acciones.php?export=1" class="btn btn-sm"  style="background-color: #59359a; color: white;">Exportar Todos los Datos</i></a>
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
    </body>
    <footer>
        <div class="footer" style="position: fixed; bottom: 0; right: 0; left: 0; background-color: #fdff00; color: white;">
            <div class="container text-center">
                <p>Autor: Ferreyra Tomás</p>
            </div>
        </div>
    </footer>
</html>
