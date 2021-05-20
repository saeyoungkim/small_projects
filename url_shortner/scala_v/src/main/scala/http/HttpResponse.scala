package http

object HttpResponse {
  val CRLF = "\\r\\n"

  def redirectResult(version: String, location: String): String =
    s"""
      |$version 301 Moved Permanently
      |Location: $location
      |Content-Length: 0
      |Connection: Close
      |
      |""".stripMargin

  def result(version: String, path: String, body: String): String =
    s"""
       |$version 404 Not Found
       |Content-Length: ${body.length.toString}
       |Connection: Close
       |
       |$body
       |""".stripMargin
}
