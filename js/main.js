function login() {
    let username = document.getElementById("input_name").value;
    let password = document.getElementById("input_pass").value;
    
    console.log(username);
    console.log(password);
}
 
function register() {
    let username = document.getElementById("input_name").value;
    let password = document.getElementById("input_pass").value;
    let password_confirm = document.getElementById("input_pass_confirm").value;
    let email = document.getElementById("input_email").value;

    console.log(username);
    console.log(password);
    console.log(password_confirm);
    console.log(password === password_confirm);
    console.log(email);
}
