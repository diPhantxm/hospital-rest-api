CREATE TABLE doctors(
    id int primary key identity(1, 1),
    firstName varchar(32) not null,
    lastName varchar(32) not null,
    specialty varchar(32) not null
);