def application(env, start_resp):
    start_resp('200 OK', [('Content-Type', 'text/html')])
    return [b'<h1>Hello, Python Web</h1>']