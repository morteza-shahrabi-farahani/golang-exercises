package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)


type KeyboardButton struct {
	Text string `json:"text"`
	RequestContact bool `json:"request_contact"`
	RequestLocation bool `json:"request_location"`
}

type InlineKeyboardButton struct {
	Text string `json:"text"`
	CallbackData string `json:"callback_data"`
	Url string `json:"url"`
}

type ReplyMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
	Keyboard [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard bool `json:"resize_keyboard"`
	OnTimeKeyboard bool `json:"one_time_keyboard"`
	Selective bool `json:"selective"`
}


type SendMessage struct {
	ChatID           interface{} `json:"chat_id"`
	Text             string `json:"text"`
	ParseMode        string `json:"parse_mode"`
	ReplyMarkup      interface{} `json:"reply_markup"`
}

type MessageError struct {
    message string
}


func (e *MessageError) Error() string {
	return fmt.Sprintf(e.message)
}


func ReadSendMessageRequest(fileName string) (*SendMessage, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		//fmt.Println(err)
	}
	//fmt.Println("Successfully Opened file")
	var sendMessage SendMessage
	var marshalReply ReplyMarkup

	err2 := json.Unmarshal([]byte(data), &sendMessage)
	if err2 != nil {
		//fmt.Println(err2)
	}
	// fmt.Println("text is : " + sendMessage.Text)
	// fmt.Println("parse mode is : " + sendMessage.ParseMode)
	// fmt.Println("id is: " + fmt.Sprintf("%v", sendMessage.ChatID))
	// fmt.Println("reply is: " + fmt.Sprint(sendMessage.ReplyMarkup))
	// fmt.Println(sendMessage)

	reply, err3 := json.Marshal((sendMessage.ReplyMarkup))
	if err3 != nil {
		//fmt.Println(err3)
	}
	//fmt.Println(reply)
	err4 := json.Unmarshal(reply, &marshalReply)
	if err4 != nil {
		//fmt.Println(err4)
	}
	//fmt.Println(marshalReply)
	sendMessage.ReplyMarkup = marshalReply
	if sendMessage.ChatID == nil {
		return nil, &MessageError{message: "chat_id is empty"}
	}
	if sendMessage.Text == "" {
		return nil, &MessageError{message: "text is empty"}
	}
	return &sendMessage, nil
}

// func main() {
// 	ReadSendMessageRequest("input_sample2.json")
// }
