DEFAULT="http://a.very.long.url"
URL={\"url\":\"${1:-$DEFAULT}\"}

curl -sX POST -H 'Content-Type: application/json' 'localhost:8080/shorten' -d $URL
