package main

import (
	"errors"
	"io"
)

func readBytes(r io.Reader, need uint32) ([]byte, error) {
	buf := make([]byte, need)

	for {
		n, err := r.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				// EOFの場合は読み出したデータが足りているかチェック
				if uint32(n) >= need {
					return buf, nil
				}
				return nil, io.ErrUnexpectedEOF
			}
			return nil, err
		}

		if uint32(n) >= need {
			return buf[:n], nil
		}
	}
}

func readUint24(b []byte) uint32 {
	_ = b[2]

	// 3バイト分のバイト列から整数値の変換はbinaryパッケージにないので自前で定義
	// MySQL Packageのintはリトルエディアンかつ実行環境もリトルエディアンなのでそのまま詰め替え
	// 重みの分だけ左シフトしてずらしたあとにORでビットを立てる
	//            10100000
	// 1 00001100 00000000
	// 1 00000000 00000000
	// 1 10000110 10100000
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16
}
