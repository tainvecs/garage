# Fail2ban

Previously, I migrated from **"denyhosts"** to **"fail2ban"**.
The IPs banned by **"denyhosts"** in the past few years are stored in an artifact file.
To migrate all these IPs, I wrote a Python script to parse the banned IPs from the artifact and shell scripts to import those IPs to **"fail2ban"**.


## Environment

- Ubuntu 22.04.1
- fail2ban-client 0.11.2
- Python 3.9.12


## Parse IPs from a DenyHosts Artifact

Run the python script will read the artifact from `resources/hosts.deny` and generate output file to `output/banned_ip_list.txt` by default.
```bash
python scripts/generate_banned_ip_list_from_denyhosts_denied.py
```
To change the **input** or **output** file path, use arguments `--input` and `--output` when running the script.


## Import Banned IPs to Fail2ban

Run the shell script will read IPs from `resources/banned_ip_list.txt` and ban those IPs from sshd with fail2ban-client.

```bash
sh scripts/fail2ban-sshd_ban_ip_from_ip_list.zsh
```

After running the script, check the status with **fail2ban-client**.
```bash
sudo fail2ban-client status sshd
```
