{
  description = "MongoDB exporter development environment";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs, ... }:
    let
      systems = [
        "aarch64-darwin"
        "x86_64-darwin"
        "aarch64-linux"
        "x86_64-linux"
      ];
      forAllSystems = nixpkgs.lib.genAttrs systems;
      mkPkgs = system: import nixpkgs {
        inherit system;
        config.allowUnfreePredicate = pkg: nixpkgs.lib.getName pkg == "mongodb";
      };
      mongoDB50 = pkgs:
        pkgs.stdenvNoCC.mkDerivation (finalAttrs: {
          pname = "mongodb";
          version = "5.0.31";

          src = pkgs.fetchurl {
            url = "https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-${finalAttrs.version}.tgz";
            hash = "sha256-EYOeDK5EY8q45em72+sqXkF1D0Ke1Y/qwghT6qPr3N8=";
          };

          dontUnpack = true;
          installPhase = ''
            mkdir -p "$out/bin"
            for binary in mongo mongod mongos; do
              ${pkgs.gnutar}/bin/tar --extract --file "$src" --gzip \
                --strip-components=2 --directory "$out/bin" \
                "mongodb-macos-x86_64-${finalAttrs.version}/bin/$binary"
            done
          '';

          meta = {
            description = "MongoDB Community Server 5.0.31 x86_64 macOS binaries";
            homepage = "https://www.mongodb.com/";
            license = pkgs.lib.licenses.sspl;
            mainProgram = "mongod";
            platforms = [ "aarch64-darwin" ];
          };
        });
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = mkPkgs system;
        in
        pkgs.lib.optionalAttrs (system == "aarch64-darwin") {
          mongodb-5_0 = mongoDB50 pkgs;
        });

      devShells = forAllSystems (system:
        let
          pkgs = mkPkgs system;
        in
        {
          default = pkgs.mkShell {
            packages = with pkgs; [
              go
              gnumake
              prometheus
              python3
            ] ++ pkgs.lib.optionals (system == "aarch64-darwin") [
              self.packages.${system}.mongodb-5_0
            ];
          };
        });
    };
}
