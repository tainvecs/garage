GARAGE_F2B_ROOT="$(dirname $(cd $(dirname $0) >/dev/null 2>&1; pwd -P;))"
DENYHOST_DENY_FILE=$GARAGE_F2B_ROOT/resources/hosts.deny

i=0
all_n=$(wc -l < $DENYHOST_DENY_FILE)
while IFS="" read -r in_str; do

    ((i+=1))

    # reomve comment
    in_str="$(echo $in_str | sed 's/^#.*//g')"

    # remove sshd prefix
    ip_str="$(echo $in_str | sed 's/sshd\:\(.*\)/\1/')"

    # trim space
    trimmed_ip_str="$(echo $ip_str | sed 's/\s*//g')"

    # skip empty
    if [[ -z "$trimmed_ip_str" ]]; then
        continue
    fi

    sudo fail2ban-client set sshd banip "$trimmed_ip_str"
    echo " $i / $all_n: banned $trimmed_ip_str"

done < $DENYHOST_DENY_FILE
