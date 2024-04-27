# install the latest version of elichika from scratch
# set the environmental variable BRANCH to pick a specific branch (latest version only)
# this can be set with export or just set when invoking bash
# Honestly this is only for testing but you can think of it as a hidden feature
# if BRANCH is not provided, then default to master
BRANCH=${BRANCH:-"master"}
# install git and golang
pkg install golang git -y || echo "assuming go and git are already installed"
# clone the source code
git clone --depth 1 https://github.com/arina999999997/elichika.git --branch $BRANCH --single-branch && \
cd elichika && \
# get the submodules (i.e. assets and other)
git submodule update --init && \
# build server, fallback to not using CGO to work on some devices
(go build || CGO_ENABLED=0 go build) && \
# set the permission
chmod +rx elichika && \
echo "Installed succesfully!"

if [ $? -eq 0 ]; then
    echo "cd $PWD && ./elichika" > ~/run_elichika && \
    echo "cd $PWD && git pull && \
    git submodule deinit -f . && \
    git submodule update --init --recursive --checkout && \
    (go build || CGO_ENABLED=0 go build)" > ~/update_elichika && \
    chmod +x ~/run_elichika && \
    chmod +x ~/update_elichika && \
    echo "Use \"~/run_elichika\" in termux to run the server!" && \
    echo "Use \"~/update_elichika\" in termux to update the server!"
else
    echo "Error installing"
fi