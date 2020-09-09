-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE weathers (
    id SERIAL NOT NULL PRIMARY KEY,
    weather_id INT NOT NULL,
    main VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    hourly_weather_info_id INT REFERENCES hourly_weather_infos NOT NULL,

    CREATED_AT TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UPDATED_AT TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    DELETED_AT TIMESTAMP WITH TIME ZONE
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE weathers;
