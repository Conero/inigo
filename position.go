// position
package main

import (
	"fmt"
)
import (
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
)

/*
	ip信息
*/
func ip_info() string {
	dial := "baidu.com:80"
	conn, err := net.Dial("udp", dial)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()
	var ip4dial = conn.LocalAddr().String() //保号时的端口号和/IP
	var ip string = strings.Split(ip4dial, ":")[0]
	return ip
}

/*
	地理信息
*/
func pos_info() string {
	ipmd := ""
	dosres, err := http.Get("http://httpbin.org/ip") //获取公网IP(末端)
	if err == nil {
		defer dosres.Body.Close()
		body, Rerr := ioutil.ReadAll(dosres.Body)
		if Rerr == nil {
			d1 := json_decode(string(body))
			d1Ip := d1["origin"].(string)
			dosres1, err2 := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + d1Ip)
			if err2 == nil {
				defer dosres1.Body.Close()
				body1, Rerr1 := ioutil.ReadAll(dosres1.Body)
				if Rerr1 == nil {
					d2 := json_decode(string(body1)) //strconv.Itoa
					d2cld := jsua(d2["data"])        //实验1
					ipmd += "\r\ncode:" + strconv.Itoa(int(d2["code"].(float64)))
					for d2ck, d2cv := range d2cld {
						ipmd += "\r\n" + d2ck + ": " + d2cv.(string)
					}
				}
			} else {
				fmt.Println(err2)
			}
		}
		//fmt.Println(body) //[123 10 32 32 34 111 114 105 103 105 110 34 58 32 34 50 50 50 46 49 55 49 46 50 53 48 46 53 56 34 10 125 10]
	}
	return ipmd
}
func pos_guaqi() {
	fmt.Println("Hello World!")
}
