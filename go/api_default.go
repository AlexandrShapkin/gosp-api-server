/*
 * Gosp API Specification
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/gin-gonic/gin"
)

type DefaultAPI interface {


    // CreateChat Post /chats
     CreateChat(c *gin.Context)

    // CreateMessage Post /chats/:chatId/messages
     CreateMessage(c *gin.Context)

    // DeleteChat Delete /chats/:chatId
     DeleteChat(c *gin.Context)

    // GetChat Get /chats/:chatId
     GetChat(c *gin.Context)

    // GetChats Get /chats
     GetChats(c *gin.Context)

    // GetMe Get /users/me
     GetMe(c *gin.Context)

    // GetMessage Get /chats/:chatId/messages/:messageId
     GetMessage(c *gin.Context)

    // GetMessages Get /chats/:chatId/messages
     GetMessages(c *gin.Context)

    // GetUsersList Get /users
     GetUsersList(c *gin.Context)

    // LoginUser Post /auth/login
     LoginUser(c *gin.Context)

    // RegisterUser Post /auth/register
     RegisterUser(c *gin.Context)

    // UpdateChat Put /chats/:chatId
     UpdateChat(c *gin.Context)

    // UpdateMessage Put /chats/:chatId/messages/:messageId
     UpdateMessage(c *gin.Context)

}