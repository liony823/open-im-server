package mcache

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/openimsdk/open-im-server/v3/pkg/common/storage/cache"
)

var (
	globalOnlineCache cache.OnlineCache
	globalOnlineOnce  sync.Once
)

func NewOnlineCache() cache.OnlineCache {
	globalOnlineOnce.Do(func() {
		globalOnlineCache = &onlineCache{
			user:           make(map[string]map[int32]struct{}),
			userOnlineTime: make(map[string]string),
		}
	})
	return globalOnlineCache
}

type onlineCache struct {
	lock           sync.RWMutex
	user           map[string]map[int32]struct{}
	userOnlineTime map[string]string
}

func (x *onlineCache) GetOnline(ctx context.Context, userID string) ([]int32, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	pSet, ok := x.user[userID]
	if !ok {
		return nil, nil
	}
	res := make([]int32, 0, len(pSet))
	for k := range pSet {
		res = append(res, k)
	}
	return res, nil
}

func (x *onlineCache) SetUserOnline(ctx context.Context, userID string, online, offline []int32) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	pSet, ok := x.user[userID]
	if ok {
		for _, p := range offline {
			delete(pSet, p)
		}
	}
	if len(online) > 0 {
		if !ok {
			pSet = make(map[int32]struct{})
		for _, p := range online {
			pSet[p] = struct{}{}
			x.userOnlineTime[userID] = strconv.FormatInt(time.Now().Unix(), 10)
		}
		}
	}
	if len(pSet) == 0 {
		delete(x.user, userID)
	}
	return nil
}

func (x *onlineCache) GetAllOnlineUsers(ctx context.Context, cursor uint64) (map[string][]int32, uint64, error) {
	if cursor != 0 {
		return nil, 0, nil
	}
	x.lock.RLock()
	defer x.lock.RUnlock()
	res := make(map[string][]int32)
	for k, v := range x.user {
		pSet := make([]int32, 0, len(v))
		for p := range v {
			pSet = append(pSet, p)
		}
		res[k] = pSet
	}
	return res, 0, nil
}
func (x *onlineCache) GetOnlineTime(ctx context.Context, userID string) (int64, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	tStr, ok := x.userOnlineTime[userID]
	if !ok {
		return 0, nil
	}
	tInt, err := strconv.ParseInt(tStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return tInt, nil
}
