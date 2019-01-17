# little-bird

LittleBird is a little blog for sharing technical stuffs...

## Demo

New version of Little Bird with VueJS is available at: https://goway.herokuapp.com/

##  Build

This project used [glide](https://glide.sh/) for dependency management hence please make sure you have installed it before building the project.

```
go get github.com/golovers/little-bird

cd $GOPATH/src/github.com/golovers/little-bird
glide install
```

## Deploy

### Variables

LittleBird uses environment variables for configuration. Hence please make sure following variables are set accordingly.

```
# Set this to ingore authentication for local testing, if this is set to true you can ignore the Google OAuth2 variables
LITTLE_BIRD_IGNORE_AUTH=true

# Google OAuth2, more details at https://developers.google.com/identity/sign-in/web/sign-in
OAUTH2_CLIENT_ID=xxxx.apps.googleusercontent.com
OAUTH2_CLIENT_SECRET=xxxxxxxxxxxxxxxx
OAUTH2_CALLBACK=http://localhost:8080/oauth2callback

# Article database
ARTICLE_DB_URI=mongodb://127.0.0.1:27017
ARTICLE_DB_NAME=littlebird
ARTICLE_DB_USERNAME=user
ARTICLE_DB_PASSWORD=pass

# Comment database
COMMENT_DB_URI=mongodb://127.0.0.1:27017
COMMENT_DB_NAME=littlebird
COMMENT_DB_USERNAME=user
COMMENT_DB_PASSWORD=pass

# Vote database
VOTE_DB_URI=mongodb://127.0.0.1:27017
VOTE_DB_NAME=littlebird
VOTE_DB_USERNAME=user
VOTE_DB_PASSWORD=pass
```

### Local Deployment
Make sure mongodb started: `sudo service mongod start`

Put the variables into a `local.env` file.

Run `make run-local`

Open the site at: http://localhost:8080/

### Deploy to heroku
You can use online mongodb servive for storing data: https://mlab.com/

Put appropriate variables into `heroku.env` file.

Make sure you have an heroku account & heroku-cli is installed.

Run following commands:

```
cd $GOPATH/src/github.com/golovers/little-bird
heroku create little-bird
make heroku-config
git push heroku master
```

## License

This project is licensed under MIT.

