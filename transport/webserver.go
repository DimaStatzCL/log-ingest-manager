package transport
import "net/http"

var outChannel chan <- string

func CreateWebTransport(out chan <- string){
	outChannel = out
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/event", handleEvent)
	http.ListenAndServe(":9000", nil)
}

func handleEvent(response http.ResponseWriter, request *http.Request) {
	query := request.URL.RawQuery
	outChannel <- query
	response.WriteHeader(200)
}

func handleHealth(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Log-manager is running"))
	response.WriteHeader(200)
}
