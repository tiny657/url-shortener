# url-shortener

# How to start server

- git clone https://github.com/tiny657/url-shortener.git
- cd url-shortener
- ./startServer.sh

# How to test

## Create shorten URL
- ./script/createShortenUrl.sh
- ./script/createShortenUrl.sh [OriginalUrl]

## Get shorten URL
- ./script/getOriginalUrl.sh
- ./script/getOriginalUrl.sh [ShortenUrl]

# How to implement
- Used Redis key value store.  
- Keys are shortenUrl made by SHA1.
- Values are originalUrl
