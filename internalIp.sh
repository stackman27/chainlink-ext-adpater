# Replace containerId with the actual ID of your running container
containerId="71b073c6940d"

# Get the internal IP address
internalIp=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $containerId)

# Print the internal IP address
echo "Internal IP Address of Container $containerId: $internalIp"




