# -*- mode: Python -*-

"""
* Frontend
  * Language: Go
  * Other notes: calls out to both `letters` and `numbers` microservices.
* Letters
  * Language: JavaScript
  * Other notes: Uses yarn. Does a `yarn install` for package dependencies iff they have changed.
* Numbers
  * Language: Python
  * Other notes: does a `pip install` for package dependencies. Re-installs dependencies iff they have changed.
"""
load('ext://restart_process', 'docker_build_with_restart')

k8s_yaml([
    'fe/fe.yaml',
    'numbers/numbers.yaml',
])

# Port-forward services so you can hit it them locally -- e.g. you
# can access the 'fe' service in your browser at http://localhost:8000
k8s_resource('fe', port_forwards='8000')
k8s_resource('numbers', port_forwards='8001')

# For all services, tell Tilt how to build the docker image, and how to Live Update
# that service -- i.e. how to update a running container in place for faster iteration.
# See docs: https://docs.tilt.dev/live_update_tutorial.html

# Service: fe
docker_build_with_restart('random_number/fe', 'fe',  # ≈ docker build -t random_number/fe ./fe
    entrypoint='/go/bin/fe',  # command to run on container start/re-run on live update
    live_update=[
        sync('./fe', '/go/src/github.com/windmilleng/random_number/fe'),
        run('go install github.com/windmilleng/random_number/fe'),
    ])

# Service: numbers
docker_build('random_number/numbers', 'numbers',  # ≈ docker build -t random_number/numbers ./numbers
    live_update=[
        sync('./numbers', '/app'),

        # run `pip install` IF `requirements.txt` has changed
        run('cd /app && pip install -r requirements.txt', trigger='numbers/requirements.txt'),
    ])
