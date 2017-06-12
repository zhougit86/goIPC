package main

import(
	"encoding/json"
	"fmt"
	//"gamePlat/ipc"
)

func main()  {
	raw := make(map[int]string)
	raw[0]="shit"
	raw[1]="fuck"

	fmt.Println(raw)
	trans := make([]byte,5)
	trans,_=json.Marshal(raw)
	fmt.Println(string(trans))
	fmt.Printf("%T\n",trans)


	raw2:= make(map[int]string)
	ok := json.Unmarshal(trans,&raw2)

	fmt.Println(ok, raw2)

}
