# genkit-golang-cloud-run-sample

`genkit-golang-cloud-run-sample` is a sample repository designed to help you learn how to build Large Language Model (LLM) applications using Firebase Genkit with Golang, deployed to Google Cloud Run.

- [Requirements](#requirements)
- [Usage](#usage)
- [License](#license)

## Requirements

- **Go**: Follow the [Go - Download and install](https://go.dev/doc/install) to install Go.
- **Genkit**: Follow the [Firebase Genkit - Get started](https://firebase.google.com/docs/genkit/get-started) to install Genkit.
- **Google Cloud CLI (gcloud)**: Follow the [Google Cloud - Install the gcloud CLI](https://cloud.google.com/sdk/docs/install) to install gcloud.
- **golangci-lint**: Follow the [golangci-lint - Install](https://golangci-lint.run/welcome/install/) to install golangci-lint.

Verify your installations:

```bash
$ go version
v22.4.1
$ genkit --version
0.5.4
$ gcloud --version
Google Cloud SDK 489.0.0
alpha 2024.08.16
bq 2.1.8
core 2024.08.16
gcloud-crc32c 1.0.0
gsutil 5.30
$ golangci-lint --version
golangci-lint has version 1.60.3 built with go1.23.0 from c2e095c on 2024-08-22T21:45:24Z
```

## Usage

### Run Genkit

Set your API key and start Genkit:

```bash
$ export GOOGLE_GENAI_API_KEY=your_api_key
$ make genkit # Starts Genkit
```

Open your browser and navigate to [http://localhost:4000](http://localhost:4000) to access the Genkit UI.

### Run HTTP Server Locally

To start the local http server, run the following command:

```bash
$ make dev
```

To test the application, use the following curl command:

```bash
$ curl -X POST -H "Content-Type: application/json" \
-d '"https://firebase.blog/posts/2024/04/next-announcements/"' http://localhost:3400/summarizeFlow
{"result": "Firebase announces new product updates including Firestore vector search, Vertex AI SDKs, and Gemini integration for AI-powered app development. \n"}
```

### Deploy

Set your secret values in `./.env.yaml`:

```bash
$ cp -p ./.env.example.yaml ./.env.yaml
$ vim ./.env.yaml # replace the secrets with your own values
GOOGLE_GENAI_API_KEY: your_api_key
```

Follow these steps to deploy the application:

```bash
$ gcloud auth application-default login
$ gcloud config set core/project [your-project-id]
$ make deploy
```

**CAUTION**: This deployment uses `.env.yaml` for environment variables, including the API key. This is not recommended for production. Instead, use Google Cloud Secret Manager for better security.

To test the deployed application, use the following curl command:

```bash
$ curl -X POST -H "Content-Type: application/json" \
-d '"https://firebase.blog/posts/2024/04/next-announcements/"' \
https://summarize-application-[your-cloud-run-id]-uc.a.run.app/summarizeFlow
{"result": "Firebase announces new product updates including Firestore vector search, Vertex AI SDKs, and Gemini integration for AI-powered app development. \n"}
```

Replace `[your-cloud-run-id]` with your Cloud Run service URL value, found in the Cloud Run Console.

### Code Formatting

To ensure your code is properly formatted, run the following command:

```bash
$ make tidy
```

## License

MIT
