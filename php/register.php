<?php
include "conn.php";

if(isset($_POST["input_name"]) && isset($_POST["input_pass"]) && isset($_POST["input_pass_confirm"])) {
    function validate($data) {
        $data = trim($data);
        $data = stripslashes($data);
        $data = htmlspecialchars($data);
        return $data;
    }
}

$uname = validate($_POST["input_name"]);
$pass = validate($_POST["input_pass"]);
$pass_confirm = validate($_POST["input_pass_confirm"]);

if(empty($uname)) {
    header("Location: ../pages/register.html?error=Username required");
    exit();
} elseif(empty($pass)) {
    header("Location: ../pages/register.html?error=Password required");
    exit();
} elseif(empty($pass_confirm)) {
    header("Location: ../pages/register.html?error=Password confirmation required");
    exit();
} elseif ($pass != $pass_confirm) {
    header("Location: ../pages/register.html?error=Passwords must match");
    exit();
}

$sql = "INSERT INTO users(username, password) values ($uname, $pass)";
?>
