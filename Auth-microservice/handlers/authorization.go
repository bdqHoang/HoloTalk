package handlers

import (
	"Auth-microservice/cache"
	"context"
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func Enforcer(ctx context.Context, sub, obj, act string, adapper *gormadapter.Adapter) (resp interface{}, err error) {
	// create key for redis
	key := sub + obj + act

	// check key in redis
	result, err := cache.RedisClient.Get(ctx, key).Result()
	if err == nil {
		boolValue, err := strconv.ParseBool(result)
		if err != nil {
			return nil, err
		}
		return boolValue, nil
	}

	// reload casbin policy
	enforcer, err := casbin.NewEnforcer("auth_model.conf", adapper)
	if err != nil {
		return false, fmt.Errorf("error reloading policy: %v", err)
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("error loading policy: %v", err)
	}

	// check policy
	resp, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, fmt.Errorf("error checking policy: %v", err)
	}
	

	// save result to redis
	cache.RedisClient.Set(ctx, key, strconv.FormatBool(resp), 0).Err()

	return resp, nil
}