package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/bellwood4486/sandbox-go/restapi-tweet/utils"

	"github.com/bellwood4486/sandbox-go/restapi-tweet/models"
)

func (c Controller) GetTweets(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tweet models.Tweet
		tweets := make([]models.Tweet, 0)
		var errorObj models.Error
		rows, err := db.Query("select * from tweets")
		if err != nil {
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&tweet.ID, &tweet.Content, &tweet.UserName, &tweet.CommentNum, &tweet.StarNum, &tweet.ReTweetNum)
			if err != nil {
				log.Println(err)
				errorObj.Message = "Server error"
				utils.Respond(w, http.StatusInternalServerError, errorObj)
				return
			}
			tweets = append(tweets, tweet)
		}
		if err := rows.Err(); err != nil {
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		utils.Respond(w, http.StatusOK, tweets)
	}
}

func (c Controller) GetTweet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tweet models.Tweet
		var errorObj models.Error
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			errorObj.Message = `"id" is wrong`
			utils.Respond(w, http.StatusBadRequest, errorObj)
			return
		}
		rows := db.QueryRow("select * from tweets where id=$1", id)
		err = rows.Scan(&tweet.ID, &tweet.Content, &tweet.UserName, &tweet.CommentNum, &tweet.StarNum, &tweet.ReTweetNum)
		if err != nil {
			if err == sql.ErrNoRows {
				errorObj.Message = "The tweet does not exist"
				utils.Respond(w, http.StatusBadRequest, errorObj)
				return
			}
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		utils.Respond(w, http.StatusOK, tweet)
	}
}
func (c Controller) AddTweet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tweet models.Tweet
		var errorObj models.Error
		if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
			panic(err)
		}
		if tweet.Content == "" {
			errorObj.Message = "\"content\" is missing"
			utils.Respond(w, http.StatusBadRequest, errorObj)
			return
		}
		if tweet.UserName == "" {
			errorObj.Message = "\"user_name\" is missing"
			utils.Respond(w, http.StatusBadRequest, errorObj)
			return
		}
		if tweet.CommentNum < 0 {
			errorObj.Message = "\"comment_num\" must be greater than or equal to 0"
			utils.Respond(w, http.StatusBadRequest, errorObj)
			return
		}
		if tweet.StarNum < 0 {
			errorObj.Message = "\"star_num\" must be greater than or equal to 0"
			utils.Respond(w, http.StatusBadRequest, errorObj)
			return
		}
		if tweet.ReTweetNum < 0 {
			errorObj.Message = "\"re_tweet_num\" must be greater than or equal to 0"
			utils.Respond(w, http.StatusBadRequest, errorObj)
			return
		}
		err := db.QueryRow(
			"insert into tweets (content, user_name, comment_num, star_num, re_tweet_num)"+
				" values($1, $2, $3, $4, $5) RETURNING id;",
			tweet.Content, tweet.UserName, tweet.CommentNum, tweet.StarNum, tweet.ReTweetNum).Scan(&tweet.ID)
		if err != nil {
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		utils.Respond(w, http.StatusCreated, tweet)
	}
}

func (c Controller) PutTweet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tweet models.Tweet
		var errorObj models.Error
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			errorObj.Message = "\"id\" is wrong"
			utils.Respond(w, http.StatusBadRequest, errorObj)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
			panic(err)
		}
		if tweet.CommentNum < 0 {
			errorObj.Message = "\"comment_num\" must be greater than or equal to 0"
			utils.Respond(w, http.StatusBadRequest, errorObj)
			return
		}
		if tweet.StarNum < 0 {
			errorObj.Message = "\"star_num\" must be greater than or equal to 0"
			utils.Respond(w, http.StatusBadRequest, errorObj)
			return
		}
		if tweet.ReTweetNum < 0 {
			errorObj.Message = "\"re_tweet_num\" must be greater than or equal to 0"
			utils.Respond(w, http.StatusBadRequest, errorObj)
			return
		}
		var update models.Tweet
		rows := db.QueryRow("select * from tweets where id=$1", id)
		err = rows.Scan(&update.ID, &update.Content, &update.UserName, &update.CommentNum,
			&update.StarNum, &update.ReTweetNum)
		if err != nil {
			if err == sql.ErrNoRows {
				errorObj.Message = "The tweet does not exist"
				utils.Respond(w, http.StatusBadRequest, errorObj)
				return
			}
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		if tweet.Content != "" {
			update.Content = tweet.Content
		}
		if tweet.UserName != "" {
			update.UserName = tweet.UserName
		}
		if tweet.CommentNum > 0 {
			update.CommentNum = tweet.CommentNum
		}
		if tweet.StarNum > 0 {
			update.StarNum = tweet.StarNum
		}
		if tweet.ReTweetNum > 0 {
			update.ReTweetNum = tweet.ReTweetNum
		}
		result, err := db.Exec(
			"update tweets set content=$1, user_name=$2, comment_num=$3, star_num=$4, re_tweet_num=$5"+
				" where id=$6 RETURNING id",
			&update.Content, &update.UserName, &update.CommentNum, &update.StarNum, &update.ReTweetNum, id)
		if err != nil {
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		_, err = result.RowsAffected()
		if err != nil {
			if err == sql.ErrNoRows {
				errorObj.Message = "The tweet does not exist"
				utils.Respond(w, http.StatusBadRequest, errorObj)
				return
			}
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		utils.Respond(w, http.StatusOK, update)
	}
}

func (c Controller) RemoveTweet(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errorObj models.Error
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			errorObj.Message = "\"id\" is wrong"
			utils.Respond(w, http.StatusBadRequest, errorObj)
			return
		}
		result, err := db.Exec("delete from tweets where id = $1", id)
		if err != nil {
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		_, err = result.RowsAffected()
		if err != nil {
			if err == sql.ErrNoRows {
				errorObj.Message = "The tweet does not exist"
				utils.Respond(w, http.StatusBadRequest, errorObj)
				return
			}
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		utils.Respond(w, http.StatusNoContent, "")
	}
}
