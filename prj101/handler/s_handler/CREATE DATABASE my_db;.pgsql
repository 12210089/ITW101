

CREATE DATABASE HospitalRegistration;
CREATE TABLE UserData(
    UserName varchar(45) not null,
    Email varchar(45) not null,
    UserPassword varchar(45) not NULL,
    PRIMARY KEY(Email),
    UNIQUE(Email)
);



create table PatientDetails(
    PatientName varchar(45) not null,
    CID varchar(45) not null,
    Age varchar(45) not null,
    Address varchar(45) not null,
    Disease varchar(45) not null,
    Date varchar(45) not null,
    primary key(CID)
);

-- drop table PatientDetails

-- truncate table PatientDetails;
-- truncate table userdata;


