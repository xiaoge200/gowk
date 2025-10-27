# gowk


gowk 是 WuKongIM REST API 的 Go 语言 SDK。

[wukongim api docs](https://docs.githubim.com/zh/api)

### example

``` go
package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
    "github.com/xiaoge200/gowk"
)

func main() {
	baseUrl := "xxx" // wukongim im api url

	client := gowk.NewClient(baseUrl)

    ctx := context.Background()
    resp,err := client.SendMessage(ctx,gowk.Message{
        Header: gowk.Header{
            NoPersist: gowk.Ptr(0),
			RedDot:    gowk.Ptr(1),
			SyncOnce:  gowk.Ptr(0),
        },
        FromUID:     "123",
        ChannelID:   "321",
        ChannelType: gowk.ChannelType_Person,
        Payload:     base64.StdEncoding.EncodeToString([]byte("hello"))
    })
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", result)
}
```