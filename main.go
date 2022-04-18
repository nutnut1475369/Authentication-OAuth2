package main

import (
	mcontext "googleauth/middleware"
	"googleauth/router"
	"googleauth/service/db"
	"googleauth/service/jwt"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file ")
    }
}

func main() {
	dsn := "host=localhost user=postgres password=123456 dbname=pg_go_auth port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!= nil {
		panic("failed to connect database")
	}
	dbs, err := db.NewService(gdb);
	if err != nil {
		panic("failed to connect database")

	}
	if err := dbs.Migrate(); err != nil {
		panic("failed to connect database")
		
	}
	key := os.Getenv("SECRET_SESSION_KEY")
	maxAge := 86400 * 30  // 30 days
  	isProd := false       // Set to true when serving over https


	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true   // HttpOnly should always be enabled
	store.Options.Secure = isProd
  
	gothic.Store = store
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8080/auth/google/callback"),
	)
	router.Load(
		mcontext.SetDb(dbs),
	)
}