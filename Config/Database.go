package Config

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
)

//func NewDB() *sql.DB {
//	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123456 dbname=DBbarang sslmode=disable")
//	Constant.PanicIfError(err)
//
//	db.SetMaxIdleConns(5)
//	db.SetMaxOpenConns(20)
//	db.SetConnMaxLifetime(60 * time.Minute)
//	db.SetConnMaxIdleTime(10 * time.Minute)
//
//	return db
//}

var (
	DB *gorm.DB
)

func InitDB() {
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"))
	con, _ := gorm.Open(postgres.Open(DSN), &gorm.Config{

		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	DB = con

	g := gen.NewGenerator(gen.Config{
		OutPath:      "Model/Database",
		OutFile:      "dto",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		WithUnitTest: true,
		ModelPkgPath: "Database",
	})

	g.UseDB(DB)

	// generate Database if the Database is used only
	// don't generate unused Database

	// Generates All Table in Database
	g.GenerateAllTable()

	// Generate Specify Table in Database
	//g.GenerateModel("master_type_schedules")

	g.Execute()
}
