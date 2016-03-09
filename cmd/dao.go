// Copyright © 2016 NAME HERE <chris@uberbrodt.net>
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
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/serenize/snaker"
	"github.com/spf13/cobra"
	"github.com/uberbrodt/exemplar/data"
	"github.com/uberbrodt/exemplar/parse"
)

var finderNameFlag string
var storeNameFlag string
var templateFlag string
var tableNameFlag string
var tplLocFlag string

func init() {
	RootCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.
	modelCmd.Flags().StringVarP(&finderNameFlag, "finderName", "f", "", "name of the type where Find methods will be generated")
	modelCmd.Flags().StringVarP(&storeNameFlag, "storeName", "s", "", "name of the type where insert/update methods will be generated")
	modelCmd.Flags().StringVarP(&templateFlag, "tpl", "", "dao/sqlx", "Specify template to generate DAO's from")
	modelCmd.Flags().StringVarP(&tplLocFlag, "tplFolder", "", "./exemplar", "Specify location of custom template files. If unset, will look for builtin's or ./exemplar")
	modelCmd.Flags().StringVarP(&tableNameFlag, "tableName", "", "", "The name of the db table associated with struct")

}

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "dao",
	Short: "Generate a Data Access Object, encapsulating populating and persisting a struct to a DB",
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
		if len(path) == 0 {
			path = append(path, "./")
		}
		if typeFlag == "" {
			fmt.Println("MISSING REQUIRED ARGUMENT: --type needs to be set")
			os.Exit(-1)
		}
		if storeNameFlag == "" {
			storeNameFlag = fmt.Sprintf("%sStore", typeFlag)
		}
		if finderNameFlag == "" {
			finderNameFlag = fmt.Sprintf("%sFinder", typeFlag)
		}

		if templateFlag == "" {
			templateFlag = "dao/sqlx"
		}

		templateFlag = fmt.Sprintf("templates/%s.tmpl", templateFlag)

		action := func(typeName string, fields []parse.Field, imports []parse.Import) {

			funcMap := template.FuncMap{
				"funcCase":     funcCase,
				"updateSQL":    updateSQLXFunc,
				"insertSQL":    insertSQLXFunc,
				"insertSQLPgx": insertSQLPgx,
				"updateSQLPgx": updateSQLPgx,
				"nameOrTag":    nameOrTag,
			}

			templateData, templateErr := data.Asset(templateFlag)
			if templateErr != nil {
				panic(fmt.Sprintf("dao template not loaded! bin-data err? %s", templateErr))
			}
			tmplString := string(templateData)
			//DEBUG: fmt.Print(tmplString)
			tmpl := template.Must(template.New("dao_tmpl").Funcs(funcMap).Parse(tmplString))

			filtered := make([]parse.Field, 0)
			for _, field := range fields {
				if field.Tags["exclude_dao"].Value != "true" && field.Name != "NeedsInsert" {
					filtered = append(filtered, field)
				}
			}
			tmpl.ExecuteTemplate(&g.Buf, "dao_tmpl",
				struct {
					Imports        []parse.Import
					Fields         []parse.Field
					StructTypeName string
					TableName      string
					FinderName     string
					StoreName      string
				}{
					Imports:        imports,
					Fields:         filtered,
					StructTypeName: typeName,
					TableName:      tableNameFlag,
					FinderName:     finderNameFlag,
					StoreName:      storeNameFlag})
		}

		if outputFlag == "" {
			outputFlag = filepath.Join(strings.Replace(path[0], ".go", "", -1), strings.ToLower(fmt.Sprintf("%s_dao.go", snaker.CamelToSnake(typeFlag))))
		}

		g.Run(path, typeFlag, outputFlag, action)
	},
}

func funcCase(f string) string {
	return strings.Replace(strings.Title(CamelCase(f)), "Id", "ID", -1)
}

func insertSQLXFunc(fields []parse.Field, tableName string) string {
	insertSQL := fmt.Sprintf("INSERT INTO %s (", tableName)
	insertSQLVals := "VALUES ("
	for i := 0; i < len(fields); i++ {
		dbTag := nameOrTag(fields[i])
		if i == (len(fields) - 1) {
			insertSQL += fmt.Sprintf("%s", dbTag)
			insertSQLVals += fmt.Sprintf(":%s", dbTag)
		} else {
			insertSQL += fmt.Sprintf("%s, ", dbTag)
			insertSQLVals += fmt.Sprintf(":%s, ", dbTag)
		}
	}

	return fmt.Sprintf("%s) %s)", insertSQL, insertSQLVals)
}

func updateSQLXFunc(fields []parse.Field, tableName string) string {
	updateSQL := fmt.Sprintf("UPDATE %s SET ", tableName)
	updateWhereSQL := fmt.Sprintf(" WHERE id = :id")
	for i := 0; i < len(fields); i++ {
		dbTag := nameOrTag(fields[i])

		if i == len(fields)-1 {
			updateSQL += fmt.Sprintf("%s = :%s", dbTag, dbTag)
		} else {
			updateSQL += fmt.Sprintf("%s = :%s, ", dbTag, dbTag)
		}
	}
	return updateSQL + updateWhereSQL
}

func insertSQLPgx(fields []parse.Field, tableName string) string {
	insertSQL := fmt.Sprintf("INSERT INTO %s (", tableName)
	insertSQLVals := "VALUES ("
	for i := 0; i < len(fields); i++ {
		dbTag := nameOrTag(fields[i])
		if i == (len(fields) - 1) {
			insertSQL += fmt.Sprintf("%s", dbTag)
			insertSQLVals += fmt.Sprintf("$%d", i+1)
		} else {
			insertSQL += fmt.Sprintf("%s, ", dbTag)
			insertSQLVals += fmt.Sprintf("$%d, ", i+1)
		}
	}

	return fmt.Sprintf("%s) %s)", insertSQL, insertSQLVals)

}

func updateSQLPgx(fields []parse.Field, tableName string) string {

	updateSQL := fmt.Sprintf("UPDATE %s SET ", tableName)
	var updateWhereSQL string
	for i := 0; i < len(fields); i++ {
		dbTag := nameOrTag(fields[i])
		if dbTag == "id" {
			updateWhereSQL = fmt.Sprintf(" WHERE id = $%d", i+1)
			continue
		}

		if i == len(fields)-1 {
			updateSQL += fmt.Sprintf("%s = $%d", dbTag, i+1)
		} else {
			updateSQL += fmt.Sprintf("%s = $%d, ", dbTag, i+1)
		}
	}
	return updateSQL + updateWhereSQL
}

func nameOrTag(field parse.Field) string {
	if strings.TrimSpace(field.Tags["db"].Value) != "" {
		return strings.TrimSpace(field.Tags["db"].Value)
	} else {
		return snaker.CamelToSnake(field.Name)
	}
}
