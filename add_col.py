import pandas as pd

# Read the existing Markdown table from a file
file_path = './poems/proverb.md'  # Replace with the actual path to your file
df = pd.read_table(file_path, sep='|',header=0,  skipinitialspace=True, engine='python').dropna(axis=1, how='all')


# Strip whitespace from column names
df.columns = df.columns.str.strip()
df = df.sort_values(by='hanzi')

# add col
df['new_col'] = ['dummy'] * len(df)

# Convert the DataFrame to a Markdown table
markdown_table = df.to_markdown(index=False)

# review
print(markdown_table)
# Save the updated Markdown table to a new file
new_file_path = 'updated_table.md'  # Replace with the desired path for the new file
with open(new_file_path, 'w') as file:
    file.write(markdown_table)
