provider "jenkins" {
    url = "http://192.168.1.145:25551"
    username = "user"
    password = "pass"
}

resource "jenkins_username_credential" "admin" {
    identifier = "cred_id"
    username = "admin"
    password = "admin"
}

resource "jenkins_job_xml" "admin" {
    name = "cred_id"
    xml = "admin"
}