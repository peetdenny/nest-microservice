# Useful if you want to deploy on Heroku. Configures your env variables in true 12 Factor App fashion!
. .env
exec > /dev/null
heroku config:set PORT=$PORT --app $HEROKU_APP
heroku config:set DEVICE=$DEVICE --app $HEROKU_APP
heroku config:set BEARER=$BEARER --app $HEROKU_APP
