#!/usr/bin/env bash
set -euo pipefail

STATE_FILE="nginx-blocklist/ips.txt"

FILE="ips.txt"
TMP_FILE="ips.txt.tmp"

mkdir -p "$(dirname "$STATE_FILE")"
touch "$STATE_FILE"

is_valid_target() {
   local target="$1"
   [[ "$target" =~ ^([0-9]{1,3}\.){3}[0-9]{1,3}(/([0-9]|[1-2][0-9]|3[0-2]))?$ ]]
}

cmd="${1:-}"

echo "$cmd"

case "$cmd" in
  add)
     ip="${2:-}"
     is_valid_target "$ip" || { echo "blocklist IP/CIDR: $ip"; exit 1; }
     echo "$ip" >> "$STATE_FILE"
     ;;
  list)
     # ip =""
     ip="${2:-}"
     cd nginx-blocklist && cat ips.txt
     ;;
  del)
     ip="${2:-}"
     is_valid_target "$ip" || { echo "blocklist IP/CIDR: $ip"; exit 1; }
     echo "$ip"
     cd nginx-blocklist
     if [[ -n "$ip" ]]; then
        echo "$ip not found" >&2
        exit 1
     fi
     awk -v target="$ip" '
        {
          line = $0
          sub(/\r&/, "", line)
          gsub(/^[[:space:]]+|[[:space:]]+$/, "", line)
          if (line != target) print $0
        }
     ' "$FILE" > "$TMP_FILE" && mv "$TMP_FILE" "$FILE"

     echo "$ip delete"
     ;;
  *)
     ;;
esac
