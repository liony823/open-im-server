package cachekey

import (
	"strings"
	"time"
)

const (
	OnlineKey           = "ONLINE:"
	OnlineChannel       = "online_change"
	OnlineExpire        = time.Hour / 2
	LatestOnlineTimeKey = "ONLINETIME:" // OWLIM 的 新加
)

func GetOnlineKey(userID string) string {
	return OnlineKey + userID
}

func GetOnlineKeyUserID(key string) string {
	return strings.TrimPrefix(key, OnlineKey)
}

/* OWLIM 的 新加 */

func GetLatestOnlineTimeKey(userID string) string {
	return LatestOnlineTimeKey + userID
}
