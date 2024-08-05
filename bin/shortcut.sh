# Setup shortcuts
echo "cd $PWD && ./elichika" > ~/run_elichika && \
echo "cd $PWD && curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/update.sh | bash"  > ~/update_elichika && \
echo "cd $PWD && curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/basic_update.sh | bash"  > ~/basic_update_elichika && \
chmod +x ~/run_elichika && \
chmod +x ~/update_elichika && \
chmod +x ~/basic_update_elichika && \
echo "Use \"~/run_elichika\" in termux to run the server!" && \
echo "Use \"~/update_elichika\" in termux to update the server!"
echo "Use \"~/basic_update_elichika\" in termux to update the server using basic logic (will be slower but will work even if you have a really old version)!"