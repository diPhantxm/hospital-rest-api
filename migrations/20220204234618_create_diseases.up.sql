CREATE TABLE diseases(
    id int primary key identity(1, 1),
    diseaseName varchar(64) not null,
    treatment varchar(1024),
    startDate date not null,
    patientId int not null foreign key references patients(id)
);