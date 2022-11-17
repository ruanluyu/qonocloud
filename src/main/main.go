package main

import "milai/qonocloud/server"

//"fmt"
//"net/http"

var defaultPort int = 8000

/*func handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, fmt.Sprintf("MiLai Cloud Server \nMethod: %s\nHost: %s\nURL: %s",
		req.Method,
		req.Host,
		req.URL))
	fmt.Printf("Request recieved:\n%s\n%s\n%s\n", req.Method, req.Host, req.URL)
}*/

func main() {
	server.Hello1()
	/*fmt.Println("Start...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), nil)*/
}
