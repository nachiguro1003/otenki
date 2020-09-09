-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE hourly_weather_infos (
    id SERIAL NOT NULL PRIMARY KEY,
    date INT NOT NULL UNIQUE ,
    temperature FLOAT NOT NULL,

    CREATED_AT TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UPDATED_AT TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    DELETED_AT TIMESTAMP WITH TIME ZONE
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE hourly_weather_infos;
