# update the server using the following process:
# - update the code (elichika)
# - update the submodules (assets)
# - rebuild the binary
# - rebuild the assets
# note that this will destroy the current state of serverdata.db and the submodules
# backup you files, or better yet, make sure that serverdata.db only store derived data or special supported data (event)

git pull && \
git submodule deinit -f assets && \
git submodule update --init --recursive --checkout assets && \
(go build || CGO_ENABLED=0 go build) && \
./elichika rebuild_assets && \
echo "Updated succesfully!"


if [ $? -eq 0 ]; then
    chmod +rwx ./bin/shortcut.sh && \
    ./bin/shortcut.sh
else
    echo "Error updating!"
fi