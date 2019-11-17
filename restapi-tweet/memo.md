```shell script
$ docker run --rm --name my_postgres -p 5432:5432 -e POSTGRES_USER=puser -e POSTGRES_PASSWORD=ppassword -e POSTGRES_DB=testdb -v $(pwd)/initdb:/docker-entrypoint-initdb.d -d postgres:9.6
$ curl -v -H "Accept:application/json" -H "Content-Type:application/json" -X POST -d '{"content":"This is my first tweet.","user_name":"@keidrun","comment_num":5,"star_num":15,"re_tweet_num":25}' http://localhost:3000/api/tweets | jq
$ curl -v -H "Accept:application/json" -H "Content-Type:application/json" -X POST -d '{"content":"Golang is my favorite language!","user_name":"@superdeveloper","comment_num":22,"star_num":222,"re_tweet_num":2222}' http://localhost:3000/api/tweets | jq
$ curl -v -H "Accept:application/json" -H "Content-Type:application/json" -X POST -d '{"content":"I am nothing. Just an ordinary guy.","user_name":"@person"}' http://localhost:3000/api/tweets | jq
$ curl -v -H "Accept:application/json" http://localhost:3000/api/tweets/1 | jq
$ curl -v -H "Accept:application/json" -H "Content-Type:application/json" -X PUT -d '{"content":"I am excellent guy!!","user_name":"@awesomeperson","comment_num":99,"star_num":999,"re_tweet_num":9999}' http://localhost:3000/api/tweets/3 | jq
$ curl -v -H "Accept:application/json" -X DELETE http://localhost:3000/api/tweets/2 | jq
$ curl -v -H "Accept:application/json" http://localhost:3000/api/tweets | jq
```