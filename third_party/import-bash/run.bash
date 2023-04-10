#!/usr/bin/env bash

script_path=$(dirname $(realpath "$0"))

read leaksdb_fp notify_url < <(echo $(cat "$script_path/args.json" | jq -r '.leaksdb_fp, .notify_url'))

while read -r leak_fp
do
    read -r context
    read -r leakers
    read -r platforms
    read -r share_date

    docker run \
        --mount "type=bind,src=$leaksdb_fp,dst=$leaksdb_fp" \
        --mount "type=bind,src=$leak_fp,dst=$leak_fp" \
        import --database-path="$leaksdb_fp" --leak-path="$leak_fp" --context="$context" --platforms="$platforms" --share-date="$share_date" --leakers="$leakers" --notify-url="$notify_url" --skip=true
done < <( cat "$script_path/args.json" | jq -cr '.leaks[] | (.leak_fp, .context, .leakers, .platforms, .share_date)')