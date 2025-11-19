package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*

		// Serve static files (CSS, images)
		fs := http.FileServer(http.Dir("static"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))

		// Serve HTML Pages
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "home.html")
		})

		http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "about.html")
		})

		http.HandleFunc("/programs", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "programs.html")
		})

		http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "contact.html")
		})

		log.Println("Server started on port 8080")
		http.ListenAndServe(":8080", nil)
	*/

	/*

		level := flag.String("level", "INFO", "log level to filter for ")
		flag.Parse()

		f, err := os.Open("./log.txt")
		if err != nil {
			log.Fatalf("failed to open file: %s", err)
		}
		defer f.Close()
		buffReader := bufio.NewReader(f)
		for line, err := buffReader.ReadString('\n'); err == nil; line, err = buffReader.ReadString('\n') {
			if strings.Contains(line, *level) {
				fmt.Print(line)
			}
		}
	*/

	fruits := []int{10, 20, 30, 40, 50}
	for i := range fruits {
		fmt.Println("this is no ", fruits[i])
		fmt.Println(reflect.TypeOf(fruits))
	}

}
