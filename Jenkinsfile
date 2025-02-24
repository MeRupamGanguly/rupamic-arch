pipeline {
    agent any

    environment {
        GO_VERSION = '1.24'  // Specify the Golang version you are using
        GO_PATH = '/go'      // Set Go workspace path
        REPO_URL = 'https://github.com/MeRupamGanguly/rupamic-arch.git'
        BRANCH_NAME = 'jenk' // Change this to the specific branch
    }

    stages {
        stage('Checkout') {
            
            steps {
                // Checkout the code from GitHub repository
                git branch: "${BRANCH_NAME}", url: "${REPO_URL}"
            }
        }

        stage('Install Dependencies') {
            
            steps {
                script {
                    // Install Go dependencies (assumes go.mod and go.sum are present)
                    sh "go mod tidy"
                }
            }
        }

        stage('Install GolangCI-Lint') {
            
            steps {
                script {
                    // Install golangci-lint if it is not installed
                    sh '''
                    if ! command -v golangci-lint &> /dev/null
                    then
                        echo "golangci-lint not found, installing..."
                        curl -sSfL https://github.com/golangci/golangci-lint/releases/download/v1.64.5/golangci-lint-1.64.5-linux-amd64.tar.gz | tar -xz -C /usr/local/bin
                    else
                        echo "golangci-lint is already installed"
                    fi
                    '''
                }
            }
        }

        stage('Lint and Static Analysis') {
            
            steps {
                script {
                    // Run Go linters and static checks using golangci-lint
                    sh 'golangci-lint run'  // Run all linters configured in .golangci.yml
                }
            }
        }

        stage('Run Tests') {
           
            steps {
                script {
                    // Run Golang tests
                    // sh 'go test -v ./...'
                    sh 'go test -v $(go list ./... | grep -v '/integration')'
                }
            }
        }

        stage('Build Binaries') {
           
            steps {
                script {
                    // Set GOOS and GOARCH to build for the desired platform
                    sh 'GOOS=linux GOARCH=amd64 go build -o userservice user/cmd/user.go' // Adjust this if your entrypoint is different
                }
            }
        }

        stage('Docker Build') {
           
            steps {
                script {
                    // Create Docker image using the minimal base image
                    sh '''
                    # docker build -t userservice:latest -f Dockerfile.user .
                    echo "Create Docker image using the minimal base image"
                    '''
                }
            }
        }

        stage('Run Docker Container') {
           
            steps {
                script {
                    // Run the Docker container
                    sh '''
                    # docker run -d --name userservice-container -p 5002:5002 userservice:latest
                    echo "Run the Docker container"
                    '''
                }
            }
        }

        stage('Cleanup') {
           
            steps {
                script {
                    // Clean up unnecessary files after build
                    sh 'go clean -cache -modcache'
                }
            }
        }
    }

    post {
        always {
            // Clean up Docker containers and images after the pipeline
            sh '''
            # docker rm -f userservice-container || true
            # docker rmi -f userservice:latest || true
            echo "Post Stages Runing... Nothing ToDo"
            '''
        }
    }
}