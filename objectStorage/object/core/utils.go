package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"object/service"
	"os"
)

// 校验并保存文件；
func checkAndSaveFile(request *service.FileBinary) error {
	//fileName := request.File.FileName
	tempHash := request.File.HashCode
	data := request.Data
	//fmt.Println(data)
	reader := bytes.NewReader(data)
	calHash := calculateHash(data)
	//fmt.Println(calHash)
	if tempHash != calHash {
		return errors.New("file hash mismatch, please try again later")
	}
	f, err := os.Create(fmt.Sprintf("./testData/data/%s", calHash))
	if err != nil {
		log.Println("object service create file failed: ", err)
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(f)
	_, err = io.Copy(f, reader)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// 计算文件的hash值；
func calculateHash(b []byte) string {
	hash := sha256.New()
	hash.Write(b)
	return hex.EncodeToString(hash.Sum(nil))
}

// 获取文件
func getFile(request *service.FileRequest) ([]byte, error) {
	f, err := os.ReadFile(fmt.Sprintf("./testData/data/%s", request.FileName))
	if err != nil {
		log.Println("read file failed...", err)
		return nil, err
	}
	return f, nil
}

//// 用于判断请求的文件是否存在，消息队列
//func locate() bool {
//	ch, err := MQ.Channel()
//	defer func(ch *amqp.Channel) {
//		err := ch.Close()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}(ch)
//	if err != nil {
//		log.Panicln("init rabbitMQ chan failed: ", err)
//	}
//	queue, _ := ch.QueueDeclare("info_queue", true, false, false, false, nil)
//	_ = ch.Qos(1, 0, false)
//	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
//	if err != nil {
//		log.Println("get msg from rabbitMQ failed: ", err)
//		return false
//	}
//
//	for msg := range msgs {
//		var fileName string
//		err := json.Unmarshal(msg.Body, &fileName)
//		if err != nil {
//			log.Println("unmarshall filename failed", err)
//			continue
//		}
//		//fmt.Println(fileName)
//		_, err = os.Stat(fmt.Sprintf("./testData/data/%s", fileName))
//		return !os.IsNotExist(err)
//	}
//	return false
//}
