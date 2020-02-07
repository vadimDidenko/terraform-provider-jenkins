provider "jenkins" {
    url = "http://jenkins_url:port"
    username = "user"
    password = "pass"
}

resource "jenkins_username_credential" "admin" {
    identifier = "cred_id"
    username = "admin"
    password = "admin"
}