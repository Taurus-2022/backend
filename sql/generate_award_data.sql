DELIMITER $$
DROP PROCEDURE IF EXISTS insert_current_data_uuid $$
CREATE PROCEDURE insert_award_data(IN item INTEGER)
BEGIN
    DECLARE counter INT;
    SET counter = item;
    START TRANSACTION ;
    WHILE counter >= 1
        DO
            INSERT INTO award (code, type, is_used)
            VALUES (replace(uuid(), '-', ''), floor(1 + rand() * 3), 0);
            SET counter = counter - 1;
        END WHILE;
    COMMIT ;
END
$$

CALL insert_award_data(1000);