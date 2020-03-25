use db;

CREATE TABLE IF NOT EXISTS player
(
    id                INT NOT NULL AUTO_INCREMENT,
    name              VARCHAR(64) NOT NULL COMMENT 'プレイヤーネーム',
    playing_stage_id  INT COMMENT 'プレイ中のステージID',
    create_time       TIMESTAMP NOT NULL COMMENT 'アカウント作成時刻',
    last_login_time   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '最終ログイン時刻',
    PRIMARY KEY (id)
);
