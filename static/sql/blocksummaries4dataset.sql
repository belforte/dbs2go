
SELECT (
    SELECT NVL(SUM(BS.BLOCK_SIZE), 0)
    FROM {{.Owner}}.BLOCKS BS
    JOIN {{.Owner}}.DATASETS DS ON BS.DATASET_ID=DS.DATASET_ID
    WHERE DS.dataset=:dataset
) AS FILE_SIZE,
(
    SELECT NVL(SUM(BS.FILE_COUNT),0)
    FROM {{.Owner}}.BLOCKS BS
    JOIN {{.Owner}}.DATASETS DS ON BS.DATASET_ID=DS.DATASET_ID
    WHERE DS.dataset=:dataset
) AS NUM_FILE,
(
    SELECT NVL(SUM(FS.EVENT_COUNT),0)
    FROM {{.Owner}}.FILES FS
    JOIN {{.Owner}}.BLOCKS BS ON BS.BLOCK_ID=FS.BLOCK_ID
    JOIN {{.Owner}}.DATASETS DS ON BS.DATASET_ID=DS.DATASET_ID
    WHERE DS.dataset=:dataset
) AS NUM_EVENT
FROM DUAL
