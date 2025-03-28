# The New Hire Challenge 🚀

## Welcome 🎉

You have just embarked on your journey in this team, and what better way to make an impression than by solving a **real-world problem** that has been haunting us for ages? Well, at least making us wait forever...

Fear not! Your challenge is clear, and the path to success is paved with clean code and high performance.

---

## Background Story 📖

In our glorious quest for efficiency, we are required to export **all certificates from our database separately** in the **PEM format**. Simple, right? 🧐

- We always seem to be running late. Deadlines are looming. The pressure is real. 😨
- The current solution? A **Bash script** that takes so long that we suspect it's secretly mining crypto. ⛏️
- We value **high-quality software**, so unit test coverage must be at least **60%**. (Yes, that number is non-negotiable. *Das ist ein Muss.*) 
- **Cross-compilation** is crucial! Due to budget constraints (since the company prefers to cover employees' business trips where they indulge in sushi at posh restaurants ([1](https://www.instagram.com/nijulondon/), [2](https://www.instagram.com/sushisamba/)) in London 🍣 and sip cocktails at rooftop bars 🍸 ([1](https://www.instagram.com/sg_skygarden/)), we have a fragmented environment - Windows, Linux, and macOS users all in one team.

With that in mind, our language of choice is [Golang](https://go.dev/)! It builds things fast and efficiently, just like we expect from you.

## Your Mission (Should You Choose to Accept It)

1. Analyze the existing Bash script. Brace yourself; it’s not pretty.
2. Write a better, faster (`goroutine` 💨), and more elegant solution in Go.
3. Ensure unit test coverage is **at least** 60%.
4. Add meaningful debug logging and error handling.
5. Ensure that directory names are derived from the issuing CA name, replacing all whitespace and hyphens with underscores.
6. Profit! 🚀

---

## Technical Requirements ⚙️

- Stick to **pure Go** - use only standard Go libraries or well-maintained third-party Go modules. 📦
- **No shell execution allowed** (e.g., calling `openssl` or similar tools from within your Go code). 🚫
- **No CGO dependencies** - your solution must be easily cross-compiled without additional system dependencies. 🚫

## What We Expect from You 🏆

- Readable, maintainable, and efficient code.
- Proper unit tests. Yes, we are going to check. No, you can’t fake it. 🥲
- Output directory structure must be the following:

  ```shell
    # Output Directory Structure:
    .
    ├── issuing_CA_1
    │   ├── 221256C1E1D18FC98BE83E63D4434F09A73FBF95.pem
    │   ├── 4EEC8531F25DDE0B0A21B6AC7AFBE6C2AF802F48.pem
    ├── issuing_CA_2
    │   ├── 1B1593132285EC4282030192D3E9B77354B886D1.pem
  ```
  
  > Go represents certificate serial numbers as `big.Int`, but we need them in **hex format**. Because apparently, life wasn’t challenging enough already.

- **(Optional)**: Create a ZIP archive of the output directory *(including all subdirectories and files)* and save it in the parent directory of the output directory.
  - Generate a SHA256 checksum of the archive file and use [Trusted Timestamping Authority](https://en.wikipedia.org/wiki/Trusted_timestamping) (TSA) to digitally sign the checksum for integrity and authenticity verification.

---

## Facts ✌🏻

- Group the exported PEM certificates into directories by the issuing CA
- Use the sample CSV [file](certificates.csv) in this repository as input file
  - `certificate_der` column contains DER encoded certificate in hex
- Use [cobra](https://github.com/spf13/cobra), [viper](https://github.com/spf13/viper) for configuration and CLI interactions
- Use [testify](https://github.com/stretchr/testify) for unit testing

Good luck, and may this challenge be ever in your favor!

---

## Getting Started ⚡

- Clone this repository
- Inspect the **Bash script** below

  ```shell
  #!/usr/bin/env bash

  while IFS= read -r line; do
    ca_name="$(echo ${line} | awk -F';' '{print $1}' -)"
    certificate_serial="$(echo ${line} | awk -F';' '{print $2}' -)"
    certificate_der="$(echo ${line} | awk -F';' '{print $3}' -)"
    target_dir="${ca_name// /_}"
 
    # Create directory if it does not exist
    [ -d "${target_dir}" ] || mkdir "${target_dir}"
 
    # Save PEM formatted certificate
    echo "${certificate_der}" | xxd -r -p - | openssl x509 -inform der > "${target_dir}/${certificate_serial}.pem"
  done < certificates.csv
    ```

  > Oh, and before you ask - the sample CSV file **does not** include a `certificate_serial` column. Why? Because we expect you to extract the serial numbers in Go like the coding wizard you are. 🔮💻

- Start coding your Go solution in [exporter.go](cmd/export.go#L55)
- Run tests, ensure coverage
- Submit your solution and bask in the glory of your success!

## Where to Begin? 🏁

To kick things off, check out the `exporter.Export` function - it’s a starting point for tackling this challenge.

---

## Bonus: Automatically Generated Documentation 📚

Take a peek at the automatically **generated** documentation [here](docs/certificate-exporter.md). It's like magic, but with more Markdown and fewer wizards. 🧙✨
