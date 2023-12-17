package main

import (
	"context"
	"fmt"
)

func main() {

	//Boş Bir Context Nesnesi
	ctx := context.Background()

	//Context İçinde Key Value Şekilde Degerlerimizi Tutabiliriz
	ctx = context.WithValue(ctx, "any_key", "any_value")

	//Key degeri ile de okuyabilirz
	val := ctx.Value("any_key")
	fmt.Println(val)
}
