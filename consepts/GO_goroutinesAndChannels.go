import (
    "fmt"
    "log"
    "sync"
    )
type Message struct {
    addition int
    multiplication int
} 
func getUser(a int,b int,ch chan<- *Message,wg *sync.WaitGroup) {
    c:=a+b
    ch<- & Message {
        addition: c,
    }
    wg.Done()
}
func getId(a int,b int,ch chan<- *Message,wg *sync.WaitGroup){
    c:=a*b
    ch<- & Message {
        multiplication: c,
    }
    wg.Done()
}
func main() {
  ch := make(chan *Message ,2)
  wg := &sync.WaitGroup{}
  wg.Add(2)
  a:=0
  fmt.Scan(&a)
  b:=0
  fmt.Scan(&b)
  go getUser(a,b,ch,wg)
  go getId(a,b,ch,wg)
  wg.Wait()
  close(ch)
  for mes := range ch {
      log.Println(mes)
  }
  
}

