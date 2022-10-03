"""prune hosts.deny file and generate banned ip list"""

import os
import argparse
import logging
import re


class IpInfo:

    """IP String Info"""

    ip_str  = ""
    ip_list = []

    def __init__(self, _ip: str, only_ipv4=False):

        self.ip_str  = _ip

        if only_ipv4:
            self.ip_list = [ int(sub_ip) for sub_ip in re.split(r'[\.\:]', _ip) ]
        else:
            self.ip_list = re.split(r'[\.\:]', _ip)



def parse_args():

    """parse and process input arguments"""

    parser = argparse.ArgumentParser("prune hosts.deny file and generate banned ip list")

    parser.add_argument(
        "--input",
        default="../resources/hosts.deny",
        type=str,
        help='the relative file path of input denyhost hosts.deny'
    )

    parser.add_argument(
        "--output",
        default="../output/banned_ip_list.txt",
        type=str,
        help='the relative file path of output file for banned ip'
    )

    parser.add_argument(
        "--only_ipv4",
        default=False,
        type=bool,
        help='set if only input ipv4'
    )

    raw_args = parser.parse_args()

    return raw_args


if __name__ == '__main__':

    args = parse_args()
    logging.basicConfig(encoding='utf-8', level=logging.INFO)

    code_path = os.path.dirname(os.path.abspath(__file__))

    # load input file
    deny_file_path = os.path.join(code_path, args.input)
    if not os.path.isfile(deny_file_path):
        logging.warning("input file %s not exist", args.input)

    ip_info_list = []
    with open(deny_file_path, 'r', encoding='utf-8') as infile:

        for line in infile:

            # filter empty line
            if not line.strip():
                continue

            # skip comment
            if re.match(r'\s*#.*', line):
                continue

            # match ip str
            ip_match = re.search(r'\s*sshd:\s*([\.\:[a-z0-9]+)\s*', line)
            if not ip_match:
                continue

            ip_str = ip_match.group(1).strip()
            ip_info_list.append(
                IpInfo(ip_str, only_ipv4=args.only_ipv4)
            )

    # sort and to string list
    sorted_ip_info_list = sorted(
        ip_info_list,
        key=lambda s: s.ip_list,
    )
    ip_str_list = [ ip_info.ip_str for ip_info in sorted_ip_info_list ]

    # output
    out_file_path = os.path.join(code_path, args.output)
    with open(out_file_path, 'w', encoding='utf-8') as ofile:
        ofile.write('\n'.join(ip_str_list)+'\n')

    print("extract {} ips from {}".format(
        len(ip_str_list),
        args.input,
    ))
