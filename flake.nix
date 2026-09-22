{
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

  outputs = { self, nixpkgs }:
    let
      system = "aarch64-darwin";
      pkgs = import nixpkgs { inherit system; };
    in {
      packages.${system}.goCurrent = pkgs.go;
      devShells.${system}.default = pkgs.mkShell {
        packages = [ pkgs.go pkgs.goreleaser pkgs.qemu ];
      };
    };
}
