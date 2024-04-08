package operate

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-redis/redis/v8"
)

func redisCommandExecuteError(cmd string, err error, args ...string) string {
	return fmt.Sprintf("ERROR: execute command '%v %v' occurs error, %v", cmd, strings.Join(args, " "), err.Error())
}

func redisCommandExecuteResult[T any](cmd string, result T, args ...string) string {
	return fmt.Sprintf("INFO: execute command '%v %v' result, %v", cmd, strings.Join(args, " "), result)
}

func connectRedis(urlString string) (redis.Cmdable, error) {
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
	rdb, err := connectRedis(argURL)
	if err != nil {
		fmt.Println("ERROR: connect redis with url", argURL, "occurs error,", err.Error())
		return
	}

	matchKeySlice, err := rdb.Keys(context.Background(), argRegexp).Result()
	if err != nil {
		fmt.Println(redisCommandExecuteError("keys", err, argRegexp))
		return
	}

	if len(matchKeySlice) == 0 {
		fmt.Println(redisCommandExecuteResult("keys", matchKeySlice, argRegexp))
		return
	}

	for _, key := range matchKeySlice {
		switch option {
		case "del":
			result, err := rdb.Del(context.Background(), key).Result()
			if err != nil {
				fmt.Println(redisCommandExecuteError("del", err, key))
				return
			}
			fmt.Println(redisCommandExecuteResult("del", result, key))
		case "hgetall":
			result, err := rdb.HGetAll(context.Background(), key).Result()
			if err != nil {
				fmt.Println(redisCommandExecuteError("hgetall", err, key))
				return
			}
			fmt.Println(redisCommandExecuteResult("hgetall", result, key))
		default:
			fmt.Println(redisCommandExecuteResult("keys", matchKeySlice, argRegexp))
		}
	}
}
