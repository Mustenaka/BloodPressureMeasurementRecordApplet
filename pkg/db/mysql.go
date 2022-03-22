package db

import (
	"BloodPressure/pkg/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ IDataSource = (*defaultMysqlDataSource)(nil)

// IDataSource 定义数据库数据源接口，按照业务需求可以返回主库链接Master和从库链接Slave
type IDataSource interface {
	Master() *gorm.DB // 主数据库
	Slave() *gorm.DB  // 从数据库 - 如果需要使用从数据库需要重新编译
	Close()           // 关闭数据库（主从均关闭）
}

// defaultMysqlDataSource 默认mysql数据源实现
type defaultMysqlDataSource struct {
	master *gorm.DB // 定义私有属性，用来持有主库链接，防止每次创建，创建后直接返回该变量。
	slave  *gorm.DB // 同上，从库链接
}

func (d *defaultMysqlDataSource) Master() *gorm.DB {
	if d.master == nil {
		panic("The [master] connection is nil, Please initialize it first.")
	}
	return d.master
}

func (d *defaultMysqlDataSource) Slave() *gorm.DB {
	if d.master == nil {
		panic("The [slave] connection is nil, Please initialize it first.")
	}
	return d.slave
}

func (d *defaultMysqlDataSource) Close() {
	// 关闭主库链接
	if d.master != nil {
		m, err := d.master.DB()
		if err != nil {
			m.Close()
		}
	}
	// 关闭从库链接
	if d.slave != nil {
		s, err := d.slave.DB()
		if err != nil {
			s.Close()
		}
	}
}

// 初始化默认sql - Master
func NewDefaultMysql(c config.DBConfig) *defaultMysqlDataSource {
	return &defaultMysqlDataSource{
		master: connect(
			c.Username,
			c.Password,
			c.Host,
			c.Port,
			c.Dbname,
			c.Charset,
			c.ParseTime,
			c.Loc,
			c.MaximumPoolSize,
			c.MaximumIdleSize),
	}
}

// 数据库链接
func connect(user, password, host, port, dbname, charset, parseTime, loc string, maxPoolSize, maxIdle int) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		user, password,
		host, port, dbname,
		charset, parseTime, loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // 缓存每一条sql语句，提高执行速度
	})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetConnMaxLifetime(time.Hour)
	// 设置连接池大小
	sqlDb.SetMaxOpenConns(maxPoolSize)
	sqlDb.SetMaxIdleConns(maxIdle)
	return db
}
