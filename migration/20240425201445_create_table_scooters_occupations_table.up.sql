USE scootin_aboot;

CREATE TABLE scooters_occupations(
    id BINARY(16) NOT NULL DEFAULT(UUID_TO_BIN(UUID())),
    scooter_id BINARY(16) NOT NULL,
    user_id BINARY(16) NOT NULL,
    PRIMARY KEY(id),
    UNIQUE KEY unique_scooter_id(scooter_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE `utf8mb4_unicode_ci` ENGINE = InnoDB;
