package main

import (
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/calendar"
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/handlers"
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/service"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

func main() {

	model := calendar.NewCalendar()
	serv := service.NewService(model)
	controller := handlers.NewController(serv)

	router := http.NewServeMux()
	router.HandleFunc("/create_event", controller.Create)
	router.HandleFunc("/update_event", controller.Update)
	router.HandleFunc("/delete_event", controller.Delete)
	router.HandleFunc("/events_for_day", controller.FilterByDay)
	router.HandleFunc("/events_for_week", controller.FilterByWeek)
	router.HandleFunc("/events_for_month", controller.FilterByMonth)

	logging := log.New(os.Stderr, "MyLogger:", log.LstdFlags)
	loggingMiddleware := LoggingMiddleware(logging)

	loggedRouter := loggingMiddleware(router)

	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logging.Println("init server")
	logging.Fatal(srv.ListenAndServe())
}

func LoggingMiddleware(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Printf("error:%v\n Traceback:\n %v\n", err, debug.Stack())

				}
			}()

			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Printf("%v request: %v method: %v\n", start, r.URL.EscapedPath(), r.Method)
		}
		return http.HandlerFunc(fn)
	}
}
