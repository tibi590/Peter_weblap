CREATE TABLE users (
    id int(1) PRIMARY KEY auto_increment,
    username Varchar(64) NOT null unique,
    password Varchar(64) NOT null
);
