{ pkgs ? import <nixpkgs> {} }:

pkgs.stdenv.mkDerivation rec {
  name = "dextryx";
  src = ./.;

  buildInputs = [ pkgs.go ];

  buildPhase = ''
    export GOPATH=$(mktemp -d)
    mkdir -p $GOPATH/src/${name}
    cp -r ./* $GOPATH/src/${name}/
    cd $GOPATH/src/${name}
    go build -o $out/bin/my-go-server
  '';

  installPhase = ''
    mkdir -p $out/bin
    cp -r ./* $out/bin/
  '';
}
