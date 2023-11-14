
function download_db() {
    lang=$1
    echo "Downloading asset files for $lang"
    asset_path="2d61e7b4e89961c7"
    if [ $lang == "ja" ]
    then
        asset_path="b66ec2295e9a00aa"
    fi
    cdn_server="https://llsifas.catfolk.party/static/$asset_path"
    files=("masterdata_a_$lang" "masterdata_i_$lang" "asset_a_$lang.db" "dictionary_$lang""_android.db" "asset_i_$lang.db"
        "dictionary_$lang""_dummy.db" "dictionary_$lang""_inline_image.db" "dictionary_$lang""_ios.db" "dictionary_$lang""_k.db"
        "dictionary_$lang""_m.db" "dictionary_$lang""_petag.db" "dictionary_$lang""_s.db" "dictionary_$lang""_v.db" 
    )

    if [ $lang == "ja" ] || [ $lang == "en" ]
    then
        files+=("masterdata.db")
    fi

    mkdir -p $"static/$asset_path"
    for file in ${files[@]}
    do
        echo -n "Downloading $file: "
        exit_code=1
        try=3
        while [ $exit_code != 0 ] && [ $try != 0 ]
        do
            try=$(($try-1))
            curl -L -s "$cdn_server/$file" -o "static/$asset_path/$file"
            
            exit_code=$?
            if [ $exit_code != 0 ]
            then
                if [ $try != 0 ]
                then
                    echo "Failed to download, retrying ($try left)" 
                else
                    echo "Failed to download assets, retry later!"
                    return $exit_code
                fi
            else
                echo "Success!"
            fi
        done
    done
    return 0
}
# install this version of elichika from scratch
# download this file manually and run it
# assume this is a fresh install
# install git and golang
pkg install golang git -y && \
# clone the source code
git clone https://github.com/arina999999997/elichika.git && \
cd elichika && \
# build server
go build && \
# download the patched database
download_db "ja" && \
download_db "en" && \
download_db "ko" && \
download_db "zh" && \
# download the config file
# set the permission
chmod +rx elichika && \
echo "Installed succesfully!"
# todo: edit .bashrc to make an easier command
