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
	Use:   "model",
	Short: "Generate a DB backed Store type for a provided struct",
	Long: `This command will generate more-or-less a DAO for a given struct. For Example given the following struct:

type Foo struct {
	ID int "db: id"
	Bar string "db: bar"
}

model will generate:

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
			g.Printf("import \"database/sql\"\n")

			//generate the struct we need
			g.Printf("type %s struct {}\n", storeNameFlag)
			generateFindByMethods(g, storeNameFlag, tableNameFlag, typeName, fields)
		}

		if outputFlag == "" {
			outputFlag = filepath.Join(args[0], strings.ToLower(fmt.Sprintf("%s_model.go", typeFlag)))
		}

		g.Run(path, typeFlag, outputFlag, propertizer)
	},
}

func generateFindByMethods(g OutputBuffer, daoObjectName string, tableName string, structTypeName string,
	fields []parse.Field) {

	for _, field := range fields {
		g.Printf("\n ")
		g.Printf("func (store *%s) FindBy%s(p %s) (*%s, error) {\n", daoObjectName,
			field.Name, field.TypeName, structTypeName)
		g.Printf(" var x %s\n", structTypeName)
		g.Printf(" sql := \"SELECT * FROM %s WHERE %s = ?\"\n", tableName, field.Tags["db"].Value)
		g.Printf(" err := db.QueryRowx(sql, p).StructScan(&x)\n")
		g.Printf(" return &x, err\n")
		g.Printf("}")
	}
}
