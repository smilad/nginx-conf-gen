package app

import (
	"context"
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
	"nginx/adapter/store"
	srv "nginx/app/server"
	"nginx/config"
	"nginx/controller/admin/http"
	trace "nginx/pkg/tracer"
	"nginx/repository"
	"nginx/service/admin"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// StartApplication func
func Start() {
	fmt.Printf("\n\n---------------  %s ----------------- \n\n", aurora.BgGreen("nginx"))
	// if go code crashed we get error and line
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// initial application configuration as Global variable
	config.InitConfig()

	database := store.NewGorm()

	closer, err := trace.InitTracer(config.C())
	if err != nil {
		log.Fatalf("error in tracer initialization")
	}
	defer closer.Close()

	//	 init layers
	dr := repository.NewDomainRepository(database)
	zr := repository.NewZoneRepository(database)
	fr := repository.NewConfigGeneratorRepo()
	ds := admin.NewDomainService(dr, zr, fr)
	ah := controller.NewAdminHandler(ds)

	server := srv.NewServer(srv.DeliveryContainer{
		Handler: ah,
	})

	quit := make(chan os.Signal)
	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf("error in server running %v", err.Error())
		}

	}()

	// gracefully shutdown
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server GraceFully ....... ")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Error In Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 3 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 3 seconds.")
	}
	log.Println("Server exiting")
}
