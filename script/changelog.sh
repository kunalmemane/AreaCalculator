#!/bin/bash

TAGS=($(git tag --sort=-creatordate | head -n 2))
LATEST_TAG="${TAGS[0]}"
PREVIOUS_TAG="${TAGS[1]}"

if [ ${#TAGS[@]} -lt 2 ]; then
    echo "Not enough tags found to generate changelog."
    exit 1
fi

CURRENT_DATE=$(date "+%Y-%m-%d")

echo -e "\n# Changelog\n"
echo -e "## Release Notes from $PREVIOUS_TAG to $LATEST_TAG\n"
echo -e "### $CURRENT_DATE\n"

declare -A tag_map=(
    ["chore"]="Minor change"
    ["docs"]="Document"
    ["feat"]="Feature Added"
    ["bug"]="Fixed"
    ["ci"]="CI"
    ["refactor"]="Changed"
    ["test"]="Test"
)

# Collect and categorize commit messages
for type in "${!tag_map[@]}"; do
    
    title="${tag_map[$type]}"

    if git log "$PREVIOUS_TAG..$LATEST_TAG" --pretty='%h - %s (%an)' | grep -q "$type"; then
        echo -e "\n### $title \n"

       git log "$PREVIOUS_TAG..$LATEST_TAG" --pretty='%h - %s (%an)' | awk -v type="$type" '$0 ~ "^.{10}(" type ": )" { sub(type ": ", ""); print }'
    fi 
done

echo -e "\nChangelog generated for release $LATEST_TAG."