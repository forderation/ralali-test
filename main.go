package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/forderation/ralali-test/docs"
	"github.com/forderation/ralali-test/internal/delivery"
	"github.com/forderation/ralali-test/internal/repository"
	"github.com/forderation/ralali-test/internal/usecase"
	"github.com/forderation/ralali-test/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go.uber.org/mock/mockgen/model"
)

func main() {
	loadConfigFile()
	mySqlDB := initMysqlDB(viper.GetString("db_dsn"))
	cakeDBRepository := repository.NewCakeDBRepository(mySqlDB, viper.GetString("cakes_table"))
	cakeUsecase := usecase.NewCakeUsecase(cakeDBRepository)
	cakeDelivery := delivery.NewCakeDelivery(cakeUsecase)

	docs.SwaggerInfo.Title = "Ralali App"
	docs.SwaggerInfo.Description = "ralali cake demo app"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8081"
	docs.SwaggerInfo.Schemes = []string{"http"}

	routes := initRoute(cakeDelivery)
	address := viper.GetString("service_addr")
	srv := &http.Server{Addr: address, Handler: routes}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// gracefully shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown service ...")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	closeMySQLDB(ctx, mySqlDB)
	select {
	case <-ctx.Done():
	}
	log.Println("service exiting")
}

func loadConfigFile() {
	filePath := fmt.Sprintf("config.toml")
	viper.SetConfigType("toml")
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("using config file:", viper.ConfigFileUsed())
}

func initRoute(cakeDelivery *delivery.CakeDelivery) *gin.Engine {
	baseRoot := gin.Default()
	baseRoot.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	baseRoot.Use(util.CORSMiddleware())
	cakeRoutes := baseRoot.Group("/cakes")
	cakeRoutes.GET("", cakeDelivery.GetCakes)
	cakeRoutes.GET("/:id", cakeDelivery.GetCake)
	cakeRoutes.POST("", cakeDelivery.CreateCake)
	cakeRoutes.PUT("/:id", cakeDelivery.UpdateCake)
	cakeRoutes.DELETE("/:id", cakeDelivery.DeleteCake)
	return baseRoot
}

func initMysqlDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Panic("error init mysql db: ", err.Error())
	}
	return db
}

func closeMySQLDB(ctx context.Context, db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal("error on close mysqldb: ", err.Error())
	}
}
