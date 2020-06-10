# build shoppie image
# docker build -t shoppie .

# launch shoppie backend service
docker run -d -p 8090:8080 -v $PWD/config/settings.yml:/root/config/settings.yml --rm -it --name shoppie-backend shoppie