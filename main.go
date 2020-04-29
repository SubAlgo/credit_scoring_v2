package main // import "github.com/subalgo/credit_scoring_v2"

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/subalgo/credit_scoring_v2/internal/app/admin"
	"github.com/subalgo/credit_scoring_v2/internal/app/auth"
	"github.com/subalgo/credit_scoring_v2/internal/app/forgotPassword"
	"github.com/subalgo/credit_scoring_v2/internal/app/location"
	"github.com/subalgo/credit_scoring_v2/internal/app/permissionSetting"
	"github.com/subalgo/credit_scoring_v2/internal/app/province"
	"github.com/subalgo/credit_scoring_v2/internal/app/questionnaire"
	"github.com/subalgo/credit_scoring_v2/internal/app/questionnaireOption"

	_ "github.com/lib/pq"
	"github.com/subalgo/credit_scoring_v2/internal/app/user"
	"github.com/subalgo/credit_scoring_v2/internal/pkg/dbctx"
	"log"
	"net/http"
	"os"
)

func main() {
	// connStr := "postgres://dbUser:dbPassword@dbHost/dbName?sslmode=<verify-full | disable>"

	dbURL := os.Getenv("CREDIT_SCORING_DB_URL")

	if dbURL == "" {
		dbURL = "postgres://localhost:5432/credit_scoring_v2?sslmode=disable"
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can not connect to database;", err)
	}
	defer db.Close()

	redisAddr := os.Getenv("CREDIT_SCORING_REDIS_ADDR")
	redisPassword := os.Getenv("CREDIT_SCORING_REDIS_PASSWORD")

	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:       redisAddr,
		Password:   redisPassword,
		MaxRetries: 3,
	})

	defer redisClient.Close()

	auth.RedisClient = redisClient
	mux := http.NewServeMux()
	mux.Handle("/", http.NotFoundHandler())
	mux.Handle("/admin/", http.StripPrefix("/admin", admin.Handler()))
	mux.Handle("/auth/", http.StripPrefix("/auth", auth.Handler()))
	mux.Handle("/user/", http.StripPrefix("/user", user.Handler()))
	mux.Handle("/questionnaire/", http.StripPrefix("/questionnaire", questionnaire.Handler()))
	mux.Handle("/permission_setting/", http.StripPrefix("/permission_setting", permissionSetting.Handler()))
	mux.Handle("/questionnaireOption/", http.StripPrefix("/questionnaireOption", questionnaireOption.Handler()))
	mux.Handle("/location/", http.StripPrefix("/location", location.Handler()))
	mux.Handle("/province/", http.StripPrefix("/province", province.Handler()))
	mux.Handle("/forgot_password/", http.StripPrefix("/forgot_password", forgotPassword.Handler()))

	h := chain(
		dbctx.Middleware(db),
		auth.FetchAuth,
	)(mux)

	log.Println("staring server on :8000")
	err = http.ListenAndServe(":8000", h)
	if err != nil {
		log.Fatal(err)
	}
}

func chain(hs ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		for i := len(hs); i > 0; i-- {
			h = hs[i-1](h)
		}
		return h
	}
}
