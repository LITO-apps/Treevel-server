ALTER TABLE records MODIFY stage_id VARCHAR(128) NOT NULL;

INSERT INTO players VALUES (3, 'test3', NULL, '2020-05-02 12:48:35', '2020-05-02 12:48:35');
INSERT INTO players VALUES (4, 'test4', NULL, '2020-05-02 12:48:35', '2020-05-02 12:48:35');
INSERT INTO players VALUES (5, 'test5', NULL, '2020-05-02 12:48:35', '2020-05-02 12:48:35');
INSERT INTO players VALUES (6, 'test6', NULL, '2020-05-02 12:48:35', '2020-05-02 12:48:35');

UPDATE records SET stage_id = 'Spring-1-1' WHERE stage_id = '1';
UPDATE records SET stage_id = 'Spring-1-2' WHERE stage_id = '2';

INSERT INTO records VALUES (9001, 4, 'Spring-1-1', TRUE, 0, 1, 200.11);
INSERT INTO records VALUES (9000, 3, 'Spring-1-1', TRUE, 0, 1, 100.02);
INSERT INTO records VALUES (9002, 5, 'Spring-1-1', TRUE, 0, 1, 150.25);
INSERT INTO records VALUES (9003, 6, 'Spring-1-1', TRUE, 0, 1, 300.45);
INSERT INTO records VALUES (9004, 3, 'Spring-1-2', TRUE, 0, 1, 50.02);
INSERT INTO records VALUES (9005, 4, 'Spring-1-2', TRUE, 0, 1, 133.11);
INSERT INTO records VALUES (9006, 5, 'Spring-1-2', TRUE, 0, 1, 122.25);
INSERT INTO records VALUES (9007, 6, 'Spring-1-2', TRUE, 0, 1, 331.45);
