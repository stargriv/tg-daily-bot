# Daily Reflections Directory

This directory contains your daily reflections markdown file.

## For Local Development

Place your `daily_reflections_structured.md` file here with the following format:

```markdown
## 01-01

*Your reflection or quote here*

Content for January 1st goes here.
This can be multiple paragraphs.

*Closing thought or meditation*

## 01-02

Content for January 2nd...
```

**Format Requirements:**
- Each date entry must start with: `## MM-DD` (e.g., `## 01-15` for January 15)
- Content continues until the next date header
- Supports basic markdown formatting (bold `**text**`, italic `*text*`)

## For GitHub Actions

The `files/` directory is ignored by git by default (see `.gitignore`).

**You have two options:**

1. **Commit the file to git**: Remove `files/` from `.gitignore` and commit your markdown file
2. **Use a private GitHub Gist**: Store your file in a private gist (see [GIST_SETUP.md](../GIST_SETUP.md))

The GitHub Actions workflow will automatically fetch from the gist if configured.
