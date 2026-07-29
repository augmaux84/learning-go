package main

import "fmt"
import "net/http"
// import "log"

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// if 
		if r.FormValue("q") != "" {
            fmt.Fprintf(w, `
                <body>
                    <h1>Hello, %s</h1>
                </body>
            `, r.FormValue("q"))
            return
        }
		
		fmt.Fprintf(w, `
			<body>
				<form action="/" method="GET">
					<label>Enter your name:</label> <br>
					<input name="q"> <br>
					<button type="submit">Submit</button>
				</form>
			</body>
		`)

	})	

    http.ListenAndServe(":3000", nil)
	// log.Fatal(http.ListenAndServe(":3000", nil))
}