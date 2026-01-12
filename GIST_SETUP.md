# Setting Up Private Gist for Daily Reflections

If you don't want to commit your daily reflections markdown file to the repository, you can store it in a private GitHub Gist.

## Step 1: Create a Private Gist

1. Go to https://gist.github.com/
2. Click the **"+"** button or go directly to https://gist.github.com/ to create a new gist
3. Set the filename to: `daily_reflections_structured.md`
4. Paste your markdown content with date headers (format: `## MM-DD`)
5. Select **"Create secret gist"** (NOT public)
6. Click **"Create secret gist"**

## Step 2: Get Your Gist ID

After creating the gist, look at the URL in your browser:

```
https://gist.github.com/YOUR_USERNAME/abc123def456ghi789jkl012mno345pq
                                      ↑
                                      This is your GIST_ID
```

Copy the alphanumeric string after your username. This is your `GIST_ID`.

## Step 3: Create a Personal Access Token (for private gists)

The workflow needs authentication to access your private gist.

1. Go to GitHub Settings → Developer settings → Personal access tokens → Tokens (classic)
   - Direct link: https://github.com/settings/tokens
2. Click **"Generate new token"** → **"Generate new token (classic)"**
3. Give it a descriptive name: `Telegram Bot Gist Access`
4. Set expiration (recommend: 1 year or no expiration for automation)
5. Select scope: **`gist`** (only this one is needed)
6. Click **"Generate token"**
7. **IMPORTANT:** Copy the token immediately (you won't see it again!)

## Step 4: Add GitHub Secrets and Variables

Go to your repository: **Settings → Secrets and variables → Actions**

### Add Secrets (Secrets tab)

Add or update these secrets:

| Secret Name | Value | Description |
|-------------|-------|-------------|
| `BOT_TOKEN` | `123456:ABC-DEF...` | Your Telegram bot token from BotFather |
| `GIST_TOKEN` | `ghp_abc123...` | Your personal access token from Step 3 |

### Add Variables (Variables tab)

Add or update these variables:

| Variable Name | Value | Description |
|---------------|-------|-------------|
| `CHAT_ID` | `-1001234567890` | Your group chat ID (negative number) |
| `MESSAGE_THREAD_ID` | `123` | Your topic/thread ID |
| `GIST_ID` | `abc123def456...` | Your gist ID from Step 2 |

## Step 5: Test the Workflow

1. Go to the **Actions** tab in your repository
2. Select **"Send Daily Message"** workflow
3. Click **"Run workflow"** → **"Run workflow"**
4. Watch the workflow run and check the logs

## Updating Your Daily Reflections

To update the content:

1. Go to your gist: https://gist.github.com/YOUR_USERNAME/YOUR_GIST_ID
2. Click **"Edit"**
3. Update the markdown content
4. Click **"Update secret gist"**

The next workflow run will automatically use the updated content!

## Troubleshooting

### Error: "Failed to download gist"
- Verify your `GIST_ID` is correct
- Ensure the `GIST_TOKEN` has the `gist` scope
- Check that the gist filename is exactly `daily_reflections_structured.md`

### Error: "404 Not Found"
- Your gist might be deleted or the ID is wrong
- The filename in the gist must match exactly: `daily_reflections_structured.md`

### Error: "401 Unauthorized"
- Your `GIST_TOKEN` is invalid or expired
- Generate a new token and update the secret

## Alternative: Public Gist (Not Recommended)

If you don't mind the content being public, you can create a **public gist** instead and skip the token:

1. Create a public gist instead of secret
2. Only add `GIST_ID` to secrets (no need for `GIST_TOKEN`)
3. Update the workflow to remove the Authorization header

**Note:** Anyone with the gist URL can view public gists.
