<?php
    if(isset($_POST["back"])) {
        header("Location: ../index.html");
    }  elseif (isset($_POST["login"])) {
        echo "login";;
    } else {
        echo "action missing";
    }
?>
