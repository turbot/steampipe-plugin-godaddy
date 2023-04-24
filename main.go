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
	"log"

	"github.com/alyx/go-daddy/daddy"
)

func main() {
	// OTE env
	// oteApiKey := "3mM44UcgzrkBAX_LDn59wMWcKiDphbVdakhdjaljsd"
	// oteApiSecret := "dx2sLKnreubBDBdsakjhda"
	// oteApiUrl := "https://api.ote-godaddy.com"

	// PROD Environment
	prodApiKey := "gHVbRVFtAFXW_TmcsFgxJQBvLjEkfdjhak"
	prodSecretKey := "Cmm6sqEBA4vVbUD8Ddsakh"
	// prodApiUrl := "https://api.godaddy.com"

	client, err := daddy.NewClient(prodApiKey, prodSecretKey, false)
	if err != nil {
		panic(err.Error())
	}

	myDomains, err := client.Domains.List(nil, nil, 2, "", nil, "")
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range myDomains {
		log.Println(value.Domain)
	}
	// Output ---- [{2023-04-24 05:07:28 +0000 UTC bootcloudlab.in 394663659 false 2024-04-24 05:07:28 +0000 UTC false false true [] false false true ACTIVE 2023-06-23 05:07:28 +0000 UTC false}]

}
