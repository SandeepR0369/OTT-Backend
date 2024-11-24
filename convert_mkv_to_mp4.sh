#!/bin/bash

MOVIE_FOLDER="/Users/sandeepreddy/Movies"

# Loop through all .mkv files in the folder
find "$MOVIE_FOLDER" -type f -name "*.mkv" -print0 | while IFS= read -r -d '' mkv_file; do
  # Generate the output .mp4 filename
  mp4_file="${mkv_file%.*}.mp4"

  # Skip if the .mp4 file already exists
  if [ -f "$mp4_file" ]; then
    echo "Skipping: $mp4_file already exists."
    continue
  fi

  # Convert .mkv to .mp4
  echo "Converting: $mkv_file to $mp4_file"
  ffmpeg -nostdin -i "$mkv_file" -c:v copy -c:a aac -strict experimental "$mp4_file"
done

echo "Conversion complete!"
