## Compilation

Install `Go >= 1.23`, then get the source code:

    git clone https://github.com/audit-brands/assurcast-reader.git

Then run one of the corresponding commands:

    # create an executable for the host os
    make build_macos    # -> _output/macos/Assurcast Reader.app
    make build_linux    # -> _output/linux/assurcast-reader
    make build_windows  # -> _output/windows/assurcast-reader.exe

    # host-specific cli version (no gui)
    make build_default  # -> _output/assurcast-reader

    # ... or start a dev server locally
    make serve          # starts a server at http://localhost:7049

    # ... or build a docker image
    docker build -t assurcast-reader -f etc/dockerfile .

## ARM compilation

The instructions below are to cross-compile Assurcast Reader to `Linux/ARM*`.

Build:

    docker build -t assurcast-reader.arm -f etc/dockerfile.arm .

Test:

    # inside host
    docker run -it --rm assurcast-reader.arm

    # then, inside container
    cd /root/out
    qemu-aarch64 -L /usr/aarch64-linux-gnu/ assurcast-reader.arm64

Extract files from images:

    CID=$(docker create assurcast-reader.arm)
    docker cp -a "$CID:/root/out" .
    docker rm "$CID"
