# プレイヤー情報
CREATE TABLE IF NOT EXISTS player
(
    id                INT NOT NULL AUTO_INCREMENT,
    name              VARCHAR(64) NOT NULL COMMENT 'プレイヤーネーム',
    playing_stage_id  INT COMMENT 'プレイ中のステージID',
    create_time       TIMESTAMP NOT NULL COMMENT 'アカウント作成時刻',
    last_login_time   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '最終ログイン時刻',

    PRIMARY KEY (id)
);

# 各ステージに対する，各プレイヤーの記録
CREATE TABLE IF NOT EXISTS record
(
    id                INT NOT NULL AUTO_INCREMENT,
    player_id         INT NOT NULL,
    stage_id          INT NOT NULL,
    is_clear          BOOLEAN DEFAULT FALSE NOT NULL COMMENT 'クリア回数',
    play_times        INT NOT NULL COMMENT 'プレイ回数',
    first_clear_times INT COMMENT '初回クリアまでのプレイ回数',
    min_clear_time    TIME COMMENT '最短クリア時間',

    PRIMARY KEY (id),
    FOREIGN KEY (player_id) REFERENCES player(id) ON DELETE CASCADE
);
