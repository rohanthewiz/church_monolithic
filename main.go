package main

// go:generate sqlboiler postgres

import (
	"github.com/rohanthewiz/church/chweb/config"
	"github.com/rohanthewiz/church/chweb/db"
	"github.com/rohanthewiz/church/chweb"
	. "github.com/rohanthewiz/logger"
	"github.com/rohanthewiz/roredis"
	"github.com/sirupsen/logrus"
	"gopkg.in/relistan/rubberneck.v1"
)

var Version, GitCommitHash, BuildTimestamp string  // buildtime vars

func init() {
	config.InitConfig(Version, GitCommitHash, BuildTimestamp)
}

func main() {
	InitLog(LogOptions{
		AppName: config.APP_NAME,
		Environment: config.AppEnv,
		Format: config.Options.Log.Format,
		Level: config.Options.Log.Level,
		InfoPath: config.Options.Log.InfoPath,
		ErrorPath: config.Options.Log.ErrorPath,
	})
	defer CloseLog()

	err := db.InitDB(db.DBOpts{
		DBType: db.DBTypes.Postgres,
		Host: config.Options.PG.Host,
		Port: config.Options.PG.Port,
		User: config.Options.PG.User,
		Word: config.Options.PG.Word,
		Database: config.Options.PG.Database,
	})
	if err != nil {
		LogErr(err, "Could not setup database")
	}
	defer db.CloseDB()

	roredis.InitRedis(roredis.RedisCfg{
		Host: "localhost", DB: 0, // default -> Port: "6379",
	})

	logrus.Println(config.APP_NAME, "is starting...")
	rubberneck.NewPrinter(logrus.Infof, rubberneck.NoAddLineFeed).Print(config.Options) // print config to log

	Log("Info", "Version Info", "Version", config.Version, "CommitHash", config.GitCommitHash,
			"BuildTimestamp", config.BuildTimestamp)

	chweb.Serve()
}
