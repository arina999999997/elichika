# update the server using the following process:
# - update the code (elichika)
# - update the submodules (assets)
# - rebuild the binary
# - rebuild the assets
# note that this will destroy the current state of serverdata.db
# backup you files, or better yet, make sure that serverdata.db only store derived data

git pull && \
git submodule deinit -f . && \
git submodule update --init --recursive --checkout && \
(go build || CGO_ENABLED=0 go build) && \
elichika reinit && \
echo "Updated succesfully!"


if [ $? -eq 0 ]; then
    echo "cd $PWD && ./elichika" > ~/run_elichika && \
    echo "cd $PWD && curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/install.sh | bash"  > ~/update_elichika && \
    chmod +x ~/run_elichika && \
    chmod +x ~/update_elichika && \
    echo "Use \"~/run_elichika\" in termux to run the server!" && \
    echo "Use \"~/update_elichika\" in termux to update the server!"
else
    echo "Error updating!"
fi