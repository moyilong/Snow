* Snow cluster test by light environment.
* Just only 3 instance

``` bash
# Startup by docker-compose
docker build -t snow-test-image ../../
docker compose up
```

or 

```bash
make test # only compose run
make test-image # build image
make test-rebuild # compose + image
```