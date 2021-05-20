import server.ServerMutex

object Main {
  def main(args: Array[String])
  {
    ServerMutex.add("/gg", "https://www.google.com")
    ServerMutex.add("/go", "https://www.golang.org")

    val mux = ServerMutex.create(9000)

    mux.listenAndServe()
  }
}
