# update the server using the following process:
# - Backup userdata.db (for user data)
# - Backup serverdata.db (for events state)
# - Comletely remove elichika and reinstall
# - Restore userdata.db and serverdata.db
# - Rebuild serverdata.db to new state
# running this command also potentially remove outdated backup

mv -f userdata.db ../userdata.db.temp && \
mv -f serverdata.db ../serverdata.db.temp && \
mv -f config.json ../config.json.temp && \
echo "Backed up databases, reinstalling" && \
cd .. && rm -rf elichika && \
curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/install.sh | bash && \
echo "Restoring old databases" && \
mv userdata.db.temp elichika/userdata.db && \
mv serverdata.db.temp elichika/serverdata.db && \
mv config.json.temp elichika/config.json && \
cd elichika && \
./elichika rebuild_assets && \
echo "Updated succesfully!"


if [ $? -eq 0 ]; then
    chmod +rwx ./bin/shortcut.sh && \
    ./bin/shortcut.sh
else
    echo "Error updating!"
fi