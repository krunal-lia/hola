# Hello Translator!

## How to use

### Step 1: 
`git clone github.com/krunal-lia/hola`

### Step 2:
`go mod download`

### Step 3: 
#### If you have temporal pre installed 
`temporal server start-dev`
#### Alternatively use docker compose
```
git clone https://github.com/temporalio/docker-compose.git
cd  docker-compose
docker-compose up
```

### Step 4:
`go run worker/main.go`

### Step 4:
`go run worker/start.go`

Enter your name and the language you want to translate to!

### Supported languages
English, Spanish, French, German, Italian, Japanese, Russian, Chinese, Arabic, Hindi
