// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	r "gopkg.in/dancannon/gorethink.v2"
)

func Example() {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:32770",
	})
	if err != nil {
		log.Fatalln(err)
	}

	res, err := r.Expr("Hello World").Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	type Post struct {
		Title   string `gorethink:"title"`
		Content string `gorethink:"content"`
	}

	log.Debugln("started insert 5000 rows")
	for i := 0; i < 5000; i++ {
		_, _ = r.DB("achmed").Table("meins").Insert(Post{
			Title:   "Lorem ipsum",
			Content: "Dolor sit amet",
		}).RunWrite(session)

		// fmt.Println("%d row inserted", resp.Inserted)
	}
	log.Debugln("end insert 5000 rows")

	if err != nil {
		fmt.Print(err)
		return
	}

	res, err = r.DB("achmed").Table("meins").Count().Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	var hero string
	err = res.One(&hero)

	log.Debugln(hero)

	// var row Post
	// for res.Next(&row) {
	// 	log.Debugln(row)
	// 	// Do something with row
	// }

	// Output:
	// Hello World
}

func main() {
	log.SetLevel(log.DebugLevel)
	fmt.Println("Hello, 世界")
	Example()
}
