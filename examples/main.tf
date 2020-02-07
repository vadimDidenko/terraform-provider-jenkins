provider "jenkins" {
  url = "http://localhost"
  username = "user"
  password = "bitnami"
}

data "jenkins_username_credential" "admin" {
  identifier = "cred_id"
  username = "admin"
  password = "admin"
}

resource "jenkins_job_xml" "admin" {

  name = "test-job"
  xml = file("./job.xml")
}