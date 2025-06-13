import os
import json
import re
import requests

SBOM_FILE = "sbom.json"
OUTPUT_FILE = "NOTICE.md"
INTERNAL_DEPENDENCIES = {"actions/checkout", "copy-to-branches"}
GITHUB_API = "https://api.github.com/repos"
GITHUB_TOKEN = os.getenv("GITHUB_TOKEN", "")

HEADERS = {
    "Accept": "application/vnd.github.v3+json"
}
if GITHUB_TOKEN:
    HEADERS["Authorization"] = f"Bearer {GITHUB_TOKEN}"


def is_internal(name):
    for term in INTERNAL_DEPENDENCIES:
        if term.lower() in name.lower():
            return True
    return False


def extract_github_repo(purl):
    match = re.search(r"github\.com/([^@]+)", purl)
    return match.group(1) if match else None


def get_repo_license(repo):
    url = f"{GITHUB_API}/{repo}/license"
    res = requests.get(url, headers=HEADERS)
    if res.status_code != 200:
        return None, None, None
    data = res.json()
    license_name = data.get("license", {}).get("name", "UNKNOWN")
    license_url = data.get("license", {}).get("url", "")
    license_text = requests.get(data.get("download_url", "")).text
    return license_name, license_url, license_text


def parse_sbom(sbom_path):
    with open(sbom_path, "r") as f:
        data = json.load(f)
    packages = data.get("packages", [])
    deps = []

    for pkg in packages:
        name = pkg.get("name", "")
        purls = [ref["referenceLocator"] for ref in pkg.get("externalRefs", []) if "purl" in ref.get("referenceType", "")]
        if not purls:
            continue

        purl = purls[0]
        if is_internal(name):
            continue

        version = pkg.get("versionInfo", "")
        github_repo = extract_github_repo(purl)
        deps.append({
            "name": name,
            "version": version,
            "purl": purl,
            "repo": github_repo
        })
    return deps


def write_notice(dependencies):
    with open(OUTPUT_FILE, "w") as f:
        # Boilerplate header
        f.write("""Copyright 2025 LinkedIn Corporation
All Rights Reserved.

Licensed under the LinkedIn Learning Exercise File License (the "License").
See LICENSE in the project root for license information.

ATTRIBUTIONS:\n\n""")

        # Metadata section
        license_blocks = []
        for dep in dependencies:
            repo_url = f"https://github.com/{dep['repo']}" if dep["repo"] else ""
            license_name, license_url, license_text = get_repo_license(dep["repo"]) if dep["repo"] else ("UNKNOWN", "", "")
            copyright_line = f"Copyright © {dep['name']} contributors"

            f.write(f"{dep['name']}\n")
            f.write(f"{repo_url}\n")
            f.write(f"{copyright_line}\n")
            f.write(f"License: {license_name}\n")
            f.write(f"{license_url}\n\n")

            if license_text:
                license_blocks.append((dep["name"], license_text.strip()))

        # Footer text
        f.write("""Please note, this project may automatically load third party code from external 
repositories (for example, NPM modules, Composer packages, or other dependencies). 
If so, such third party code may be subject to other license terms than as set 
forth above. In addition, such third party code may also depend on and load 
multiple tiers of dependencies. Please review the applicable licenses of the 
additional dependencies.

===========================================================\n\n""")

        # License text sections
        for name, text in license_blocks:
            f.write(f"{name}:\n\n{text}\n\n")
            f.write("===========================================================\n\n")


if __name__ == "__main__":
    print("🔍 Parsing SBOM and generating NOTICE.md...")
    deps = parse_sbom(SBOM_FILE)
    write_notice(deps)
    print(f"✅ NOTICE.md generated with {len(deps)} dependencies.")
