PRAGMA foreign_key = "ON";

CREATE TABLE shelf (
id integer primary key,
name varchar
);

CREATE TABLE game (
id integer primary key, 
name varchar,
shelf_id references shelf(id) on update cascade on delete cascade
);
