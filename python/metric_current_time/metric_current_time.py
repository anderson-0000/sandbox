from http.server import BaseHTTPRequestHandler, HTTPServer
import time

class RequestHandler(BaseHTTPRequestHandler):
    def do_GET(request):
        if request.path == "/metric":
            request.send_response(200)
            request.send_header('Content-type', 'text/plain')
            request.end_headers()
            now = time.time()
            message = "# TYPE current_time_seconds gauge\n"
            message += "current_time_seconds %s\n" % int(now)
            message += "# TYPE current_time_nanoseconds gauge\n"
            message += "current_time_nanoseconds %s\n" % int(now * (10**9))
            request.wfile.write(bytes(message, "utf8"))
        return

def run():
    httpd = HTTPServer(('', 50005), RequestHandler)
    print('running server...')
    httpd.serve_forever()

if __name__ == '__main__':
    run()

