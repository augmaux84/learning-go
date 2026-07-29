package main

import "fmt"
import "net/http"
import "strconv"

func main() {
	
	// initial tests with the GO language

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	if r.Method == "POST" {
    //         fmt.Fprintf(w, `
    //             <body>
    //                 <h1>Hello, %s</h1>
    //             </body>
    //         `, r.FormValue("q"))
    //         return
    //     }
		
	// 	fmt.Fprintf(w, `
    //         <body>
    //             <form action="/" method="POST">
    //                 <label>Counter</label>
    //                 <input name="counter" value="1">
    //                 <button type="submit">Add</button>
    //             </form>
    //         </body>
    //         <a href="/">Reset</a>
    //     `)
    //     })
    
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
		    count, _ := strconv.Atoi(r.FormValue("counter"))
            count++
            fmt.Fprintf(w, `
            <body>
                <form action="/" method="POST">
                    <label>Counter</label>
                    <input name="counter" value="%d" readonly>
                    <button type="submit">Add</button>
                </form>
                <a href="/">Reset</a>
            </body>
            `, count)
            return
		}

        fmt.Fprintf(w, `
            <body>
                <form action="/" method="POST">
                    <label>Counter</label>
                    <input name="counter" value="1" readonly>
                    <button type="submit">Add</button>
                </form>

                <a href="/">Reset</a>
            </body> 
        `)
	})

	http.ListenAndServe(":3000", nil)
}