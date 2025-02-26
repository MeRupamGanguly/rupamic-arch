pipeline {
    agent any

    environment {
        GO_VERSION = '1.24.0'  // Specify the Golang version you are using
        GO_PATH = '/go'        // Set Go workspace path
        WORKSPACE_DIR = 'user_repo'
        GO_BIN_PATH = '/var/jenkins_home/go/bin'  // Set the explicit Go binary path
        LINT_BIN_PATH = '/var/jenkins_home/bin'  // Directory where golangci-lint will be installed
        PATH = "${GO_BIN_PATH}:${LINT_BIN_PATH}:${PATH}"  // Add Go and golangci-lint binary paths to the PATH globally
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    if (fileExists(WORKSPACE_DIR)) {
                        deleteDir() // This will delete the entire workspace if it exists
                    }
                    sh 'git --version'
                    // Checkout the code from GitHub repository
                    git branch: "jenk", url: "https://github.com/MeRupamGanguly/rupamic-arch.git"
                }
            }
        }

        stage('Install Golang') {
            steps {
                script {
                    // Install Go if it's not already installed or update to the latest version
                    sh '''
                    if ! command -v go &> /dev/null
                    then
                        echo "Go not found, installing Go version ${GO_VERSION}..."
                        curl -LO https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
                        tar -C $HOME -xvzf go${GO_VERSION}.linux-amd64.tar.gz  # Install in the Jenkins home directory
                        export PATH=$HOME/go/bin:$PATH
                        echo "Go ${GO_VERSION} installed successfully"
                    else
                        echo "Go is already installed"
                    fi
                    '''
                }
            }
        }

        stage('Install Dependencies') {
            steps {
                script {
                    // Ensure Go binary is in the PATH for this step
                    sh "go mod tidy"
                }
            }
        }

      stage('Install GolangCI-Lint') {
    steps {
        script {
            // Install golangci-lint if it's not installed
            sh '''
            if ! command -v golangci-lint &> /dev/null
            then
                echo "golangci-lint not found, installing..."
                mkdir -p /var/jenkins_home/bin  # Create a directory for golangci-lint in the Jenkins user's home
                curl -sSfL https://github.com/golangci/golangci-lint/releases/download/v1.64.5/golangci-lint-1.64.5-linux-amd64.tar.gz | tar -xz -C /var/jenkins_home/bin
                echo "golangci-lint installed successfully"
                # Move golangci-lint to the correct location
                mv /var/jenkins_home/bin/golangci-lint-1.64.5-linux-amd64/golangci-lint /var/jenkins_home/bin/golangci-lint
            else
                echo "golangci-lint is already installed"
            fi
            '''
            // List the contents of /var/jenkins_home/bin to verify that the binary is extracted correctly
            echo "Listing files in /var/jenkins_home/bin:"
            sh 'ls -l /var/jenkins_home/bin'
        }
    }
}

stage('Lint and Static Analysis') {
    steps {
        script {
            // Re-export PATH to include golangci-lint directory in this shell environment
            sh '''
            export PATH=$PATH:/var/jenkins_home/bin
            echo "Checking golangci-lint version"
            
            # Navigate to the directory containing Go code (adjust the path to where your Go code is located)
            cd user  # Replace 'user' with your actual Go project directory if needed
            
            # Run golangci-lint within the directory where the Go files are located
            golangci-lint run  // Run linters directly
            '''
        }
    }
}
        stage('Run Tests') {
            steps {
                script {
                    // Run Golang tests, skipping integration tests if desired
                    sh 'go test -v $(go list ./... | grep -v \'/integration\')'
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
            echo "Post Stages Running... Nothing To Do"
            '''
            deleteDir() // Delete the entire workspace directory
        }
    }
}
