-- +goose Up
CREATE TABLE users (
    id bigint(20) NOT NULL AUTO_INCREMENT, PRIMARY KEY (id),
    email varchar(50) NOT NULL,
    password varchar(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS users;
