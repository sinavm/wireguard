# PureVPN WireGuard Improved Fetcher

Improved version of Rikpat/wireguard with GitHub Actions integration, secrets management, and GitHub Pages UI for easy config generation.

## Setup

1. Fork/clone this repo.
2. Add secrets: PUREVPN_USERNAME, PUREVPN_PASSWORD, PUREVPN_SUB_USER.
3. Enable GitHub Pages on gh-pages branch.
4. Replace 'yourusername' in go.mod, main.go, index.html with your GitHub username.
5. Commit and push to gh-pages for Pages.
6. For city IDs, adjust in workflow if needed (extract from browser dev tools on PureVPN site).

## Usage

- Local: `go run cmd/wireguard full` with env vars.
- Docker: `docker run -e PUREVPN_USERNAME=... yourimage`.
- UI: Visit Pages site, click button, enter token, check Actions logs for config (copy to clipboard).

## Improvements

- Secrets via GitHub Actions.
- UI with location buttons.
- No plaintext creds in code/config.
- Basic error handling.
- Artifact upload for download.

Note: Selectors in chromedp may need update if PureVPN changes UI.