create database max_inventory;
use max_inventory;

create table USERS (
    id int not null auto_increment,
    mail varchar(255) not null, 
    name varchar(50) not null,
    password varchar(50) not null,
    primary key(id)
);

create table PRODUCTS(
    id int not null auto_increment,
    description varchar(255) not null, 
    name varchar(50) not null,
    price float not null,
    create_by int not null,
    primary key(id),
    foreign key (create_by) references USERS(id)
);

create table roles(
    id not null auto_increment,
    name varchar(255) not null,
    primary key(id)
);

create table user_roles(
    id int not null auto_increment,
    user_id int not null, 
    role_id int not null,
    primary key(id),
    foreign key (user_id) references USERS(id),
    foreign key (role_id) references ROLES(id)
);

insert into roles(id, name) values (1, 'admin');
insert into roles(id, name) values (2, 'seller');
insert into roles(id, name) values (3, 'customer');