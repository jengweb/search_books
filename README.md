# search_books
Stack RESTful API
```
Golang + Elasticsearch

Start Docker by commandline
docker-compose up --build -d
```

Ceate Book
```
Method: POST
URL: localhost/insert
Post form: data =
{
  "title": "Elasticsearch: The Definitive Guide",
  "author_name": [
    "Clinton Gormley",
    "Zachary Tong"
  ],
  "tag": [
    "search",
    "computer"
  ],
  "isbn": "1449358543",
  "price": 44.3,
  "total_page": 724,
  "description": "A Distributed Real-Time Search and Analytics Engine"
}

command line test
curl -X POST \
  http://goelasticsearch.test/insert \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Postman-Token: 8802fe5b-32cf-4e66-a10f-5f43ede54d5c' \
  -H 'cache-control: no-cache' \
  -d 'data=%7B%0A%20%20%22title%22%3A%20%22Elasticsearch%3A%20The%20Definitive%20Guide%22%2C%0A%20%20%22author_name%22%3A%20%5B%0A%20%20%20%20%22Clinton%20Gormley%22%2C%0A%20%20%20%20%22Zachary%20Tong%22%0A%20%20%5D%2C%0A%20%20%22tag%22%3A%20%5B%0A%20%20%20%20%22search%22%2C%0A%20%20%20%20%22computer%22%0A%20%20%5D%2C%0A%20%20%22isbn%22%3A%20%221449358543%22%2C%0A%20%20%22price%22%3A%2044.3%2C%0A%20%20%22total_page%22%3A%20724%2C%0A%20%20%22description%22%3A%20%22A%20Distributed%20Real-Time%20Search%20and%20Analytics%20Engine%22%2C%0A%20%20%22detail%22%3A%20%22Learn%20how%20to%20use%20Elasticsearch%2C%20an%20open%20source%2C%20distributed%2C%20RESTful%20search%20engine%20built%20on%20top%20of%20Apache%20Lucene.%20Each%20chapter%20in%20this%20book%20tackles%20a%20particular%20facet%20of%20Elasticsearch%20with%20separate%20sections%20for%20beginners%20more%20advanced%20programmers.%20If%20you%E2%80%99re%20a%20beginner%2C%20advanced%20techniques%20are%20not%20required%20reading%2C%20but%20you%20can%20revisit%20them%20once%20you%20have%20a%20solid%20understanding%20of%20the%20basics.%22%0A%7D&undefined='
```

Search Book
```
Method: GET
URL: localhost/search?query=Elasticsearch
```

Update Book
```
Method: POST
URL: localhost/update
Post form: 
_id = [data id from method Get filed "_id"]
data =
{
  "title": "Elasticsearch",
  "author_name": [
    "Clinton Gormley",
    "Zachary Tong"
  ],
  "tag": [
    "search",
    "computer"
  ],
  "isbn": "1449358543",
  "price": 44.3,
  "total_page": 724,
  "description": "A Distributed Real-Time Search and Analytics Engine"
}

command line test
curl -X POST \
  http://goelasticsearch.test/update \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Postman-Token: e343adc9-153e-435c-b5c4-b764c076c6b7' \
  -H 'cache-control: no-cache' \
  -d 'data=%7B%0A%20%20%22title%22%3A%20%22Elasticsearch%22%2C%0A%20%20%22author_name%22%3A%20%5B%0A%20%20%20%20%22Clinton%20Gormley%22%2C%0A%20%20%20%20%22Zachary%20Tong%22%0A%20%20%5D%2C%0A%20%20%22tag%22%3A%20%5B%0A%20%20%20%20%22search%22%2C%0A%20%20%20%20%22computer%22%0A%20%20%5D%2C%0A%20%20%22isbn%22%3A%20%221449358543%22%2C%0A%20%20%22price%22%3A%2044.3%2C%0A%20%20%22total_page%22%3A%20999%2C%0A%20%20%22description%22%3A%20%22A%20Distributed%20Real-Time%20Search%20and%20Analytics%20Engine%22%2C%0A%20%20%22detail%22%3A%20%22Learn%20how%20to%20use%20Elasticsearch%2C%20an%20open%20source%2C%20distributed%2C%20RESTful%20search%20engine%20built%20on%20top%20of%20Apache%20Lucene.%20Each%20chapter%20in%20this%20book%20tackles%20a%20particular%20facet%20of%20Elasticsearch%20with%20separate%20sections%20for%20beginners%20more%20advanced%20programmers.%20If%20you%E2%80%99re%20a%20beginner%2C%20advanced%20techniques%20are%20not%20required%20reading%2C%20but%20you%20can%20revisit%20them%20once%20you%20have%20a%20solid%20understanding%20of%20the%20basics.%22%0A%7D&_id=KbqilnAB1DYbAWo8f7zJ&undefined='
```
