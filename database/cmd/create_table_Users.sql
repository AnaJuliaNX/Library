CREATE TABLE Users (
    `id` BINARY(16) NOT NULL,
    `Name` VARCHAR(100) NOT NULL,
    `Email` VARCHAR(100),
    `Password` VARCHAR(100) NOT NULL,
    `owner` BOOLEAN NOT NULL DEFAULT FALSE,
    `Phone` VARCHAR(20),
    `DateCreated` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRiMARY KEY (`id`)
)
ENGINE = InnoDB;