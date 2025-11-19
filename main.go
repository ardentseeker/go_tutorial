package main

import (
	"fmt"
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

	// this is a method of printing datat type
	/*
		fruits := []int{10, 20, 30, 40, 50}
		for i := range fruits {
			fmt.Println("this is no ", fruits[i])
			fmt.Println(reflect.TypeOf(fruits))
		}
	*/
	/*
		// use of pointer
		a := 42
		b := &a
		fmt.Println("value of a:", a)
		fmt.Println("address of b:", *b)
		*b = 21
		fmt.Println("new value of b:", *b)
	*/

	// use of Array
	var arr [5]int = [5]int{10, 20, 30, 40, 50}
	var arr2 [5]int = arr

	fmt.Println("Array:", arr2)
	fmt.Println("Length of Array:", arr)
	arr2[3] = 100
	fmt.Println("Modified Array:", arr2)

}
