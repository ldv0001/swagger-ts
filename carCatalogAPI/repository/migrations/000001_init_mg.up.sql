create table if not exists people(
id int generated always as identity,
name text,
surname text, 
patronymic text,
primary key(id)
);

create table if not exists cars(
reg_num varchar(10),
mark text,
model text,
year int,
primary key(reg_num),
people_id int,
foreign key( people_id) references people(id) 
);