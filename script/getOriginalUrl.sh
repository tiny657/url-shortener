DEFAULT="http://localhost/ea7c045ff32d2295067299b0f6b214f000114ccb"
URL={\"short\":\"${1:-$DEFAULT}\"}

curl -sX GET -H 'Content-Type: application/json' 'localhost:8080/original' -d $URL
