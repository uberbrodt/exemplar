// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/serenize/snaker"
	"github.com/spf13/cobra"
	"github.com/uberbrodt/exemplar/parse"
)

var storeNameFlag string
var dbtypeFlag string
var tableNameFlag string

func init() {
	RootCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.
	modelCmd.Flags().StringVarP(&storeNameFlag, "storename", "s", "", "name of the type that will be the store")
	modelCmd.Flags().StringVarP(&dbtypeFlag, "dbtype", "d", "pg", "Currently supports only pg, though shouldn't matter much b/c sqlx ")
	modelCmd.Flags().StringVarP(&tableNameFlag, "tableName", "", "", "The name of the db table associated with struct")

}

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "dao",
	Short: "Generate a DB backed Store type for a provided struct",
	Long: `This command will generate more-or-less a DAO for a given struct. For Example given the following struct:

type Foo struct {
	ID int "db: id"
	Bar string "db: bar"
}

dao will generate:

type FooStorePg struct {
}

func (store *FooStorePg) GetByID(id int) Foo {
	//sql stuff to get a Foo from Postgres
}
....`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		g := new(parse.Generator)
		path := cmd.Flags().Args()

		propertizer := func(typeName string, fields []parse.Field, imports []parse.Import) {

			g.Printf("import \"github.com/jmoiron/sqlx\"\n")
			g.Printf("import \"github.com/uberbrodt/exemplar/model\"\n")
			g.Printf("import \"fmt\"\n") // Used by all methods.
			//g.Printf("import \"database/sql\"\n")

			for _, imprt := range imports {
				g.Printf("import \"%s\" \n", imprt.ImportedName)
			}

			//generate the struct we need
			g.Printf("type %s struct { DB  model.ExemplarSqlx }\n", storeNameFlag)
			generateFindByMethods(g, storeNameFlag, tableNameFlag, typeName, fields)
			generateFindByTxMethods(g, storeNameFlag, tableNameFlag, typeName, fields)

			generateFindSliceByMethods(g, storeNameFlag, tableNameFlag, typeName, fields)
			generateFindSliceByTxMethods(g, storeNameFlag, tableNameFlag, typeName, fields)
			generateSaveMethods(g, storeNameFlag, tableNameFlag, typeName, fields)
		}

		if outputFlag == "" {
			outputFlag = filepath.Join(args[0], strings.ToLower(fmt.Sprintf("%s_dao.go", snaker.CamelToSnake(storeNameFlag))))
		}

		g.Run(path, typeFlag, outputFlag, propertizer)
	},
}

func generateFindByMethods(g OutputBuffer, daoObjectName string, tableName string, structTypeName string,
	fields []parse.Field) {

	for _, field := range fields {
		if field.Name == "NeedsInsert" {
			continue
		}
		g.Printf("\n ")
		g.Printf("func (store *%s) FindBy%s(p %s) (*%s, error) {\n", daoObjectName,
			funcCase(field.Name), field.TypeName, structTypeName)
		g.Printf(" var x %s\n", structTypeName)
		g.Printf(" sql := \"SELECT * FROM %s WHERE %s = ? LIMIT 1\"\n", tableName, field.Tags["db"].Value)
		g.Printf(" err := store.DB.QueryRowx(sql, p).StructScan(&x)\n")
		g.Printf(" return &x, err\n")
		g.Printf("}\n")
	}
}

func generateFindByTxMethods(g OutputBuffer, daoObjectName string, tableName string, structTypeName string,
	fields []parse.Field) {

	for _, field := range fields {
		if field.Name == "NeedsInsert" {
			continue
		}
		g.Printf("\n ")
		g.Printf("func (store *%s) FindBy%sTx(p %s, tx *sqlx.Tx) (*%s, error) {\n", daoObjectName,
			funcCase(field.Name), field.TypeName, structTypeName)
		g.Printf(" var x %s\n", structTypeName)
		g.Printf(" sql := \"SELECT * FROM %s WHERE %s = ? LIMIT 1\"\n", tableName, field.Tags["db"].Value)
		g.Printf(" err := tx.QueryRowx(sql, p).StructScan(&x)\n")
		g.Printf(" return &x, err\n")
		g.Printf("}\n")
	}
}

func generateFindSliceByMethods(g OutputBuffer, daoObjectName string, tableName string,
	structTypeName string, fields []parse.Field) {

	for _, field := range fields {
		if field.Name == "NeedsInsert" {
			continue
		}
		g.Printf("\n")
		g.Printf("func (store *%s) FindSliceBy%s(p %s, orderBy interface{}, limit int, offset int) ([]*%s, error) {\n",
			daoObjectName, funcCase(field.Name), field.TypeName, structTypeName)
		g.Printf("x := make([]*%s, 0)\n", structTypeName)
		g.Printf("sql := \"SELECT * FROM %s WHERE %s = ? \"\n", tableName, field.Tags["db"].Value)
		g.Printf(" if orderBy != \"\" { \n")
		g.Print(`  sql += fmt.Sprintf("ORDER BY %s ", orderBy)`)
		g.Printf("\n}\n")
		g.Printf("if limit > 0 { \n")
		g.Print(` sql += fmt.Sprintf("LIMIT %d", limit)` + "\n")
		g.Printf("}\n")
		g.Printf("if offset > 0 {\n")
		g.Print(` sql += fmt.Sprintf("OFFSET %d", offset)` + "\n")
		g.Printf("}\n")
		g.Printf("rows, err := store.DB.Queryx(sql, p)\n")
		g.Printf("if err != nil { return x, err }\n")
		g.Printf("defer rows.Close()\n")
		g.Printf("for rows.Next() {\n")
		g.Printf("i := new(%s)\n", structTypeName)
		g.Printf("scanErr := rows.StructScan(i)\n")
		g.Printf("if scanErr != nil { return x,scanErr }\n ")
		g.Printf("x = append(x, i)")
		g.Printf("}\n")
		g.Printf("return x,nil")
		g.Printf("}\n")

		//ORDER BY %s LIMIT OFFSET ", )

	}

}
func generateFindSliceByTxMethods(g OutputBuffer, daoObjectName string, tableName string,
	structTypeName string, fields []parse.Field) {

	for _, field := range fields {
		if field.Name == "NeedsInsert" {
			continue
		}
		g.Printf("\n")
		g.Printf("func (store *%s) FindSliceBy%sTx(p %s, tx *sqlx.Tx, orderBy interface{}, limit int, offset int) ([]*%s, error) {\n",
			daoObjectName, funcCase(field.Name), field.TypeName, structTypeName)
		g.Printf("x := make([]*%s, 0)\n", structTypeName)
		g.Printf("sql := \"SELECT * FROM %s WHERE %s = ? \"\n", tableName, field.Tags["db"].Value)
		g.Printf(" if orderBy != \"\" { \n")
		g.Print(`  sql += fmt.Sprintf("ORDER BY %s ", orderBy)`)
		g.Printf("\n}\n")
		g.Printf("if limit > 0 { \n")
		g.Print(` sql += fmt.Sprintf("LIMIT %d", limit)` + "\n")
		g.Printf("}\n")
		g.Printf("if offset > 0 {\n")
		g.Print(` sql += fmt.Sprintf("OFFSET %d", offset)` + "\n")
		g.Printf("}\n")
		g.Printf("rows, err := tx.Queryx(sql, p)\n")
		g.Printf("if err != nil { return x, err }\n")
		g.Printf("defer rows.Close()\n")
		g.Printf("for rows.Next() {\n")
		g.Printf("i := new(%s)\n", structTypeName)
		g.Printf("scanErr := rows.StructScan(i)\n")
		g.Printf("if scanErr != nil { return x,scanErr }\n ")
		g.Printf("x = append(x, i)")
		g.Printf("}\n")
		g.Printf("return x,nil")
		g.Printf("}\n")

		//ORDER BY %s LIMIT OFFSET ", )

	}
}

func generateSaveMethods(g OutputBuffer, daoObjectName string, tableName string,
	structTypeName string, fields []parse.Field) {

	g.Printf("func (store *%s) Save(o *%s) error {\n", daoObjectName, structTypeName)
	g.Printf("if o.NeedsInsert {\n")
	g.Printf(" return store.insert(o) \n")
	g.Printf("} else { \n")
	g.Printf("return store.update(o) }\n")
	g.Printf("}\n")

	g.Printf("func (store *%s) SaveTx(o *%s, tx *sqlx.Tx) error {\n", daoObjectName, structTypeName)
	g.Printf("if o.NeedsInsert {\n")
	g.Printf(" return store.insertTx(o, tx) \n")
	g.Printf("} else { \n")
	g.Printf("return store.updateTx(o, tx) }\n")
	g.Printf("}\n")

	insertSQL := fmt.Sprintf("INSERT INTO %s (", tableName)
	insertSQLVals := "VALUES("
	updateSQL := fmt.Sprintf("UPDATE %s SET", tableName)
	updateWhereSQL := fmt.Sprintf("WHERE id = :id")

	for _, f := range fields {
		if f.Name == "NeedsInsert" {
			continue
		}
		dbTag := f.Tags["db"].Value
		insertSQL += fmt.Sprintf("%s,", dbTag)
		insertSQLVals += fmt.Sprintf(":%s,", dbTag)
		updateSQL += fmt.Sprintf("%s = :%s,", dbTag, dbTag)
	}

	insertSQL = fmt.Sprintf("%s) %s)", insertSQL, insertSQLVals)
	updateSQL += updateWhereSQL
	insertMethods(g, insertSQL, daoObjectName, structTypeName)
	updateMethods(g, updateSQL, daoObjectName, structTypeName)
}

func insertMethods(g OutputBuffer, sql, daoObjectName, structTypeName string) {
	g.Printf("func (store *%s) insert(o *%s)  error {\n",
		daoObjectName, structTypeName)
	g.Printf("_, err := store.DB.NamedExec(\"%s\", o)\n", sql)
	g.Printf("if err == nil {\n")
	g.Printf("o.NeedsInsert = false\n")
	g.Printf("}\n")
	g.Printf("return err\n")
	g.Printf("}\n")

	g.Printf("func (store *%s) insertTx(o *%s, tx *sqlx.Tx)  error {\n",
		daoObjectName, structTypeName)
	g.Printf("_, err := tx.NamedExec(\"%s\", o)\n", sql)
	g.Printf("if err == nil {\n")
	g.Printf("o.NeedsInsert = false\n")
	g.Printf("}\n")
	g.Printf("return err\n")
	g.Printf("}\n")
}

func updateMethods(g OutputBuffer, sql, daoObjectName, structTypeName string) {
	g.Printf("func (store *%s) update(o *%s)  error {\n",
		daoObjectName, structTypeName)
	g.Printf("_, err := store.DB.NamedExec(\"%s\", o)\n", sql)
	g.Printf("return err\n")
	g.Printf("}\n")

	g.Printf("func (store *%s) updateTx(o *%s, tx *sqlx.Tx)  error {\n",
		daoObjectName, structTypeName)
	g.Printf("_, err := tx.NamedExec(\"%s\", o)\n", sql)
	g.Printf("return err\n")
	g.Printf("}\n")
}

func funcCase(f string) string {
	return strings.Replace(strings.Title(CamelCase(f)), "Id", "ID", -1)
}
