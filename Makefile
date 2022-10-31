sudo docker image build -f Dockerfile . -t lab3
sudo docker run -p 1323:1323 --name cont lab3