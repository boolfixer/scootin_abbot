USE scootin_aboot;

INSERT INTO scooters (id, name, latitude, longitude)
VALUES
    (UUID_TO_BIN(UUID()), 'Scooter 1', 2, 4),
    (UUID_TO_BIN(UUID()), 'Scooter 2', 3, 7),
    (UUID_TO_BIN(UUID()), 'Scooter 3', 3, 9),
    (UUID_TO_BIN(UUID()), 'Scooter 4', 5, 15),
    (UUID_TO_BIN(UUID()), 'Scooter 5', 6, 5),
    (UUID_TO_BIN(UUID()), 'Scooter 6', 7, 11),
    (UUID_TO_BIN(UUID()), 'Scooter 7', 9, 14),
    (UUID_TO_BIN(UUID()), 'Scooter 8', 11, 11),
    (UUID_TO_BIN(UUID()), 'Scooter 9', 11, 18),
    (UUID_TO_BIN(UUID()), 'Scooter 10', 13, 2),
    (UUID_TO_BIN(UUID()), 'Scooter 11', 14, 13),
    (UUID_TO_BIN(UUID()), 'Scooter 12', 17, 17),
    (UUID_TO_BIN(UUID()), 'Scooter 13', 18, 5),
    (UUID_TO_BIN(UUID()), 'Scooter 14', 18, 9),
    (UUID_TO_BIN(UUID()), 'Scooter 15', 19, 3);