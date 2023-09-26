<?php
session_start();
include "conn.php";

if(isset($_POST["input_name"]) && isset($_POST["input_pass"])) {
    function validate($data) {
        $data = trim($data);
        $data = stripslashes($data);
        $data = htmlspecialchars($data);
        return $data;
    }
}

$uname = $_POST["input_name"];
$pass = $_POST["input_pass"];

if(empty($uname)) {
    header("Location: ../pages/login.html?error=Username required");
    exit();
} elseif(empty($pass)) {
    header("Location: ../pages/login.html?error=Password required");
    exit();
}

$sql = "SELECT * FROM users WHERE username='$uname' AND password='$pass'";
$result = mysqli_query($db_conn, $sql);

if(mysqli_num_rows($result) === 1) {
    $row = mysqli_fetch_assoc($result);
    if($row["username"] === $uname && $row["password"] === $pass) {
        echo "Logged into user: $uname | $pass";
        session_reset();
        $_SESSION["id"] = $row["id"];
        $_SESSION["username"] = $row["username"];
        $_SESSION["password"] = $row["password"];
        
        header("Location: ../php/home.php");
        exit();
    } 
} else {
    header("Location: ../pages/login.html?error=Incorrect username or password");
    exit();
}
?>
