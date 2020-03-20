use db;

DROP TABLE IF EXISTS record;

CREATE TABLE record
(
    id                INT NOT NULL AUTO_INCREMENT,
    player_id         INT NOT NULL,
    stage_id          INT NOT NULL,
    is_clear          BIT(1),
    play_times        INT,
    first_clear_times INT,
    min_clear_time    TIME,
    PRIMARY KEY (id)
);
