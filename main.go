// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"

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

	type Event struct {
		Action string `gorethink:"action"`
		Eid    int    `gorethink:"eid"`
	}

	type Post struct {
		Title   string `gorethink:"title"`
		Content string `gorethink:"content"`
		Num     int    `gorethink:"num"`
		Event   Event  `gorethink:"event"`
	}

	log.Debugln("started insert 5000 rows")
	rand.Seed(42) // Try changing this number!
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	events := []Event{
		{"click", 1},
		{"open", 2},
		{"bounce", 3},
	}

	for i := 0; i < 1; i++ {
		_, _ = r.DB("achmed").Table("meins").Insert(Post{
			Title:   "Lorem ipsum",
			Content: answers[rand.Intn(len(answers))],
			Num:     rand.Intn(10),
			Event:   events[rand.Intn(len(events))],
		}).RunWrite(session, r.RunOpts{Durability: "soft"})
		if i%1000 == 0 {
			log.Infoln("1000 written")
		}
		// fmt.Println("%d row inserted", resp.Inserted)
	}
	log.Debugln("end insert 5000 rows")

	if err != nil {
		fmt.Print(err)
		return
	}

	res, err = r.DB("achmed").Table("meins").Filter(map[string]interface{}{
		// "num":     9,
		// "content": "Reply hazy try again",
		"event": map[string]interface{}{"action": "click"},
	}).Count().Run(session)
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
