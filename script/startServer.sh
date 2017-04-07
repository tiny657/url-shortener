git clone https://github.com/tiny657/url-shortener.git
cd url-shortener
docker build -t url-shortener .
docker run -d --publish 8080:8080 --name url-shortener url-shortener
