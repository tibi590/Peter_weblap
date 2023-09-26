<?php
session_start();

echo "Id: ".$_SESSION['id']."Username: ".$_SESSION['username']." | Password: ".$_SESSION['password'];
?>
