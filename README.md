A package for [Amazon Advertising API](https://advertising.amazon.com/API/docs/en-us/)

## Installing
```shell script
go get -u github.com/gzp233/amazon_advertising
```

This will retrieve the library.

## example

```go
package main

import (
	"fmt"

	"github.com/gzp233/amazon_advertising"
)

func main() {
	c, err := amazon_advertising.NewClient(amazon_advertising.ClientOption{
		Region:       "NA",
		ClientId:     "amzn1.application-oa2-client.xxxxxxxxxx",
		ClientSecret: "xxxxxxxxxxxxxxxx",
		RefreshToken: "Atzr|cxxxxxxxxxxxxxxxxxxxxxxxxxx",
	})
	if err != nil {
		fmt.Println(err)
	}

	c.SetProfileId(12345678901111)
	c.SetDebug(true)
	r, err := c.SpListAdGroupsEx()
	if err != nil {
		fmt.Println(err)
	}
	if r != nil {
		fmt.Printf("%s %d", r.Body, r.StatusCode)
	}
}
```

