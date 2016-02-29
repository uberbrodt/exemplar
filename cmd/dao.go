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
	"text/template"

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
		if storeNameFlag == "" {
			storeNameFlag = fmt.Sprintf("%sStore", typeFlag)
		}

		propertizer := func(typeName string, fields []parse.Field, imports []parse.Import) {

			funcMap := template.FuncMap{
				"funcCase": funcCase,
			}

			tmpl := template.Must(template.New("generic_tmpl").Funcs(funcMap).ParseFiles("templates/dao/generic.tmpl"))

			for idx, field := range fields {
				if field.Tags["exclude_dao"].Value == "true" || field.Name == "NeedsInsert" {
					copy(fields[idx:], fields[idx+1:])
					fields = fields[:len(fields)-1]
				}
			}

			tmpl.ExecuteTemplate(&g.Buf, "generic.tmpl",
				struct {
					Imports        []parse.Import
					StoreNameFlag  string
					Fields         []parse.Field
					StructTypeName string
					TableName      string
					DAOName        string
				}{
					Imports:        imports,
					Fields:         fields,
					StructTypeName: typeName,
					TableName:      tableNameFlag,
					DAOName:        storeNameFlag})

			generateSaveMethods(g, storeNameFlag, tableNameFlag, typeName, fields)

		}

		if outputFlag == "" {
			outputFlag = filepath.Join(strings.Replace(args[0], ".go", "", -1), strings.ToLower(fmt.Sprintf("%s_dao.go", snaker.CamelToSnake(storeNameFlag))))
		}

		g.Run(path, typeFlag, outputFlag, propertizer)
	},
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

	insertSQL := fmt.Sprintf("INSERT INTO %s ( ", tableName)
	insertSQLVals := "VALUES("
	updateSQL := fmt.Sprintf("UPDATE %s SET ", tableName)
	updateWhereSQL := fmt.Sprintf(" WHERE id = :id")

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
