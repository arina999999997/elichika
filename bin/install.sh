# install this version of elichika from scratch
# download this file manually and run it
# assume this is a fresh install
# install git and golang
pkg install golang git -y || echo "assuming go and git are already installed"
# clone the source code
git clone https://github.com/arina999999997/elichika.git && \
cd elichika && \
# build server
go build && \
# set the permission
chmod +rx elichika && \
echo "Installed succesfully!"
# todo: edit .bashrc to make an easier command
