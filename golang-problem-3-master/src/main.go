package main

import (
	"fmt"
	"regexp"
)

type Bale interface {
	AddUser(username string, isBot bool) (int, error)
	AddChat(chatname string, isGroup bool, creator int, admins []int) (int, error)
	SendMessage(userId, chatId int, text string) (int, error)
	SendLike(userId, messageId int) error
	GetNumberOfLikes(messageId int) (int, error)
	SetChatAdmin(chatId, userId int) error
	GetLastMessage(chatId int) (string, int, error)
	GetLastUserMessage(userId int) (string ,int, error)
}

var usersList []BaleImpl
var usernameChecker = regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString 
var chatList []BaleChat
var chatMessageList []BaleMessages
var likedMessages []BaleLikedMessages

type BaleImpl struct {
	username string
	id int
	baleType bool
}

type BaleChat struct {
	name string 
	chatType bool
	creatorId int 
	adminIds []int
	chatId int
}

type BaleMessages struct {
	senderId int
	receiverChatId int
	message string
	messageId int
}

func NewBaleImpl() *BaleImpl {
	user := new(BaleImpl)
	return user
}

type MessageError struct {
    message string
}

type BaleLikedMessages struct {
	userId int;
	messageId int;
}

func (e *MessageError) Error() string {
	return fmt.Sprintf(e.message)
}

func (bale BaleImpl) AddUser(username string, isBot bool) (int, error) {
	//fmt.Println("hello")
	var i int = 0
	if len(usersList) == 0 {
		if len(username) > 3 && usernameChecker(username) {
			user := new(BaleImpl)
			user.baleType = isBot
			user.username = username
			user.id = i
			usersList = append(usersList, *user)
		} else {
			return 0, &MessageError{message: "invalid username"}
		}
	} else {
		for i = 0; i < len(usersList); i++ {
			if username == usersList[i].username {
				return 0, &MessageError{message: "invalid username"}
			}
		}
		if len(username) > 3 && usernameChecker(username) {
			user := new(BaleImpl)
			user.baleType = isBot
			user.username = username
			user.id = i
			usersList = append(usersList, *user)
		} else {
			return 0, &MessageError{message: "invalid username"}
		}
	}
	
	return i, nil
}


func (bale BaleImpl) AddChat(chatname string, isGroup bool, creator int, admins []int) (int, error) {
	if usersList[creator].baleType == true {
		return 0, &MessageError{message: "could not create chat"}
	} else {
		chat := new(BaleChat)
		chat.creatorId = creator
		chat.chatType = isGroup
		chat.name = chatname
		chat.adminIds = admins
		chat.chatId = len(chatList) + 1
		chatList = append(chatList, *chat)
		return chat.chatId, nil
	}

}

func (bale BaleImpl) SendMessage(userId, chatId int, text string) (int, error) {
	isAdmin := false
	for i := 0; i < len(chatList[chatId - 1].adminIds); i++ {
		if userId == chatList[chatId - 1].adminIds[i] || chatList[chatId - 1].chatType == false {
			isAdmin = true
			break
		}
	}
	if isAdmin {
		baleMessage := new(BaleMessages)
		baleMessage.senderId = userId
		baleMessage.receiverChatId = chatId
		baleMessage.message = text
		baleMessage.messageId = len(chatMessageList) + 1
		chatMessageList = append(chatMessageList, *baleMessage)
		return baleMessage.messageId, nil
	} else {
		return 0, &MessageError{message: "user could not send message"}
	}
}

func (bale BaleImpl) SendLike(userId, messageId int) error {
	validMessage := false
	//check if message exist
	for i := 0; i <len(chatMessageList); i++ {
		if messageId == chatMessageList[i].messageId {
			validMessage = true
		}
	}

	if !validMessage {
		return &MessageError{message: "message not found"}
	}
	
	//check if this message liked before or not
	for i := 0; i < len(likedMessages); i++ {
		if likedMessages[i].messageId == messageId && likedMessages[i].userId == userId {
			return &MessageError{message: "this user has liked this message before"}
			break
		}
	}

	//succesfull process
	baleLikeMessage := new(BaleLikedMessages)
	baleLikeMessage.userId = userId
	baleLikeMessage.messageId = messageId
	likedMessages = append(likedMessages, *baleLikeMessage)
	return nil
}

func (bale BaleImpl) GetNumberOfLikes(messageId int) (int, error) {
	var i = 0
	var counter = 0
	for i = 0; i < len(likedMessages); i++ {
		if likedMessages[i].messageId == messageId {
			counter++
		}
	}

	return counter, nil
}

func (bale BaleImpl) SetChatAdmin(chatId, userId int) error {
	for i := 0; i < len(chatList); i++ {
		if chatId == chatList[i].chatId {
			for j := 0; j < len(chatList[i].adminIds); j++ {
				if userId == chatList[i].adminIds[j] {
					return &MessageError{message: "user is already admin"}
				}
			}
			chatList[i].adminIds = append(chatList[i].adminIds, userId)
			return nil
		}
	}

	return &MessageError{message: "This chat doesn't exist."}
}

func GetLastMessage(chatId int) (string, int, error) {
	for i := len(chatMessageList) - 1; i >= 0; i-- {
		if chatId == chatMessageList[i].receiverChatId {
			return chatMessageList[i].message, chatMessageList[i].messageId, nil
		}
	}

	return "", 0, &MessageError{message: "This chat doesn't exist."}
}

func GetLastUserMessage(userId int) (string ,int, error) {
	for i := len(chatMessageList) - 1; i >= 0; i-- {
		if userId == chatMessageList[i].senderId {
			return chatMessageList[i].message, chatMessageList[i].messageId, nil
		}
	}

	return "", 0, &MessageError{message: "This chat doesn't exist."}
}
