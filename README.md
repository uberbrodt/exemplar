[![Go Report Card](https://goreportcard.com/badge/github.com/uberbrodt/exemplar)](https://goreportcard.com/report/github.com/uberbrodt/exemplar)

#Exemplar: A Go Code Generator

Exemplar is an extensible code generator for generating exemplary Go code! It's based off the [Stringer](https://godoc.org/github.com/Go-zh/tools/cmd/stringer) tool that was used to [introduce](https://blog.golang.org/generate) Go generators. It parses the target package, generates code based off provided struct(s) and then verfies that the package still works. It even runs `gofmt` afterwards so that your generated code is pretty!

It currently consists of two commands:

* `dao` will generate DAO like code for a struct using [sqlx](https://github.com/jmoiron/sqlx) or [pgx](https://github.com/jackc/pgx). Support for  other drivers is unplanned at the moment, though should be easy to accomplish on your own by providing custom templates
* `propertizer` will generate getters and setters for a struct. You are able to use struct tags to exclude method generation for fields.

The following is under consideration

*  `resource` will generate a set of REST methods and muxer routes (GET /foo/:id, PUT /foo/:id, etc) for a struct. The aim will be to make it solely standard library for performance & compatibility reasons, but might support either [negroni](https://github.com/codegangsta/negroni) or [gin](https://github.com/gin-gonic/gin) out of the box

##Install
`go get install github.com/uberbrodt/exemplar` will fetch the latest source and install the `exemplar` binary into your GOPATH.

##Commands
###DAO: Generate query and persistance code for SQL databases

As a statically typed language, lacking builtin generics, writing database access code can seem tedious when compared to dynamic languages like Ruby or Python. The `dao` command solves this by generating common finder and CRUD methods for any given struct. It will automatically convert field names to snake_case (so `ID` becomese `id`, `FirstName` becomes `first_name` etc.), though a name can be set via a `db` structtag.

In order for `dao` to generate code correctly, the struct needs to meet the following requirements:

* Must have a field either tagged or inferred as `id`. The field can be any simple type supported by the target DB.
* Any persisted fields need to be of a simple type. i.e., `dao` doesn't know how to save a field of `[]string`. Any fields that should not or cannot be persisted should be marked with the following struct tag: `exclude_dao:"true"`

####How To Use
The `dao` command can be run from the commandline as a `go:generate` comment. Here's a basic example which indicates the struct to generate code for (`-type=Basicstruct`), output, and the name of the database table associated with the struct. Run `exemplar dao --help` for all the options.

`exemplar dao --type=Basicstruct -o ./basicstruct_dao.go --tableName="basic_struct" --tpl="dao/pgx" ./basicstruct.go`

The `--tpl` option is important. It indicates the file template to generate the code. Right now there are two available:

* `dao/pgx` for accessing PostgreSQL databases
* `dao/sqlx` for everything else using `jmoiron/sqlx`

In the future, you will be able to pass in any template and have the following data injected when the command is run:

					Imports        []parse.Import
					Fields         []parse.Field
					StructTypeName string
					TableName      string
					FinderName     string
					MapperName     string
####What's generated?
The default templates will generate two new structs, a Finder and a Mapper (prefixed with it's type). The Finder has methods for each field: one that returns a single struct pointer, and another that returns a slice of struct pointers. The slice methods also accept offset and limit params so you can implement basic paging. For each Finder method, there is a companion method that accepts a Transaction object, so that higher level code can implement a Unit of Work pattern or just do manual transaction manipulation.


The Mapper will have methods for Update, Insert, and Delete.

Why two objects? Well, there may be cases where you want the Find methods in a different object than the CRUD methods. For example, you may not care, and just embed both into a business domain model, as in an ActiveRecord pattern.

- - -

###Propertizer: Generate mutators and accessors for Go struct types.

Generates you Getters and Setters for structs. Can ignore fields and generate interfaces for the getters and setters as well

###Hold on, why do I need this? This isn't Java?!
Well, you generally don't. For some reason I thought I needed this, but upon further review I didn't :(

Oh well, at least I learned a lot and now I'll put that knowledge to use to create something useful.

###How to use

`go get install github.com/uberbrodt/exemplar`

Pick a type and point it towards the relevant file

`exemplar propertizer -type YourStruct /github.com/your/package/yourstruct.go`

And it'll generate you a file called `yourstruct_propertizer.go`


####examples

Here's an example struct:

	//go:generate propertizer -type Basicstruct
	type Basicstruct struct {
		id           int
		date         time.Time
		status       string
		hiddenField  string
		privateField string `propertizer:"private"`
	}

running `go generate` inside of the package that contains this struct will automatically generate `basicstruct_properties.go` which will look like this:

	// Code generated by "propertizer-type=Basicstruct testdata/basicstruct.go"; DO NOT EDIT

	package main

	import "time"

	func (this *Basicstruct) ID() int {
		return this.id
	}
	func (this *Basicstruct) SetID(x int) {
		this.id = x
	}
	func (this *Basicstruct) Date() time.Time {
		return this.date
	}
	func (this *Basicstruct) SetDate(x time.Time) {
		this.date = x
	}
	func (this *Basicstruct) Status() string {
		return this.status
	}
	func (this *Basicstruct) SetStatus(x string) {
		this.status = x
	}
	func (this *Basicstruct) HiddenField() string {
		return this.hiddenField
	}
	func (this *Basicstruct) SetHiddenField(x string) {
		this.hiddenField = x
	}


####Private fields

You can mark a field as private (ie. Don't generate getters or setters) with the `propertizer:"private"` struct tag.

####Getter Prefix
Adding the flag `--getterPrefix` will preface the accessors with "Get". This is necessary where you have a struct that is also being used for JSON marshaling, and you need your field names to remain public. Go will throw an error when a method has the same name as a field ie. obj.ID() and obj.ID

####TODO
* Handle _ and aliased imports
* Auto generate tests for structs? We already compile the package after running, so really only useful if you're worried about code coverage metrics

