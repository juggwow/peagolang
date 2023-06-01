pipeline {
    agent any
    
    stages {
        stage('Build') {
            steps {
                sh """
                    docker build --rm \
                    -f Dockerfile \
                    -t registry.hub.docker.com/patna/api \
                    -t registry.hub.docker.com/patna/api \
                    .
                """
            }
        }
        

        stage('Push') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'patna-docker', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                sh """
						docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD registry.hub.docker.com
                		docker push registry.hub.docker.com/patna/static-web:v2
					  """
                }
            }
        }

        stage('Deploy to server') {
            steps {
                script {
                    sshagent (credentials: ["DEV-Server"]){
                        sh """
                            ssh -o StrictHostKeyChecking=no -l ubuntu 13.229.66.4 'mkdir -p patna/api/'
                            scp docker-compose.yaml ubuntu@13.229.66.4:patna/api/

                            ssh -o StrictHostKeyChecking=no -l ubuntu 13.229.66.4 \"
                                cd patna/api/
                                docker compose -f docker-compose.yaml up -d
                            \"
                        """
                    }
                }
            }
        }
    }
}

