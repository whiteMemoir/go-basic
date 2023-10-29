 create table employees(
 	id serial primary key,
 	nip varchar(50) unique not null,
 	name varchar(100) not null,
 	address varchar(255),
 	created_at timestamp default now(),
 	updated_at timestamp default now()
 );
 
select * from employees;

CREATE TABLE "menus" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar(100) not null,
    "category" varchar(100) not null,
    "description" varchar(255),
    "price" int not null,
    image_url varchar(255) not null,
    "created_at" timestamp default now(),
    "updated_at" timestamp default now()
);

CREATE TABLE "transactions" (
    "id" SERIAL PRIMARY KEY,
    "employee_id" int,
    "menu_id" int,
    "quantity" int,
    "total" int,
    "created_at" timestamp,
    "updated_at" timestamp
);

alter table menus 
add coba varchar(50);

alter table menus 
drop coba;

alter table transactions 
add foreign key (menu_id) references menus(id);

alter table transactions 
add foreign key (employee_id) references employees(id);

create table test (
	id serial primary key,
	name varchar(100) not null,
	created_at timestamp default now(),
	updated_at timestamp default now()
);

drop table test ;
