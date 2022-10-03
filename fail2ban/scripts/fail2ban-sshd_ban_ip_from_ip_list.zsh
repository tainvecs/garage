GARAGE_F2B_ROOT="$(dirname $(cd $(dirname $0) >/dev/null 2>&1; pwd -P;))"
BANNED_IP_FILE=$GARAGE_F2B_ROOT/resources/banned_ip_list.txt

i=0
all_n=$(wc -l < $BANNED_IP_FILE)
all_n="$(echo $all_n | sed 's/\s*//g')"

while IFS="" read -r in_ip; do

    ((i+=1))

    # trim space
    trimmed_ip="$(echo $in_ip | sed 's/\s*//g')"

    # skip empty
    if [[ -z "$trimmed_ip" ]]; then
        continue
    fi

    sudo fail2ban-client set sshd banip "$trimmed_ip"
    echo " $i / $all_n: banned $trimmed_ip"

done < $BANNED_IP_FILE
