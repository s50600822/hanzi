import os
from collections import Counter

ignored = {'(',')','<','>','＜','＞','[',']','{','}',':','.','。',',','!','?','+','-','*','/'}
in_file_ext = ".md" # better browsing on github

def process_text_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        text = file.read().replace('\n', '')
    return text

def process_directory(directory_path):
    all_text = ""
    for file_name in os.listdir(directory_path):
        if file_name.endswith(in_file_ext):
            file_path = os.path.join(directory_path, file_name)
            all_text += process_text_file(file_path)
    return all_text

def print_character_frequencies(text):
    # Remove spaces
    cleaned_text = ''.join(text.split())
    # Count character frequencies
    character_frequencies = Counter(cleaned_text)
    generate_md(character_frequencies.most_common())
    # Print characters ordered by frequencies in descending order
    for char, freq in character_frequencies.most_common():
        if char not in ignored:
            print(f"{char}: {freq}")

def generate_md(character_frequencies):
    table = "| Rank | Character | Frequency |\n|-----------|-----------|-----------|\n"
    total_count = 0

    for char, freq in character_frequencies:
        if char not in ignored:
            total_count += 1
            table += f"| {total_count} | {char} | {freq} |\n"

    # Add total row at the top
    table = f" {total_count}: characters \n" + table

    with open("Readme.md", 'w', encoding='utf-8') as readme_file:
        readme_file.write(table)

# Specify the directory containing text files
directory_path = './lyrics'

all_text = process_directory(directory_path)

print_character_frequencies(all_text)