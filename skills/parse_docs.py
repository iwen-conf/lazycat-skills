import re
import os

def parse_docs(filename, out_dir):
    with open(filename, 'r', encoding='utf-8') as f:
        content = f.read()

    # Split by exactly '---' on its own line
    parts = re.split(r'^---\n', content, flags=re.MULTILINE)
    
    docs = []
    current_url = None
    
    for part in parts:
        if part.startswith('url: '):
            current_url = part.split('url: ')[1].strip()
        elif current_url and part.strip():
            docs.append({
                'url': current_url,
                'content': part.strip()
            })
            current_url = None

    os.makedirs(out_dir, exist_ok=True)
    
    index_content = "# Lazycat Developer Docs Index\n\nWhen a user asks about any of the following topics or URLs, read the corresponding local file to get the full context.\n\n| URL | Local File |\n|---|---|\n"
    
    for doc in docs:
        # Generate safe filename
        safe_name = doc['url'].strip('/').replace('/', '_') + '.md'
        if not safe_name.endswith('.md.md'):
            # wait, if safe_name already has .md, it's fine. If not, add it.
            pass
        if safe_name.endswith('.md.md'):
            safe_name = safe_name[:-3]
        
        file_path = os.path.join(out_dir, safe_name)
        
        # Save content
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(f"# Source: https://developer.lazycat.cloud{doc['url']}\n\n")
            f.write(doc['content'])
            
        index_content += f"| `https://developer.lazycat.cloud{doc['url']}` | `references/docs/{safe_name}` |\n"
        
    with open(os.path.join(out_dir, 'INDEX.md'), 'w', encoding='utf-8') as f:
        f.write(index_content)
        
    print(f"Exported {len(docs)} documents to {out_dir}")

if __name__ == '__main__':
    parse_docs('skills/lazycat-developer-docs.md', 'skills/lazycat:developer-expert/references/docs')