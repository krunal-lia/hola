# Hello Translator!

## How to use

### Step 1: 
`git clone git@github.com:krunal-lia/hola.git`

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
```
cd ..
go run worker/main.go
```

### Step 4:
`go run start/main.go`

Enter your name and the language you want to translate to!

### Supported languages
English, Spanish, French, German, Italian, Japanese, Russian, Chinese, Arabic, Hindi


#### Sample output
![Screenshot 2024-04-23 at 8 50 51 PM](https://github.com/krunal-lia/hola/assets/29747181/a7e74a1c-45f8-40be-a0af-af00c713ff14)
