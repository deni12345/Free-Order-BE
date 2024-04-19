CREATE TABLE IF NOT EXISTS `User` (
    Id BIGINT NOT NULL AUTO_INCREMENT,
    UserName VARCHAR(255),
    Password VARCHAR(20),
    PRIMARY KEY (Id)
);

CREATE TABLE IF NOT EXISTS `Order` (
    Id BIGINT NOT NULL AUTO_INCREMENT,
    FoodName VARCHAR(255),
    Amount INT,
    CreateAt VARCHAR(25),
    PRIMARY KEY (Id)
);

CREATE TABLE IF NOT EXISTS `OrderSheet` (
    Id BIGINT NOT NULL AUTO_INCREMENT,
    SheetName VARCHAR(255),
    CreateAt VARCHAR(25),
    UserId BIGINT NOT NULL,
    OrderId BIGINT NOT NULL,
    PRIMARY KEY (Id),
    FOREIGN KEY (UserId) REFERENCES User(Id),
    FOREIGN KEY (OrderId) REFERENCES `Order`(Id)
);