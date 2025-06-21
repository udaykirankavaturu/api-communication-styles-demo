from spyne import Application, rpc, ServiceBase, Integer, Unicode
from spyne.protocol.soap import Soap11
from spyne.server.wsgi import WsgiApplication
from wsgiref.simple_server import make_server

class HelloWorldService(ServiceBase):
    @rpc(Unicode, Integer, _returns=Unicode)
    def say_hello(ctx, name, times):
        return u' '.join([f"Hello, {name}!" for _ in range(times)])

application = Application([
    HelloWorldService
], 'spyne.examples.hello.soap',
    in_protocol=Soap11(validator='lxml'),
    out_protocol=Soap11())

wsgi_app = WsgiApplication(application)

if __name__ == '__main__':
    server = make_server('0.0.0.0', 8000, wsgi_app)
    print("SOAP server running on http://0.0.0.0:8000")
    server.serve_forever()
