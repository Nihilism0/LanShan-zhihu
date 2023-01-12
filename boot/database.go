package boot

import (
	g "CSAwork/global"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"log"
	"time"
)

func InitDB() {
	//链接数据库
	dsn := ""
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	g.GlobalDb1 = db
}

//func MysqlDBSetup() {
//	config := g.Config.DataBase.Mysql
//	db, err := gorm.Open(mysql.Open(config.GetDsn()))
//	if err != nil {
//		g.Logger.Fatal("initialize mysql failed.", zap.Error(err))
//	}
//	sqlDB, _ := db.DB()
//	sqlDB.SetConnMaxIdleTime(g.Config.DataBase.Mysql.GetConnMaxIdleTime())
//	sqlDB.SetConnMaxLifetime(g.Config.DataBase.Mysql.GetConnMaxLifeTime())
//	sqlDB.SetMaxIdleConns(g.Config.DataBase.Mysql.MaxIdleConns)
//	sqlDB.SetMaxOpenConns(g.Config.DataBase.Mysql.MaxOpenConns)
//	err = sqlDB.Ping()
//	if err != nil {
//		g.Logger.Fatal("connect to mysql db failed.", zap.Error(err))
//	}
//	g.GlobalDb1 = db
//	g.GlobalDb1.AutoMigrate(&model.User{}, &model.Question{}, &model.Answer{})
//	g.Logger.Info("initialize mysql successfully!")
//}

func RedisSetup() {
	config := g.Config.DataBase.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Password: config.Password,
		DB:       config.Db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	g.Ctx = ctx
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		g.Logger.Fatal("connect to redis instance failed.", zap.Error(err))
	}
	g.RedisDb = rdb
	g.Logger.Info("initialize redis client successfully!")
}
func MongoDB() {
	clientOptions := options.Client().ApplyURI("")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	g.MongoDBClient = client
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
}
