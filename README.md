#Propertizer: A generator mutators and accessors for Go struct types.

Basically a rip off of [Stringer](https://godoc.org/github.com/Go-zh/tools/cmd/stringer), but generates you Getters and Setters for structs.

###Hold on, why do I need this? This isn't Java?!
Good point! But like Java, Go is statically typed. If you want to write tests, you'll probably need interfaces. You can't specify fields as part of an interface, so that means you'll need methods.

###How to use

`go get install github.com/uberbrodt/propertizer`

Pick a type and point it towards the relevant filed

`propertizer -type YourStruct /github.com/your/package/yourstruct.go`

And it'll generate you a file called `yourstruct_propertizer.go`