## Book Struct

```
Book struct {
	UUID        string   
	Name        string   
	AuthorList  []string 
	PublishDate string   
	ISBN        string   
}
```

## API Endpoints

| Method   | Endpoint      | Description        | Auth Required | 
| -------- | ------------- | ------------------ | ------------- |
| `POST`   | `/get-token`  | Generate JWT token | Basic         |
| `POST`   | `/books`      | Create a new book  | Basic/Bearer  | 
| `GET`    | `/books`      | Get all books      | Basic/Bearer  |
| `GET`    | `/books/{id}` | Get book by UUID   | Basic/Bearer  |
| `PUT`    | `/books/{id}` | Update a book      | Basic/Bearer  |
| `DELETE` | `/books/{id}` | Delete a book      | Basic/Bearer  |


## CLI Flags

| Flag             | Shorthand | Type   | Default    | Description                                                                                                                       |
| ---------------- | --------- | ------ | ---------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `--port`         | `-p`      | string | `"3000"`   | The port number on which the server will listen. Example: `--port=8080`                                                           |
| `--secret`       | `-s`      | string | `"secret"` | The secret key used to sign JWT tokens. Example: `--secret=mySecretKey123`                                                        |
| `--authRequired` | *(none)*  | bool   | `true`     | Enable or disable JWT authentication middleware. Use `--authRequired=false` to disable authentication for testing or open access. |


## Running Locally

```bash
git clone https://github.com/al-masood/book_server.git
cd book_server
go build -o book_server
./book_server serve 
```

## Running with Docker

```bash
docker pull almasood/book-server:latest
docker run -d -p 3000:3000 --name book_server almasood/book-server:latest serve
```