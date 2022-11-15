-- +goose Up
-- +goose StatementBegin
create table users(
    id serial primary key,
    name varchar(50),
    email varchar(100)
);

insert into users(name, email) values
    ('ivan', 'ivan@mail.ru'),
    ('andrey', 'andrey@gmail.com'),
    ('john', 'andrey@gmail.com'),
    ('slava', 'slava@example.com'),
    ('alex', 'alex@testserver')
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
