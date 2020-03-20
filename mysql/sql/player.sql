use db;

DROP TABLE IF EXISTS player;

CREATE TABLE player
(
    id                INT NOT NULL AUTO_INCREMENT,
    name              VARCHAR(64) NOT NULL,
    playing_stage_id  INT,
    create_time       TIMESTAMP ,
    last_login_time   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
