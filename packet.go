package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"slices"
)

var ErrPacketTooLarge = errors.New("packet for query is too large")

func readPacketPayload(r io.Reader) ([]byte, error) {
	header, err := readBytes(r, 4)
	if err != nil {
		return nil, err
	}
	// 先頭3バイトがパケットの長さ
	packetLength := readUint24(header[:3])
	// 次の1バイトがシーケンス番号
	sequenceId := header[3]
	slog.Info("Header", "payload_length", packetLength, "sequence_id", sequenceId)

	payload, err := readBytes(r, packetLength)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func writePacket(w io.Writer, payload []byte) error {
	// TODO: MaxPackageSizeチェック
	// Payloadサイズが2^24-1をを超える場合はヘッダーのpayload_lengthを`0xffffff`にしないといけないが一旦無視

	headerLen := 4
	packetLen := len(payload) + headerLen

	header := make([]byte, headerLen)
	header[0] = byte(packetLen)
	header[1] = byte(packetLen >> 8)
	header[2] = byte(packetLen >> 16)
	header[3] = 0 // TODO: シーケンス番号

	_, err := w.Write(slices.Concat(header, payload))
	if err != nil {
		return err
	}

	return nil
}

func readHandshakePacket(r io.Reader) error {
	payload, err := readPacketPayload(r)
	if err != nil {
		return err
	}

	startPosition := 0

	// Protocol::HandshakeV10のパース
	// https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase_packets_protocol_handshake_v10.html
	protocolVersion := payload[0]

	// server_versionはNull終端
	// [low, high)なのでhigh側は+1
	serverVersionNullIndex := bytes.IndexByte(payload[1:], 0x00)
	serverVersion := payload[1 : serverVersionNullIndex+1]
	startPosition = serverVersionNullIndex + 1

	// 次の4バイトはconnection_id
	connectionId := binary.LittleEndian.Uint32(payload[startPosition : startPosition+4+1])
	startPosition = startPosition + 4

	// server_versionの次の8バイト
	fmt.Println("a: ", serverVersionNullIndex)
	fmt.Printf("%x\n", payload[serverVersionNullIndex])
	authPluginDataPart := payload[startPosition+1 : startPosition+8+1]
	startPosition = startPosition + 8 + 1

	slog.Info("Payload",
		"protocol_version", protocolVersion,
		"server_version", serverVersion,
		"connection_id", connectionId,
		"auth_plugin_data_part_1", authPluginDataPart,
	)
	fmt.Printf("%x\n", payload[startPosition])
	return nil
}

func writeHandshakeResponse(conn io.Writer) error {
	return nil
}
