{{if .Site}}
UPDATE BLOCKS
    SET ORIGIN_SITE_NAME = ?,
        LAST_MODIFIED_BY = ?,
        LAST_MODIFICATION_DATE = ?
    WHERE BLOCK_NAME = ?
{{else}}
UPDATE BLOCKS
    SET OPEN_FOR_WRITING = ?,
        LAST_MODIFIED_BY = ?,
        LAST_MODIFICATION_DATE = ?
    WHERE BLOCK_NAME = ?
{{end}}
