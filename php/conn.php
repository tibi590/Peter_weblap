<?php
$db_host = "localhost";
$db_user = "admin";
$db_pass = "admin";
$db_name = "peter";

$db_conn = mysqli_connect($db_host, $db_user, $db_pass, $db_name);

if($db_conn -> errno) {
    echo "Connection to the database failed";
    exit();
}
?>
