import argparse
from functools import partial
from random import randint

from flask import Flask
app = Flask(__name__)


# automatically flush stdout when logging so messages show up
log = partial(print, flush=True)


@app.route('/')
def rand_number():
    log("🎲 Randomizing the numbers...")
    return str(randint(0, 100))+'\n'


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='start an http server that returns numbers')
    parser.add_argument('--port', type=int, default=5000, help='port on which to serve http')
    args = parser.parse_args()
    app.logger.info('Starting `number` server on port {}'.format(args.port))
    app.run(debug=True, host='0.0.0.0', port=args.port)
