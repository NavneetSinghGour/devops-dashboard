pipeline {
    agent any

    options {
        ansiColor('xterm')
        timestamps()
        buildDiscarder(logRotator(numToKeepStr: '10'))
        disableConcurrentBuilds()
    }

    environment {
        APP_NAME = "devops-dashboard"
        IMAGE_NAME = "devops-dashboard"
    }

    stages {

        stage('Checkout Source') {
            steps {
                checkout scm
            }
        }

        stage('Repository Information') {
            steps {
                sh '''
                    echo "===== Repository ====="
                    pwd
                    ls -lah

                    echo
                    echo "===== Latest Commit ====="
                    git log --oneline -1
                '''
            }
        }

        stage('Verify Environment') {
            steps {
                sh '''
                    echo "===== Docker ====="
                    docker --version

                    echo
                    echo "===== Kubectl ====="
                    kubectl version --client

                    echo
                    echo "===== Helm ====="
                    helm version

                    echo
                    echo "===== Trivy ====="
                    trivy --version

                    echo
                    echo "===== Git ====="
                    git --version
                '''
            }
        }

        stage('Verify Kubernetes Access') {
            steps {
                sh '''
                    echo "===== Kubernetes Cluster ====="
                    kubectl get nodes
                '''
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '''
                    echo "===== Building Docker Image ====="

                    docker build \
                        -t ${IMAGE_NAME}:${BUILD_NUMBER} \
                        .

                    echo
                    echo "===== Docker Images ====="

                    docker images | grep ${IMAGE_NAME}
                '''
            }
        }
    }

    post {

        always {
            cleanWs()
        }

        success {
            echo "Pipeline completed successfully."
        }

        failure {
            echo "Pipeline failed."
        }
    }
}
