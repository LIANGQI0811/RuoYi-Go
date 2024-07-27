package di

import (
	"RuoYi-Go/config"
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/pkg/cache"
	"RuoYi-Go/pkg/logger"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"reflect"
)

var (
	app *SuperContainer
)

type SuperContainer struct {
	appConfig *config.AppConfig
	logger    *zap.Logger
	mvc       *iris.Application
	beans     map[reflect.Type]interface{}
}

func NewSuperContainer(c config.AppConfig) (*SuperContainer, error) {
	log := logger.NewZapLogger(c.LogLevel)
	mvc := iris.New()
	freeCache := cache.NewFreeCacheClient(100 * 1024 * 1024)
	ins := make(map[reflect.Type]interface{})
	ins[reflect.TypeOf(freeCache)] = freeCache
	ins[reflect.TypeOf(c)] = c
	redis, err := cache.NewRedisClient(c, log)
	if err != nil {
		log.Error("try to get redis redis error", zap.Error(err))
	}
	ins[reflect.TypeOf(redis)] = redis
	cache := cache.NewFreeCacheClient(c.FreeCacheConfig.Size)
	ins[reflect.TypeOf(cache)] = cache
	db, err := dao.OpenDB(c)
	if err != nil {
		log.Error("try to connect database error", zap.Error(err))
	}
	app = &SuperContainer{
		appConfig: &c,
		logger:    log,
		mvc:       mvc,
		beans:     ins,
	}
	app.putBean(db)

	return app, nil

}

func (this *SuperContainer) putBean(bean interface{}) *SuperContainer {
	if bean == nil {
		this.logger.Error("we got and ")
		return this
	}
	this.beans[reflect.TypeOf(bean)] = bean
	return this
}

func (this *SuperContainer) PutBeanByType(clazz interface{}) *SuperContainer {
	_type := reflect.TypeOf(clazz)
	//this.logger.Info("try to instantiate  object of type [%s] :", zap.String("type_name", _type.String()))
	value := reflect.New(_type)
	elem := value.Elem()
	for i := 0; i < _type.NumField(); i++ {
		field := _type.Field(i)
		val := this.beans[field.Type]
		elem.FieldByName(field.Name).Set(reflect.ValueOf(val).Elem())
	}
	this.beans[_type] = value
	return this
}

func (this *SuperContainer) getBeanByType(_type reflect.Type) interface{} {
	return this.beans[_type]
}
