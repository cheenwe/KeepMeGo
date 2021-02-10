# KeepMeGO
Keep me go api


## Support platform
* docker hub


## run and deployment
It's strongly recommend to use docker to run this project 
because we're about to running some system commands from untrusted sources.

```shell script
git clone https://github.com/cheenwe/KeepMeGO
cd KeepMeGO
# change your token here, you may also add other environment variables such as `http_proxy`
vim config.ini
# create your db
touch keep.db
docker-compose up -d
```
Of course you could build your own docker image
`docker build -t keepmego .

### How to update using docker-compose
1. Use docker pull to update docker-image, and run again

## License
MIT
