package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

func monthDayYear(t time.Time) string {
	return t.Format("January 2, 2006")
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var fm = template.FuncMap{
		"fdateMDY": monthDayYear,
	}
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
	if err := tpl.ExecuteTemplate(w, "index.gohtml", time.Now()); err != nil {
		log.Fatalln(err)
	}
}

func main() {

	type User struct {
		Name   string
		Coupon string
		Amount int64
		Ul     []string
	}
	type User2 struct {
		Name   string
		Coupon string
		Amount int64
		Ul     map[string]string
	}
	// user := User{
	// 	Name:   "Rick",
	// 	Coupon: "IAMAWESOMEGOPHER",
	// 	Amount: 5000,
	// 	Ul:     []string{"yaba", "daba", "duuu"},
	// }

	// user2 := User2{
	// 	Name:   "Rick",
	// 	Coupon: "IAMAWESOMEGOPHER",
	// 	Amount: 5000,
	// 	Ul:     map[string]string{"yaba": "fffbg rgb ", "daba": "vy rere43 rr", "duuu": "b gfb4t"},
	// }

	superhero := struct {
		Name  string
		Motto string
		Ul    map[string]string
		Date  time.Time
	}{
		Name:  "Bruise Wayne",
		Motto: "I am batman",
		Ul:    map[string]string{"yaba": "fffbg rgb ", "daba": "vy rere43 rr", "duuu": "b gfb4t"},
		Date:  time.Now(),
	}

	var fm = template.FuncMap{
		"fdateMDY": monthDayYear,
	}

	tpl := template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))

	err := tpl.Execute(os.Stdout, superhero)
	if err != nil {
		panic(err)
	}

}
