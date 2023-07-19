package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/trinodb/trino-go-client/trino"
	"net"
	"net/http"
	"time"
)

type Track struct {
	//vin, datetime, timestamp, params.attached_map, params.attached_map_version, params.mov_objs
	Vin        string    `json:"vin"`
	Datetime   time.Time `json:"datetime"`
	TimeStamp  float64   `json:"timestamp"`
	MapId      int64     `json:"params.attached_map"`
	MapVersion int64     `json:"params.attached_map_version"`
	MovObjs    string    `json:"params.mov_objs"`
}

func main() {
	// 连接参数
	foobarClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig:       &tls.Config{},
		},
	}
	trino.RegisterCustomClient("foobar", foobarClient)
	//http://user@localhost:8080?source=hello&catalog=default&schema=foobar
	db, err := sql.Open("trino", "https://hsz10608:Hsz10608$@trino.uisee.com:30310?catalog=elasticsearch&schema=default")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("select vin, datetime, timestamp, params.attached_map, params.attached_map_version, params.mov_objs from \"20230710-20230716_tracking\" where vin='robobus-m-byd.car1' and datetime >= cast('2023-07-16 00:00:00.000' as timestamp) and datetime <= cast('2023-07-16 01:00:00.000' as timestamp) limit 10")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var track Track
		err := rows.Scan(&track.Vin, &track.Datetime, &track.TimeStamp, &track.MapId, &track.MapVersion, &track.MovObjs)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("track:", track)
	}
}
