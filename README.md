## Dependencies

Make sure you have these components installed
- Golang [https://golang.org/](More info)
- Buffalo [https://gobuffalo.io/en](More info)
- Postgressql [https://www.postgresql.org/](More info)
- Nodejs [https://nodejs.org/en/](More info)

### Create Your Databases

Run below command to create the database

	$ buffalo db create -a
## Starting the Application in dev mode

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

App can be accessed in [http://127.0.0.1:3000](http://127.0.0.1:3000) locally
App is also deployed in [https://blooming-sierra-32775.herokuapp.com/](Heroku)