package pili

import (
	"fmt"
	"runtime"
)

var (
	VERSION = "1.0.1"

	API_HOST          = "pili.qiniuapi.com"
	RTMP_PUBLISH_HOST = "pub.z1.glb.pili.qiniup.com"
	RTMP_PLAY_HOST    = "live.z1.glb.pili.qiniucdn.com"
	HLS_PLAY_HOST     = "hls1.z1.glb.pili.qiniuapi.com"
)

func UserAgent() string {
	return fmt.Sprintf("pili-sdk-go/%s %s %s/%s", VERSION, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}