USE scootin_aboot;

CREATE TABLE scooters_occupations(
    id BINARY(16) NOT NULL,
    scooter_id BINARY(16) NOT NULL,
    user_id BINARY(16) NOT NULL,
    occupied_at DATETIME NOT NULL,
    released_at DATETIME DEFAULT NULL,
    PRIMARY KEY(id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE `utf8mb4_unicode_ci` ENGINE = InnoDB;
