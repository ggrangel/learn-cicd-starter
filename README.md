# Project Goals

- Set up a continuous integration pipeline with GitHub Actions that ensures new PRs pass certain checks before they are merged to main:
  - Unit tests pass
  - Formatting checks pass
  - Linting checks pass
  - Security checks pass
  
- Configured a cloud-based SQLite database hosted on Turso

- Set up a continuous deployment pipeline with GitHub Actions that does the following whenever changes are merged into main:
  - Builds a new server binary
  - Builds a new Docker image for the server
  - Pushes the Docker image to the Google Artifact Registry
  - Deploys a new Cloud Run revision with the new image and serves the app to the public internet
