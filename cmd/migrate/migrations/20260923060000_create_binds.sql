
-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS binds 
(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255)
);

-- +goose StatementEnd

-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS user_binds 
(
    user_id int NOT NULL,
    bind_id int NOT NULL,
    addr VARCHAR(255) NOT NULL,
    CONSTRAINT pk_user_binds PRIMARY KEY (user_id, bind_id),
    CONSTRAINT fk_user_binds_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_user_binds_bind
        FOREIGN KEY (bind_id)
        REFERENCES binds(id)
        ON DELETE CASCADE
);

-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS user_binds;
DROP TABLE IF EXISTS binds;


-- +goose StatementEnd