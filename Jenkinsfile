pipeline {
    agent any
    tools {
        go 'go'
        dockerTool 'docker'
    }
    stages {
        stage ('Install dependencies') {
            steps {
                sh 'go get github.com/gorilla/mux'
                sh 'go get github.com/go-sql-driver/mysql'
                sh 'go get github.com/joho/godotenv'
            }
        }
        stage ('Git') {
            steps {
                git url: 'https://github.com/emorales8/Info-service'
            }
        }
        stage ('Go App Build') {
            steps {
                sh 'go build get.go'
            }
        }
        stage ('Docker build') {
            steps {
                sh 'docker build -t go-get-srv:v1.0 .'
            }
        }
    }
}
