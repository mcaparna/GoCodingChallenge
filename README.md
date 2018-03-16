# Golang Coding Challenge

Hi! Welcome to the Golang coding challenge. Below is a set of instructions that must attempt to complete within 3 days. Fork the repo when you're ready and good luck! 😀

Within this repository, you will find a hastily thrown together application. It's a very basic, a simple To-Do API with the ability to create and list your to-dos.

Here are some things we need help with.

1. We have create and list features but lack the ability to update. Introduce a new PUT endpoint at `/todos/{todoID}` that receives a JSON body containing title and status. The feature should update the existing record and return a JSON body representing the _new_ state of the todo item.  
You may notice the lack of tests in the repo, maybe set a good example and add tests to your method if you have time. That way the other devs can copy-paste from your good example.  
Once done, go ahead and open a pull request again the repo.

2. As mentioned before, we have create and list already in place. The dev team was super excited because they knocked this out faster than anyone thought, maybe too fast.  
Feel free to open a github issue and point out some of those shortcuts. If you're feeling bold, Pull Requests are always welcome 😀.

# Setup

Within the repo you will find a docker-compose.yaml file. If you're familar with docker and docker-compose, great! You can get started by simply running `docker-compose up` and that will create an API and Postgres container for you.

If you are not familar with those tools, feel free to setup whatever environment you are comfortable working within. At the very least you will need a go environment to run your API and a SQL database. There is a todo_schema.sql file that will create a basic table and sequence to get you started. To run the API, simply run `go run main.go`. You will need a few variables within the app, so please make sure those are provided.

Environment Variables

* DB_USER
* DB_PASSWORD
* DB_NAME

Also, regardless of which environment you use, you'll need to install dependencies. For this repo, we're using [Dep](https://golang.github.io/dep/). A Gopkg toml and lock file has been provided, simply run `dep ensure`.

--------
Over the last three days, I have had a great time working on this GoLang coding, took it up as a challenge as it was a new language and technology for me, just like the way things work in real-life in any organization or projects.

Here's what I have been able to accomplish in the time frame along with my other commitments. 

1. Setup the Docker, GoLang and PostGres.  
2. Understand the working of the docker compose package.
3. Understand the GoLang language and executing the basic given package with all the dependencies.
Had to spend some time in understanding the GOPATH setup in windows environment.
4. Enhancing to add the Update Feature.
5. Adding testcases and trying to execute it (phew)(handlers_test.go)
6. Trying to mock the PostGres DB calls (WIP)
 
Testcases�phew I could only put a basic framework but the postgres mocking is still posing a challenge to me

Here is the link to my github 
 
https://github.com/mcaparna/GoCodingChallenge.git

Having said that, I can already think of a lot of enhancements -
1. Like having a UI instead of using POSTMAN. 
2. A lot more validations on every input and output and the design itself can be changed to make it more modular like having Controllers etc to make adding test cases easier. 
3. Look into more go packages to include more mocking.  


