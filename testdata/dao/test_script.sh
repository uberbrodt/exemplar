#1/bin/bash

#pgx
../../exemplar dao --type=Basicstruct -o ./basicstruct_dao.go --tableName="basic_struct" --tpl="dao/pgx" ./basicstruct.go

#sqlx
../../exemplar dao --type=Basicstruct -o ./basicstruct_dao.go --tableName="basic_struct" --tpl="dao/sqlx" ./basicstruct.go
