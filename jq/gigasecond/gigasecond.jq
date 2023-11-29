if .moment | length | . == 10 then "\(.moment)T00:00:00Z" else "\(.moment)Z" end | fromdateiso8601 | . + 1000000000 | todateiso8601 | .[0:-1]
