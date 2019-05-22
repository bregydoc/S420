docker-compose -f minio-compose.yml up -d 
docker-compose -f s420-compose.yml up -d 
echo "Now you can access to your files from web at 3300 and with gRPC at 24000"