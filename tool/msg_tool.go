package tool

import (
	"fmt"
	"net/http"
	"net/url"
	. "snow/common"
	"strconv"
	"time"
)

var RemoteHttp = "127.0.0.1:8111"

var Num = 100
var InitPort = 40000

// 发送HTTP GET请求
var client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        10,               // 允许最多 200 个空闲连接
		MaxIdleConnsPerHost: 10,               // 每个主机最多 10 个空闲连接
		IdleConnTimeout:     60 * time.Second, // 60 秒超时
	},
}

func SendHttp(from string, target string, data []byte, k int) {
	if data[1] == UserMsg || data[1] == IHAVE || data[1] == ReliableMsgAck {
		port := GetPortByIp(target)
		//不给浮动的节点发送请求
		if port >= InitPort+Num {
			//fmt.Print(InitPort)
			return
		} else if port == InitPort {
			return
		}
		values := url.Values{}
		values.Add("From", from)
		values.Add("Target", target)
		values.Add("Size", fmt.Sprintf("%d", len(data)))
		if data[0] == ColoringMsg || data[0] == RegularMsg || data[0] == ReliableMsg {
			values.Add("Id", string(data[TagLen+IpLen*2:TagLen+IpLen*2+TimeLen]))
		} else if data[0] == EagerPush {
			values.Add("Id", string(data[TagLen+IpLen:TagLen+IpLen+TimeLen]))
		} else if data[0] == LazyPush || data[0] == Graft {
			//合并同一个id的消息，因为这些也是计算进消息大小的
			//data[0] = EagerPush
			values.Add("Id", string(data[TagLen:TagLen+TimeLen]))
		} else {
			//其他的消息都没有附带ip
			values.Add("Id", string(data[TagLen:TagLen+TimeLen]))
		}
		values.Add("FanOut", strconv.Itoa(k))
		values.Add("Num", strconv.Itoa(Num))

		values.Add("MsgType", strconv.Itoa(int(data[0])))
		values.Add("Size", fmt.Sprintf("%d", len(data)))

		baseURL := "http://" + RemoteHttp + "/putRing"
		fullURL := fmt.Sprintf("%s?%s", baseURL, values.Encode())

		client.Get(fullURL)
	}

}
