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
```
