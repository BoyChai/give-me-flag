package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var flag []string

func init() {
	file, err := os.Open("Batch_submit_flags/flag.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			return
		}
		flag = append(flag, string(line))
	}
}

func main() {
	push := submit{
		Host: "192.16.1.150",
		Port: "8080",
		Uri:  "flag_file.php",
	}
	push.GET()
	//push.POST()
}

type submit struct {
	Host string
	Port string
	Uri  string
}

func (s *submit) GET() {
	for _, v := range flag {
		resp, _ := http.Get(fmt.Sprintf("http://%s:%s/%s?token=team1&flag=%s",
			s.Host,
			s.Port,
			s.Uri,
			v,
		))
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
			resp.Body.Close()
			continue
		}
		fmt.Println("返回状态码====", resp.StatusCode)

	}
}

func (s *submit) POST() {
	for _, v := range flag {
		data := make(url.Values)
		data.Add("token", "team1")
		data.Add("flag", v)
		payload := data.Encode()
		resp, _ := http.Post(fmt.Sprintf("http://%s:%s/%s",
			s.Host,
			s.Port,
			s.Uri,
		), "application/x-www-form-urlencoded", strings.NewReader(payload))
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
			continue
		}
		fmt.Println("返回状态码====", resp.StatusCode)

	}

}
