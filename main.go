//https://pkg.go.dev/github.com/alyx/go-daddy/daddy
// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/alyx/go-daddy/daddy"
// )

// func main() {
// 	client, err := daddy.NewClient("gHVbRVFtAFXW_TmcsFgxJQBvLjE5LD4d77m", "Cmm6sqEBA4vVbUD8DSejQy", false)
// 	fmt.Println("My connection ---", client)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	myDomains, err := client.Domains.List(nil, nil, 0, "", nil, "")
// 	fmt.Println("My Domains ---", myDomains)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, value := range myDomains {
// 		log.Println(value.Domain)
// 	}

// My connection --- &{   gHVbRVFtAFXW_TmcsFgxJQBvLjE5LD4d77m Cmm6sqEBA4vVbUD8DSejQy https://api.godaddy.com {0xc0000ec2c0} 0xc0000ec320 0xc0000ec320 0xc0000ec320 0xc0000ec320 0xc0000ec320 0xc0000ec320 0xc0000ec320 0xc0000ec320 0xc0000ec320}
// My Domains --- [{ {{{     }        } {{     }        } {{     }        } {{     }        }} 2023-04-24T05:07:28.000Z   bootcloudlab.in 394663659 false 2024-04-24T05:07:28.000Z false true [] false false 2024-06-07T22:07:23.000Z true ACTIVE false}]
// 2023/04/24 12:19:39 bootcloudlab.in

//}

// https://pkg.go.dev/github.com/oze4/godaddygo#section-readme
package main

import (
	"context"
	"fmt"

	"github.com/oze4/godaddygo"
)

func main() {
	key := "gHVbRVFtAFXW_TmcsFgxJQBvLjE5LD4d77m"
	secret := "Cmm6sqEBA4vVbUD8DSejQy"

	// Target production GoDaddy API
	// 99% of the time this is what you are looking for
	api, err := godaddygo.NewProduction(key, secret)
	if err != nil {
		panic(err.Error())
	}

	// Target version 1 of the production API
	godaddy := api.V1()

	//
	// See `Extended Example` section below for more
	//
	output, err := godaddy.ListDomains(context.Background())
	fmt.Println("Output ----", output)
	if err != nil {
		panic(err.Error())
	}

// Output ---- [{2023-04-24 05:07:28 +0000 UTC bootcloudlab.in 394663659 false 2024-04-24 05:07:28 +0000 UTC false false true [] false false true ACTIVE 2023-06-23 05:07:28 +0000 UTC false}]

}
