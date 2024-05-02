USE scootin_aboot;

CREATE TABLE scooters(
    id BINARY(16) NOT NULL DEFAULT(UUID_TO_BIN(UUID())),
    name VARCHAR(255) NOT NULL,
    latitude SMALLINT UNSIGNED NOT NULL,
    longitude SMALLINT UNSIGNED NOT NULL,
    location_updated_at DATETIME NOT NULL,
    PRIMARY KEY(id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE `utf8mb4_unicode_ci` ENGINE = InnoDB;
