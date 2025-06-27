package storage

import (
	"bufio"
	"encoding/json"
	"os"
	"simple-chat/model"
	"sync"
)

const MessageFile = "save_messages.txt"

var FileMutex sync.Mutex

func InitFile() error {
	_, err := os.Stat(MessageFile)
	if os.IsNotExist(err) {
		file, err2 := os.Create(MessageFile)
		if err2 != nil {
			return err2
		}
		defer file.Close()
	} else if err != nil {
		return err
	}
	return nil
}

func ReadMessages() ([]model.Message, error) {
	FileMutex.Lock()
	defer FileMutex.Unlock()
	file, err := os.Open(MessageFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var messages []model.Message
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var msg model.Message
		raw := scanner.Bytes()
		err := json.Unmarshal(raw, &msg)
		if err == nil {
			messages = append(messages, msg)
		}
	}
	return messages, scanner.Err()
}

func SaveMessage(msg model.Message) error {
	FileMutex.Lock()
	defer FileMutex.Unlock()
	f, err := os.OpenFile(MessageFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = f.Write(append(data, '\n'))
	return err
}
