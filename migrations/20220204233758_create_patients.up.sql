CREATE TABLE patients(
    id int primary key identity(1, 1),
    firstName varchar(32) not null,
    lastName varchar(32) not null,
    birthDate date not null,
    residence varchar(64) not null
);