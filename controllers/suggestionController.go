package controllers

import (
	"net/http"

	"github.com/Barbra-GbR/barbra-backend/helpers"
	"github.com/Barbra-GbR/barbra-backend/models"
	"github.com/gin-gonic/gin"
)

//Provides an webInterface for generating and managing suggestions
type SuggestionController struct{}

func (controller *SuggestionController) GetSuggestions(c *gin.Context) {
	if c.Query("ids") != "" {
		controller.GetSuggestionsById(c)
		return
	}

	//TODO
	dummyArticle1, _ := models.GetSuggestion(
		"https://medium.com/@saginadir/why-i-love-golang-90085898b4f7",
		"article",
		"Why I Love Golang",
		"Medium",
		"go",
		[]string{"programming"},
		"I love the Go programming language, or as some refer to it, Golang. It\u2019s simple and it\u2019s great.\r\n\r\nI write this on a tangent. Didn\u2019t expect Golang to be so good.\r\n\r\nI first picked up go around January 2016, it had a relative small but enthusiastic community here in Israel.\r\n\r\nI didn\u2019t think much of it at the time, I was honing my programming skills and Golang was just a tool I\u2019ve used to accomplish a task.\r\n\r\nEven one year ago, using go was brilliant. The process was straightforward once I\u2019ve got the general hang of the language.\r\n\r\nI wrote a crucial piece of code for Visualead, the company I work for, and it didn\u2019t let us down, still running in production a year later with zero maintenance since then.\r\n\r\nRecently I found myself again using Golang again, and I felt compelled to write about the reasons I fell in love with Golang.\r\n\r\nThe GOPATH environment\r\nThis is one of the first things you\u2019ll have to handle once you begin writing in Go.\r\n\r\nSetup your GOPATH directory anywhere on your computer, complete with bin, src, and pkg directories and you are ready to begin writing.\r\n\r\n")

	dummyArticle2, _ := models.GetSuggestion(
		"https://medium.com/@saginadir/why-i-love-golang-90085898b4f7",
		"article",
		"Why I Love Golang 2",
		"Medium",
		"go",
		[]string{"programming"},
		"I love the Go programming language, or as some refer to it, Golang. It\u2019s simple and it\u2019s great.\r\n\r\nI write this on a tangent. Didn\u2019t expect Golang to be so good.\r\n\r\nI first picked up go around January 2016, it had a relative small but enthusiastic community here in Israel.\r\n\r\nI didn\u2019t think much of it at the time, I was honing my programming skills and Golang was just a tool I\u2019ve used to accomplish a task.\r\n\r\nEven one year ago, using go was brilliant. The process was straightforward once I\u2019ve got the general hang of the language.\r\n\r\nI wrote a crucial piece of code for Visualead, the company I work for, and it didn\u2019t let us down, still running in production a year later with zero maintenance since then.\r\n\r\nRecently I found myself again using Golang again, and I felt compelled to write about the reasons I fell in love with Golang.\r\n\r\nThe GOPATH environment\r\nThis is one of the first things you\u2019ll have to handle once you begin writing in Go.\r\n\r\nSetup your GOPATH directory anywhere on your computer, complete with bin, src, and pkg directories and you are ready to begin writing.\r\n\r\n")

	dummyArticle3, _ := models.GetSuggestion(
		"https://medium.com/@saginadir/why-i-love-golang-90085898b4f7",
		"article",
		"Why I Love Golang 3",
		"Medium",
		"go",
		[]string{"programming"},
		"I love the Go programming language, or as some refer to it, Golang. It\u2019s simple and it\u2019s great.\r\n\r\nI write this on a tangent. Didn\u2019t expect Golang to be so good.\r\n\r\nI first picked up go around January 2016, it had a relative small but enthusiastic community here in Israel.\r\n\r\nI didn\u2019t think much of it at the time, I was honing my programming skills and Golang was just a tool I\u2019ve used to accomplish a task.\r\n\r\nEven one year ago, using go was brilliant. The process was straightforward once I\u2019ve got the general hang of the language.\r\n\r\nI wrote a crucial piece of code for Visualead, the company I work for, and it didn\u2019t let us down, still running in production a year later with zero maintenance since then.\r\n\r\nRecently I found myself again using Golang again, and I felt compelled to write about the reasons I fell in love with Golang.\r\n\r\nThe GOPATH environment\r\nThis is one of the first things you\u2019ll have to handle once you begin writing in Go.\r\n\r\nSetup your GOPATH directory anywhere on your computer, complete with bin, src, and pkg directories and you are ready to begin writing.\r\n\r\n")

	c.JSON(http.StatusOK, []*models.Suggestion{dummyArticle1, dummyArticle2, dummyArticle3})
}

//GetSuggestion returns the suggestion with the specified id
func (controller *SuggestionController) GetSuggestion(c *gin.Context) {
	id, err := helpers.StringToObjectId(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	suggestion, err := models.GetSuggestionById(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, suggestion)
}

//Returns all suggestions with matching ids set in the ids query string
func (controller *SuggestionController) GetSuggestionsById(c *gin.Context) {
	ids, err := helpers.StringToObjectIds(c.Query("ids"))
	if err != nil {
		//TODO change
		c.AbortWithStatus(http.StatusTeapot)
		return
	}

	suggestions, err := models.GetSuggestionsById(ids)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, suggestions)
}
