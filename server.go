package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})

	server := &http.Server{
		Addr:    ":3030",
		Handler: mux,
	}

	// 创建系统信号接收器
	// done := make(chan os.Signal)
	// signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// go func() {
	// 	<-done

	// 	if err := server.Shutdown(context.Background()); err != nil {
	// 		log.Fatal("Shutdown server:", err)
	// 	}
	// }()

	log.Println("Starting HTTP server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}

type helloHandler struct{}

func (*helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ctmd")
}

/*background start*/
//  body{
//     height: 100%;
//     opacity: 0.968;
//     background-position: right;
//     background-size: contain;
//     background-repeat: no-repeat;
//     background-image:url('file:///Users/lizhaolun/Downloads/HupuBBS_201115110105-1242822577.jpg');
//    }
/*background end*/
