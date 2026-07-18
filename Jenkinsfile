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
        DOCKER_USERNAME = "navneet2004"
        IMAGE_NAME = "navneet2004/devops-dashboard"
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
                        -t ${APP_NAME}:${BUILD_NUMBER} \
                        .
                '''
            }
        }

        stage('Push Docker Image') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {

                    sh '''
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin

                        docker tag \
                            ${APP_NAME}:${BUILD_NUMBER} \
                            ${IMAGE_NAME}:${BUILD_NUMBER}

                        docker push ${IMAGE_NAME}:${BUILD_NUMBER}

                        docker logout
                    '''
                }
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
