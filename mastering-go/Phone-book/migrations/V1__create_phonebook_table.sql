DROP TABLE IF EXISTS phone_book;

CREATE TABLE phone_book (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar(100) NOT NULL,
    surname varchar(255),
    phone_number varchar(11) NOT NULL
);