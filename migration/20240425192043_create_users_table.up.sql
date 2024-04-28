USE scootin_aboot;

CREATE TABLE users(
    id BINARY(16) NOT NULL DEFAULT(UUID_TO_BIN(UUID())),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    api_key VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE `utf8mb4_unicode_ci` ENGINE = InnoDB;