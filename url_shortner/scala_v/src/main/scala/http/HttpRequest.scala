package http

import java.io.BufferedReader

case class HttpRequest(
    requestMethod: String,
    requestURI: String,
    requestVersion: String,
)

object HttpRequest {
    private val REQUEST_HEADER_REGEX = "([a-zA-Z]+)\\s([a-zA-Z1-9#:/.?=]+)\\s(HTTP/1.1|HTTP/2)".r

    def build(br: BufferedReader): Option[HttpRequest] = {
        val header = br.readLine()
        println(s"header : $header")

        header match {
            case REQUEST_HEADER_REGEX(method, uri, version) => {
                println("http request header parsing done.")
                Some(HttpRequest(method, uri, version))
            }
            case _ =>
                throw new RuntimeException("[error] sorry cannot build http request.")
        }
    }
}
