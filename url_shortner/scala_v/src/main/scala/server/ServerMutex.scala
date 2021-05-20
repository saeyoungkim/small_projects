package server

import java.io.{BufferedReader, InputStreamReader, PrintWriter}
import java.net.{InetSocketAddress, ServerSocket, Socket}

import http.{HttpRequest, HttpResponse}

import scala.collection.mutable
import scala.collection.mutable.{ArrayBuffer, Map}
import scala.concurrent.{Await, Future}
import scala.util.{Failure, Success, Try}
import scala.concurrent.ExecutionContext.Implicits.global
import scala.concurrent.duration.Duration

case class ServerMutex(
  private val port: Int
) {
  private val serverSocket: ServerSocket = createServerSocket()

  private def createServerSocket(): ServerSocket = {
    val sock = for {
      listenSocket <- Try(new ServerSocket())
      _ <- Try(listenSocket.bind(new InetSocketAddress("localhost", this.port)))
    } yield listenSocket

    sock match {
      case Success(ss) => ss
      case Failure(ex) =>
        println("[error] failed to initialization")
        throw ex
    }
  }

  private def listenAndServeUtil(sock: Socket): Future[Unit] = Future {
    val reader: BufferedReader = new BufferedReader(new InputStreamReader(sock.getInputStream))
    val writer: PrintWriter = new PrintWriter(sock.getOutputStream, true)

    val httpRequestOpt = HttpRequest.build(reader)

    httpRequestOpt match {
      case Some(httpRequest) => redirect(httpRequest, writer)
      case None => throw new RuntimeException("リクエストの型が正しくありません")
    }

    reader.close
    writer.close
    sock.close
  }

  def listenAndServe(): Unit = {
    val futureSeq = ArrayBuffer.empty[Future[Unit]]
    while(true) {
      val sock = this.serverSocket.accept()
      val f: Future[Unit] = listenAndServeUtil(sock)
      futureSeq += f

      f.onComplete {
        case Success(_) => {
          println("serve done")
        }
        case Failure(_) => {
          println("failure callback called...")
          futureSeq.foreach(Await.result(_, Duration.Inf))
          throw new RuntimeException("cannot serve correctly.")
        }
      }
    }
  }

  private def redirect(request: HttpRequest, pw: PrintWriter): Unit = {
    ServerMutex.pathMap.get(request.requestURI) match {
      case Some(path) => {
        println(s"send redirect response.)\n${HttpResponse.result(request.requestVersion, path)}")
        pw.write(HttpResponse.result(request.requestVersion, path))
      }
      case None => pw.write("hi")
    }
  }
}

object ServerMutex {
  private val pathMap: Map[String, String] = mutable.HashMap[String, String]()

  def add(oldPath: String, newPath: String): Unit = {
    pathMap += (oldPath -> newPath)
  }

  def create(port: Int = 8080): ServerMutex = ServerMutex(port)
}
