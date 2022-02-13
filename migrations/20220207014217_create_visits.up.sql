CREATE TABLE visits(
    id int primary key identity(1, 1),
    patientId int not null foreign key references patients(id),
    diseaseId int not null foreign key references diseases(id),
    doctorId int not null foreign key references doctors(id),
    visitDate date not null
);