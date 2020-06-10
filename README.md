# random_number
A simple "microservice" app for demoing [`tilt`](https://tilt.build/)'s Live Update capabilities, composed of:
* A Python+Flask server `numbers` which serves a random number
* A Golang server `fe` which serves the frontend and hits `numbers` for a random number

Check out the [Live Update Tutorial](https://docs.tilt.dev/live_update_tutorial.html) for more info.

For a slightly more complex version of this app, see [`abc123`](https://github.com/tilt-dev/abc123/).

## Quick start
1. [Install `tilt`](https://docs.tilt.dev/install.html)
2. `git clone https://github.com/windmilleng/random_number.git`
3. `cd random_number`
4. `tilt up`
5. When your services are green, navigate to `localhost:8080` to get your very own random number!
