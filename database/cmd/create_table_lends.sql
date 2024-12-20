CREATE TABLE Lend_books (
    `id` BINARY(16) NOT NULL,
    `User_id` BINARY(16) NOT NULL,
    `Book_id` BINARY(16) NOT NULL,
    `Status` TINYINT NOT NULL COMMENT 'Status: emprestado = 1 / devolvido = 2 / atrasado = 3',
    `DateCreated` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `DateDevoluted` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`User_id`) REFERENCES Users(`id`),
    FOREIGN KEY (`Book_id`) REFERENCES Books(`id`)
)
ENGINE = InnoDB;