pipeline {
    agent any // This specifies that Jenkins can run the pipeline on any available agent (machine or node).
    // agent {
    //     docker {
    //         image 'golang:1.20'  // Runs the pipeline inside a Docker container using the golang:1.20 image
    //     }
    // }
    // agent {
    //     node {
    //         name 'build-server'  // Runs only on the 'build-server' node
    //     }
    // }
    // agent {
    //     label 'linux-node'  // The pipeline will run on an agent with the label 'linux-node'
    // }
    environment {
        GO_VERSION = '1.24.0'  // Specify the Golang version you are using
        GO_PATH = '/go'        // Set Go workspace path
        WORKSPACE_DIR = 'user_repo'
        // Specify locations where Go binaries and the golangci-lint binary will be installed.
        GO_BIN_PATH = '/var/jenkins_home/go/bin'  // Set the explicit Go binary path
        LINT_BIN_PATH = '/var/jenkins_home/bin'  // Directory where golangci-lint will be installed
        // This updates the PATH environment variable to include both the Go binary and the linter's binary location globally in the pipeline steps.
        PATH = "${GO_BIN_PATH}:${LINT_BIN_PATH}:${PATH}"  // Add Go and golangci-lint binary paths to the PATH globally
    }

    stages {
        stage('Checkout') {
            steps { // This is where you define the actions to be performed in the pipeline stage. 
                script { // The script block is used to allow you to execute a series of shell commands.
                    // Checks if the specified directory (user_repo) exists. If it does, it deletes the entire workspace directory using deleteDir().
                    if (fileExists(WORKSPACE_DIR)) {
                        deleteDir() // This will delete the entire workspace if it exists
                    }
                    sh 'git --version' // Verifies that Git is installed and prints its version.
                    // Checks out the code from the GitHub repository on the jenk branch.
                    git branch: "jenk", url: "https://github.com/MeRupamGanguly/rupamic-arch.git"
                }
            }
        }

        stage('Install Golang') {
            steps {
                script {
                    // Install Go if it's not already installed or update to the latest version
                    // sh step in Jenkins is used to execute shell commands. This particular multi-line shell script starts with 
                    // sh ''' and ends with ''', which allows you to run several lines of shell code in one go.
                    sh '''
                    if ! command -v go &> /dev/null
                    then
                        echo "Go not found, installing Go version ${GO_VERSION}..."
                        curl -LO https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
                        tar -C $HOME -xzf go${GO_VERSION}.linux-amd64.tar.gz 
                        export PATH=$HOME/go/bin:$PATH
                        echo "Go ${GO_VERSION} installed successfully"
                    else
                        echo "Go is already installed"
                    fi
                    '''
                }
            }
            // command -v go This command checks if the go command exists in the current shell environment. It looks for the Go 
            // executable in the directories listed in the $PATH environment variable.
            // ! operator negates the command. So, if go is not found, the condition will evaluate as true.
            // &> /dev/null: This redirects both the standard output (stdout) and standard error (stderr) to /dev/null, which 
            // means any output from command -v go is discarded. This prevents unnecessary output from being displayed in the Jenkins logs.
            // If the go command is not found, the script proceeds with the block of code following the then statement. If go is 
            // already installed, the else block will execute instead.
            // curl is a command-line tool used to transfer data from or to a server.
            // -L: Follows redirects if the URL changes (some URLs may redirect to another location).
            // -O: Saves the file with the same name as the remote file. 
            // tar: This command is used to extract files from compressed archives.
            // -C $HOME: This option specifies the directory to extract the files to. $HOME is the environment variable that points to the user's home directory.
            // -xzf: These are options for tar:
            // -x: Extract files from the archive.
            // -z: Decompress the archive using gzip.
            // -f: Specifies the file to extract (in this case, go${GO_VERSION}.linux-amd64.tar.gz).
            // This command extracts the downloaded Go tarball into the $HOME directory, which will create a go directory under the user's home directory (e.g., /home/jenkins/go).
            // export PATH=$HOME/go/bin:$PATH    This command modifies the system's PATH environment variable to include the Go 
            // binary directory ($HOME/go/bin). This ensures that the Go binaries can be accessed from anywhere on the system, 
            // allowing you to run Go commands like go directly from the command line.  
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
                        mkdir -p /var/jenkins_home/bin 
                        curl -sSfL https://github.com/golangci/golangci-lint/releases/download/v1.64.5/golangci-lint-1.64.5-linux-amd64.tar.gz | tar -xz -C /var/jenkins_home/bin
                        echo "golangci-lint installed successfully"
                        # Move golangci-lint to the correct location
                        mv /var/jenkins_home/bin/golangci-lint-1.64.5-linux-amd64/golangci-lint /var/jenkins_home/bin/golangci-lint
                        export PATH=$PATH:/var/jenkins_home/bin
                    else
                        echo "golangci-lint is already installed"
                    fi
                    '''
                    // List the contents of /var/jenkins_home/bin to verify that the binary is extracted correctly
                    echo "Listing files in /var/jenkins_home/bin:"
                    sh 'ls -l /var/jenkins_home/bin'
                }
                // curl
                // -s: Silent mode, which suppresses progress output.
                // -S: Show errors if -s is used.
                // -f: Fail silently on server errors (e.g., 404).
                // -L: Follow redirects (if the URL is redirected to another location).
                // This mv step is needed because the extracted folder will have the binary in a nested directory (e.g., golangci-lint-1.64.5-linux-amd64/golangci-lint), but we want the binary to be directly available at /var/jenkins_home/bin/golangci-lint.
            }
        }

        stage('Lint and Static Analysis') {     
            steps {         
                script {             
                    // Check if any Go files exist in the user directory
                    sh '''
                    echo "Listing files in user directory"
                    ls -l user
                    
                    if find user -name "*.go" | grep -q .; then
                        cd user
                        echo "Go files found, running golangci-lint"
                        golangci-lint run --disable-all --enable staticcheck ./...
                    else
                        echo "No Go files found in the 'user' directory or subdirectories"
                        exit 1
                    fi
                    '''         
                }     
            } 
        }

        stage('Run Tests') {
            steps {
                script {
                    sh 'go test -v ./common/...'
                    // Run Golang tests, skipping integration tests if desired
                    // sh 'go test -v $(go list ./... | grep -v \'/integration\')'
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
