CREATE TABLE cakes(
    id int AUTO_INCREMENT PRIMARY KEY,
    title varchar(255) NOT NULL,
    description varchar(255),
    rating float NOT NULL DEFAULT 0,
    image varchar(255),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME
);

