CREATE DATABASE crud;


use crud;


CREATE Table usuario (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nome VARCHAR(50) not NULL,
    idade CHAR(3) not NULL ,
    cpf VARCHAR(15) not null UNIQUE,
    telefone VARCHAR(11) not NULL UNIQUE

)ENGINE=INNODB ;
