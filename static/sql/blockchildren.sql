SELECT BC.BLOCK_NAME FROM {{.Owner}}.BLOCKS BC
                        JOIN {{.Owner}}.BLOCK_PARENTS BPRTS
                            ON BPRTS.THIS_BLOCK_ID = BC.BLOCK_ID
                        JOIN {{.Owner}}.BLOCKS BP
                            ON BPRTS.PARENT_BLOCK_ID = BP.BLOCK_ID
