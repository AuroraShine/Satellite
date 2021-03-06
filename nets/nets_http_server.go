package nets

import (
	"fmt"
	"log"
	"net/http"
	. "satellite/global"
	"time"
)

func StartHttpServer(ip string, port string) {
	server := http.Server{
		Addr:         ip + ":" + port,
		WriteTimeout: HTTPWriteTimeout * time.Millisecond,
		ReadTimeout:  HTTPReadTimeout * time.Millisecond,
		Handler:      createHttpRouter(),
	}
	fmt.Println("Start Listen And Server on ", ip+":"+port)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error Listen And Server:", err)
	}
}
