package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	//Basic Command
	dateCmd := exec.Command("date")

	//Get Output
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	//Out type:byte convert string
	fmt.Printf("Date:%v", string(dateOut))

	//Error Type Check
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	//grep:düz metin veri kümelerinde düzenli bir ifadeyle eşleşen satırları aramak için kullanılan bir komut satırı programıdı
	grepCmd := exec.Command("grep", "hello")

	//Input girdisi vermek için pipe alınır
	grepIn, _ := grepCmd.StdinPipe()

	//Outputu okumak için pipe alınır
	grepOut, _ := grepCmd.StdoutPipe()

	//Konsolde çalıştırılır
	grepCmd.Start()

	//Grep programı yukarda hangi kelime ile başlattıysak içinde geçen satırı bize doner
	//burda da input pipine degerlerimizi yazıyoruz
	grepIn.Write([]byte("hello grep\ngoodbye grep"))

	//input pipe kapatılır
	grepIn.Close()

	//Output uzerinden cıktı alınır
	grepBytes, _ := io.ReadAll(grepOut)

	//Verdigimiz inputun işlenmesi vakit alabilir bu durumda işlem tamamlana kadar bekleriz
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes)) //Çıktı yine bize byte formatında gelir biz stringe cevirirz

	//Bulundugu klasordeki dosyaların ayrıntılarını cıktı olarka veir
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
