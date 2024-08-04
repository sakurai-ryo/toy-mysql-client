package main

import (
	"log/slog"
)

func main() {
	//conn, err := sql.Open("mysql", "root:root")
	//if err != nil {
	//	slog.Error(err.Error())
	//	os.Exit(1)
	//}
	//defer conn.Close()
	//
	//_, err = conn.Exec("SELECT 1")
	//if err != nil {
	//	slog.Error(err.Error())
	//}

	d := &Driver{}
	_, err := d.Open("root:root")
	if err != nil {
		slog.Error(err.Error())
	}

	// 100000
	//data := []byte{0b10100000, 0b10000110, 0b00000001}
	//fmt.Printf("data: %b\n", data)
	//fmt.Printf("uint32(data[0]): %b\n", uint32(data[0]))
	//fmt.Printf("uint32(data[1]): %b\n", uint32(data[1]))
	//fmt.Printf("uint32(data[1])<<8: %b\n", uint32(data[1])<<8)
	//fmt.Printf("uint32(data[2]): %b\n", uint32(data[2]))
	//fmt.Printf("uint32(data[2])<<16: %b\n", uint32(data[2])<<16)
	//pktLen := int(uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16)
	//fmt.Println(pktLen)
	//fmt.Printf("OR: %b\n", pktLen)
}
