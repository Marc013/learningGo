# ======================================
# VSCode Dev Container - Ubuntu Noble
# Includes: Terraform, GraphViz, Azure CLI,
# PowerShell + Az modules, Oh My Posh + Nerd Fonts
# ======================================

FROM mcr.microsoft.com/devcontainers/base:noble

# -------------------------------
# Environment
# -------------------------------
ENV DEBIAN_FRONTEND=noninteractive
ENV POWERSHELL_TELEMETRY_OPTOUT=1

# -------------------------------
# Fix apt sources for Ubuntu Noble
# -------------------------------
RUN sed -i 's|http://archive.ubuntu.com/ubuntu/|http://old-releases.ubuntu.com/ubuntu/|g' /etc/apt/sources.list || true && \
    apt-get clean && apt-get update

# -------------------------------
# Install basic utilities
# -------------------------------
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    file \
    fontconfig \
    gnupg \
    lsb-release \
    unzip \
    wget \
    sudo \
    git \
    && rm -rf /var/lib/apt/lists/*

# -------------------------------
# Install Go (Golang)
# -------------------------------
ARG GO_VERSION=1.26.4
RUN curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" -o /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz

ENV GOPATH=/home/vscode/go
ENV GOBIN=${GOPATH}/bin
ENV PATH="/usr/local/go/bin:${GOBIN}:${PATH}"

RUN mkdir -p ${GOPATH} && chown -R vscode:vscode ${GOPATH}

# -------------------------------
# Install common Go developer tools
# -------------------------------
RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# -------------------------------
# Install Azure CLI
# -------------------------------
RUN curl -sL https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor | tee /etc/apt/trusted.gpg.d/microsoft.gpg > /dev/null && \
    echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/azure-cli.list && \
    apt-get update && apt-get install -y azure-cli && rm -rf /var/lib/apt/lists/*

# -------------------------------
# Install PowerShell
# -------------------------------
RUN wget https://github.com/PowerShell/PowerShell/releases/download/v7.6.3/powershell-lts_7.6.3-1.deb_amd64.deb && \
    dpkg -i powershell-lts_7.6.3-1.deb_amd64.deb || apt-get install -f -y && \
    rm powershell-lts_7.6.3-1.deb_amd64.deb

# -------------------------------
# Install PowerShell modules
# -------------------------------
RUN pwsh -Command "Register-PSResourceRepository -Name MAR -Uri https://mcr.microsoft.com -ApiVersion ContainerRegistry -Trusted; \
                    Set-PSResourceRepository -Name PSGallery -Trusted; \
                    Install-PSResource -Name Az.Accounts -Version 5.5.0 -Scope AllUsers -Repository MAR; \
                    Install-PSResource -Name Az.Resources -Version 10.0.0 -Scope AllUsers -Repository MAR"

# -------------------------------
# Copy PowerShell profiles
# -------------------------------
COPY ./powerShellProfile.ps1 /home/vscode/.config/powershell/Microsoft.PowerShell_profile.ps1
COPY ./powerShellProfile.ps1 /home/vscode/.config/powershell/Microsoft.VSCode_profile.ps1

# -------------------------------
# Install Nerd Fonts
# -------------------------------
RUN wget -O /tmp/Meslo.zip https://github.com/ryanoasis/nerd-fonts/releases/download/v3.4.0/Meslo.zip && \
    unzip /tmp/Meslo.zip -d /usr/share/fonts/truetype/nerd-fonts && \
    fc-cache -fv && rm /tmp/Meslo.zip

# -------------------------------
# Install Oh My Posh
# -------------------------------
RUN wget https://github.com/JanDeDobbeleer/oh-my-posh/releases/latest/download/posh-linux-amd64 -O /usr/local/bin/oh-my-posh && \
    chmod +x /usr/local/bin/oh-my-posh

# -------------------------------
# Configure Oh My Posh for vscode user
# -------------------------------
USER vscode
RUN mkdir -p /home/vscode/.poshthemes && \
    wget https://github.com/JanDeDobbeleer/oh-my-posh/releases/latest/download/themes.zip -O /tmp/themes.zip && \
    unzip /tmp/themes.zip -d /home/vscode/.poshthemes && \
    chmod u+rw /home/vscode/.poshthemes/*.omp.json && \
    rm /tmp/themes.zip && \
    echo 'eval "$(oh-my-posh init bash --config ~/.poshthemes/blue-owl.omp.json)"' >> /home/vscode/.bashrc && \
    echo 'eval "$(oh-my-posh init zsh --config ~/.poshthemes/blue-owl.omp.json)"' >> /home/vscode/.zshrc

# -------------------------------
# Set default shell to PowerShell
# -------------------------------
ENV SHELL=/usr/bin/pwsh
