{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
    py = pkgs.writeShellScriptBin "py" ''
      python3 src/main.py
    '';
  in {
    devShells.${system}.default = pkgs.mkShell {
      packages =
        [
          py
        ]
        ++ (with pkgs; [
          python3
          python312Packages.pygame-ce
          python312Packages.pytmx
          tiled
        ]);

      shellHook = "tmux";
    };
  };
}
