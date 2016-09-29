package receiver

import "math/rand"

type Event struct {
	Action string `gorethink:"action"`
	Eid    int    `gorethink:"eid"`
}

type Group struct {
	Id          int `gorethink:"id"`
	Activated   int `gorethink:"activated"`
	Deactivated int `gorethink:"deactivated"`
	Registered  int `gorethink:"registered"`
}

type Receiver struct {
	Email    string  `gorethink:"email"`
	Bounced  int     `gorethink:"bounced"`
	Imported int     `gorethink:"imported"`
	Source   string  `gorethink:"source"`
	Groups   []Group `gorethink:"groups"`
	Events   []Event `gorethink:"event"`
}

func init() {
	rand.Seed(42) // Try changing this number!
}

func Randomize() Receiver {
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

	groups := []Group{
		{1, 1212, 0, 123},
		{1, 1212, 123, 0},
		{1, 1212, 131, 123},
		{1, 112212, 0, 0},
		{2, 1212, 0, 123},
		{2, 1212, 123, 0},
		{2, 1212, 131, 123},
		{2, 112212, 0, 0},
		{3, 1212, 0, 123},
		{3, 1212, 123, 0},
		{3, 1212, 131, 123},
		{3, 112212, 0, 0},
	}
	events := []Event{
		{"click", rand.Intn(10)},
		{"open", rand.Intn(10)},
		{"bounce", rand.Intn(10)},
	}

	ret := Receiver{
		Email:    "bla@fasel.de",
		Source:   answers[rand.Intn(len(answers))],
		Bounced:  rand.Intn(10),
		Imported: rand.Intn(10),
		Events:   []Event{events[rand.Intn(len(events))], events[rand.Intn(len(events))]},
		Groups:   []Group{groups[rand.Intn(len(groups))]},
	}

	return ret
}
