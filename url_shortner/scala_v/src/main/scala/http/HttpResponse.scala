package http

object HttpResponse {
  def result(version: String, location: String): String =
    s"""
      |$version 301 Moved Permanently
      |Location: $location
      |
      |
      |""".stripMargin
}
