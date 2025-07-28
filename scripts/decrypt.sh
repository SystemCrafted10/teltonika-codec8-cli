#!/bin/bash

# Check if password is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <password>"
  exit 1
fi

PASSWORD=$1

# Function to decrypt a file
decrypt_file() {
  local file=$1
  local decrypted_file="${file%.gpg}"
  echo "Decrypting $file to $decrypted_file..."
  echo "$PASSWORD" | gpg --batch --yes --passphrase-fd 0 --decrypt "$file" > "$decrypted_file"
}

# Files to be decrypted
files_to_decrypt=(
  .env.gpg
  .env.development.gpg
  _notes.md.gpg
)

# Find and decrypt the files
for file in "${files_to_decrypt[@]}"; do
  if [ -f "$file" ]; then
    decrypt_file "$file"
  else
    echo "$file not found, skipping."
  fi
done
