package main

import "net/http"

func main() {
	url := "http://xxx.xx:xx:/"
	request, _ := http.NewRequest("get", url, nil)
	request.Header.Set("flag", "56 6d 31 77 52 32 46 72 4e 56 5a 4e 56 6c 70 70 55 6c 64 34 56 6c 6c 58 64 47 46 5a 56 6c 4a 59 59 33 70 47 61 6b 31 58 64 7a 4a 57 52 33 4d 31 59 56 5a 5a 65 46 4e 73 62 47 46 57 56 32 68 51 57 57 74 61 56 6d 51 78 54 6c 6c 6a 52 6d 68 58 59 6c 64 6f 55 56 5a 47 56 6d 46 6b 4d 57 52 48 56 6d 78 6f 59 56 49 7a 61 46 68 61 56 7a 45 77 54 6b 5a 6b 56 56 4e 75 54 6c 4a 4e 52 45 5a 4a 56 57 31 34 62 31 52 73 57 58 70 68 52 58 52 57 59 6b 64 53 64 6c 6c 71 52 6c 5a 6b 4d 58 42 47 56 32 78 47 56 6c 5a 45 51 54 55 3d")
}
