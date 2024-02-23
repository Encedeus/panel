sudo apt update && sudo apt upgrade

# Install Docker Engine
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
# Start Docker on boot
sudo systemctl enable docker.service && sudo systemctl start docker.service
sudo systemctl enable containerd.service && sudo systemctl start containerd.service

# Install skyhook
wget https://cdn.encedeus.com/bin/skyhook_latest_amd64_linux
mkdir -p /home/encedeus/skyhook
mv skyhook_latest_amd64_linux /home/encedeus/skyhook/
echo -e "[Unit]\nDescription=Encedeus Skyhook instance\n\n[Service]\nExecStart=/home/encedeus/skyhook/skyhook_latest_amd64_linux\n\n[Install]\nWantedBy=multi-user.target" > /etc/systemd/system/encedeus_skyhook.service
# Start Skyhook on boot
sudo systemctl enable encedeus_skyhook.service
sudo systemctl start encedeus_skyhook.service

