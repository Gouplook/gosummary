package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

//var resDir = flag.String("o", "", "the res dir")
func HttpClient(method string, url string, params []byte, headParams map[string]string) ([]byte, error) {
	// 会话代理
	client := http.Client{
		Timeout: 20 * time.Second,
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}
	// set head of request
	for k, v := range headParams {
		request.Header.Set(k, v)
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func main() {
	// data := map[int64]float64{85899345922: 1000, 90194313218: 1000, 180388626434: 1000, 184683593730: 1000, 274877906946: 1000, 279172874242: 1000, 304942678018: 1000, 309237645314: 1000, 313532612610: 1000, 317827579906: 1000, 322122547202: 1000, 326417514498: 1000}
	// for lid, weight := range data {
	//  fmt.Printf("seg:%v-seq:%v-dir:%v-%v\n",
	//      lid>>32,
	//      int((lid&0xffffffff)>>2),
	//      lid&0x3,
	//      weight)
	// }
	raw := []byte("{\"global_project_id\":\"\",\"map_version\":-1,\"map_id\":20221219,\"points\":[{\"long\":299727.88,\"lat\":2469359.2,\"theta\":0.44},{\"long\":300841.333844,\"lat\":2.468671195618e+06,\"theta\":3.39}]}")
	group := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			group.Add(1)
			defer group.Done()
			pre := time.Now().UnixMicro()
			_, err := HttpClient("POST", "http://localhost:8888/mapservice/planning/get/path", raw, nil)
			if err != nil {
				fmt.Println("request error:", err.Error())
			} else {
				delta := time.Now().UnixMicro() - pre
				fmt.Println(i, " request time:", float64(delta)/1e6)
			}
		}(i)
	}
	group.Wait()

}
