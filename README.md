# ChatGPT-api-vercel

Deploy ChatGPT API proxy on vercel quickily.

## Preparation

1. Access to [vercel](https://vercel.com) and login with your github account.
2. OpenAPI key for [OpenAI](https://openai.com) and set it to `OPENAPI_KEY` in vercel environment variables.
3. Fork this repository.
4. [Supported Languages for Serverless Functions | Vercel Docs](https://vercel.com/docs/concepts/functions/serverless-functions/supported-languages#go)

## Create a project

On [vercel](https://vercel.com) dashboard, click `New Project` and select `Import Git Repository`.

Choose your forked repository, you don't need to change anything, and click `Import and Deploy`.

Now you can request vercel API endpoint, for example `https://[name].vercel.app/api/hello`.

## Add ChatGPT Key to environment variables

You can add ChatGPT key to environment variables on vercel dashboard, or use `vercel env` command.

```bash
$cd [root of this repository]

# Add an Environment Variable
$vercel env add CHATGPT_KEY [your ChatGPT key]

#  List all variables for the specified Environment
$vercel env list
```

## Deploy

You can deploy your project on vercel dashboard, or use `vercel` command.

```bash
# Deploy to vercel, and set the environment variable
$vercel

# Deploy to vercel, and set the environment variable for production
$vercel --prod
```

## Develop locally

You can develop locally with `vercel dev` command.

⚠️Go version should be 1.20 or higher. Keep in mind that Go version should greater than @vercel/go `go.mod` version. refer to [Getting package not found error when running vercel dev in my Go project](https://github.com/orgs/vercel/discussions/1850#discussioncomment-5401091)

### Structure

Vercel will use `api` directory as the entry point for your serverless functions. And the `utils` directory is used to store the common functions.

Function name can be any name, but should be capitalized. And function should respect [type http.HandlerFunc](https://pkg.go.dev/net/http#HandlerFunc)`


