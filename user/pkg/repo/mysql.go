package repo

import (
	"fmt"
	"strings"
	"time"
	"github.com/cyicz123/todolist/user/pkg/logger"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	log "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// MysqlFactory implements the DBFactory interface for creating a MySQL user repository.
type MysqlFactory struct{}

// UserModel implements the UserRepository interface using a gorm.DB instance for MySQL.
type UserModel struct{
	db *gorm.DB
	log logger.Interface
}

// New creates a new UserModel instance that implements the UserRepository interface for a MySQL user repository.
// The method uses viper configuration library to configure the MySQL connection and connection pool settings.
func (f *MysqlFactory) New(l logger.Interface, v *viper.Viper) (UserRepository,error) {
	host := v.GetString("mysql.host")
	port := v.GetString("mysql.port")
	database := v.GetString("mysql.database")
	username := v.GetString("mysql.username")
	password := v.GetString("mysql.password")
	charset := v.GetString("mysql.charset")
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", "?charset=" + charset + "&parseTime=true"}, "")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Error("Set database error. err:", err)
		return nil,err
	}

	ret := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", database))
	if ret.Error != nil {
		return nil, ret.Error
	}

	dsn = strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")
	db, err = gorm.Open(mysql.New(mysql.Config{
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
		l.Error("Set database error. err:", err)
		return nil,err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	u := &UserModel{
		db: db,
		log: l,
	}
	f.migration(u)
	return u,nil
}

// migration performs an automatic migration of the User table in the database using the gorm.AutoMigrate method.
// The method sets the table options to utf8mb4 and logs success or panics on failure.
func (f *MysqlFactory) migration(u *UserModel) {
	//自动迁移模式
	err := u.db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&User{},
		)
	if err != nil {
		u.log.Panic("register table fail")
	}
	u.log.Info("register table success")
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

func (u *UserModel) Delete(name string) error {
	user := &User{UserName: name}
	err := u.db.Where("user_name=?", name).First(user).Delete(user).Error
	if err != nil {
		u.log.Error("failed to delete user:", err)
		return err
	}
	return nil
}

func (u *UserModel) GetByID(id uint32) (*User, error) {
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