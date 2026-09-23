
-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS services 
(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id int NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    login VARCHAR(255) NOT NULL,
    pass varchar(255) NOT NULL,
    added_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT fk_services_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS services;

-- +goose StatementEnd