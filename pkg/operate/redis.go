package operate

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-redis/redis/v8"
)

func connect(urlString string) (redis.Cmdable, error) {
	url, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	var rdb redis.Cmdable
	switch url.Scheme {
	case "cluster":
		clusterOptions := &redis.ClusterOptions{}
		clusterOptions.Username = url.User.Username()
		password, has := url.User.Password()
		if has {
			clusterOptions.Password = password
		}
		clusterOptions.Addrs = []string{url.Host}
		rdb = redis.NewClusterClient(clusterOptions)
	default:
		urlOption, err := redis.ParseURL(urlString)
		if err != nil {
			return nil, err
		}
		rdb = redis.NewClient(urlOption)
	}
	if rdb == nil {
		return nil, fmt.Errorf("create redis client with url %v option failed", urlString)
	}
	cmd := rdb.Ping(context.Background())
	if cmd == nil {
		return nil, fmt.Errorf("create redis ping command failed")
	}
	result, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	if result != "PONG" {
		return nil, fmt.Errorf("ping redis failed")
	}
	return rdb, nil
}

func OperateRedis(argURL, option, argRegexp string) {
	rdb, err := connect(argURL)
	if err != nil {
		fmt.Println("ERROR: connect redis with url", argURL, "occurs error,", err.Error())
		return
	}

	matchKeySlice, err := rdb.Keys(context.Background(), argRegexp).Result()
	if err != nil {
		fmt.Println("ERROR: execute command 'keys '", argRegexp, "'occurs error", err.Error())
		return
	}

	for _, key := range matchKeySlice {
		switch option {
		case "del":
			result, err := rdb.Del(context.Background(), key).Result()
			if err != nil {
				fmt.Println("ERROR: execute command 'del", key, "' occurs error,", err.Error())
				return
			}
			fmt.Println("INFO: execute command 'del", key, "' result", result)
		case "hgetall":
			result, err := rdb.HGetAll(context.Background(), key).Result()
			if err != nil {
				fmt.Println("ERROR: execute command 'hgetall", key, "' occurs error,", err.Error())
				return
			}
			fmt.Println("INFO: execute command 'hgetall", key, "' result", result)
		}
	}
}
