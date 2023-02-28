package repo

import (
	"strings"
	"time"
	"user/pkg/logger"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	log "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type MysqlFactory struct{}

type UserModel struct{
	db *gorm.DB
	log logger.Interface
}

func (f *MysqlFactory) New(l logger.Interface) (UserRepository,error) {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,      // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}), &gorm.Config{
		Logger: log.Default,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		l.Error("Set database error.")
		return nil,err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	return &UserModel{
		db: db,
		log: l,
	},nil
}

func (u *UserModel) Create(user *User) error {
	err := u.db.Create(user).Error
	if err != nil {
		u.log.Error("failed to create user:", err)
		return err
	}
	return nil
}

func (u *UserModel) Update(user *User) error {
	err := u.db.Save(user).Error
	if err != nil {
		u.log.Error("failed to update user:", err)
		return err
	}
	return nil
}

func (u *UserModel) Delete(id int) error {
	user := &User{UserID: uint(id)}
	err := u.db.Delete(user).Error
	if err != nil {
		u.log.Error("failed to delete user:", err)
		return err
	}
	return nil
}

func (u *UserModel) GetByID(id int) (*User, error) {
	user := &User{}
	err := u.db.First(user, id).Error
	if err != nil {
		u.log.Error("failed to get user by id:", err)
		return nil, err
	}
	return user,nil
}

func (u *UserModel) GetByName(name string) (*User, error) {
	user := &User{}
	err := u.db.Where("user_name=?", name).First(&user).Error
	if err != nil {
		u.log.Error("failed to get user by name:", err)
		return nil, err
	}
	return user,nil
}