sudo docker image build -f Dockerfile . -t lab4
sudo docker run -p 1323:1323 --name cont lab4